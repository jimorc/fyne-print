package print

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
)

// Printer is a struct containing information related to the printer.
type Printer struct {
	handle       syscall.Handle
	printerInfo2 PrinterInfo2
	pSizes       paperSizes
}

// newPrinter creates a Printer struct based on information provided in the PrinterInfo2 argument.
func newPrinter(pInfo2 *PrinterInfo2) *Printer {
	p := &Printer{printerInfo2: *pInfo2}
	printerDefs := newPrinterDefaults("RAW", pInfo2.DevMode,
		printerAccessUse)

	prHandle := openPrinter(p.Name(), printerDefs)
	p.handle = prHandle
	return p
}

// Comment returns the comment set in the printer properties.
func (p *Printer) Comment() string {
	return p.printerInfo2.Comment()
}

// defaultPrinterSize returns the PaperSize corresponding to the default paper size for the printer.
func (p *Printer) defaultPaperSize() *PaperSize {
	dm := p.printerInfo2.DevMode.dmPaperSize
	if dm != dmPaperNone {
		ps := stdPaperSizes.findPaperSizeFromDmPaperSize(dm)
		if ps != nil {
			return ps
		}
	} else {
		w := p.printerInfo2.DevMode.dmPaperWidth
		h := p.printerInfo2.DevMode.dmPaperLength
		ps := stdPaperSizes.findPaperSizeFromWindowsPaperSize(fyne.NewSize(float32(w), float32(h)))
		if ps != nil {
			return ps
		}
	}
	return nil
}

// Location returns the location set in the printer properties.
func (p *Printer) Location() string {
	return p.printerInfo2.Location()
}

// Name returns the name of the printer.
func (p *Printer) Name() string {
	return p.printerInfo2.Name()
}

// paperSizes returns the paper sizes for the printer. The first time
// this method is called, the paper sizes are retrieved.
func (p *Printer) paperSizes() (paperSizes, error) {
	if p.pSizes.isEmpty() {
		err := p.retrievePaperSizes()
		if err != nil {
			return p.pSizes, err
		}
	}
	return p.pSizes, nil
}

func (p *Printer) retrievePaperSizes() error {
	// First, get the Windows paper names.
	var pNames [maxNumPaperSizes][paperNameSize]uint16
	count, err := deviceCapabilities(
		p.Name(),
		p.PortName(),
		dcPaperNames,
		uintptr(unsafe.Pointer(&pNames)),
		p.printerInfo2.DevMode)

	if count <= 0 && err != syscall.Errno(0) {
		fyne.LogError("Error retrieving paper names", err)
		eStr := fmt.Sprintf("error retrieving paper names: %s", err.Error())
		return errors.New(eStr)
	}
	// Now get the paper sizes. We will use these to attempt to get the standard
	// paper sizes.
	var pSizes [maxNumPaperSizes]intPaperSize
	count, err = deviceCapabilities(
		p.Name(),
		p.PortName(),
		dcPaperSize,
		uintptr(unsafe.Pointer(&pSizes)),
		p.printerInfo2.DevMode)
	if count <= 0 && err != syscall.Errno(0) {
		fyne.LogError("Error retrieving paper sizes", err)
		eStr := fmt.Sprintf("error retrieving paper sizes: %s", err.Error())
		return errors.New(eStr)
	}
	p.pSizes.empty()
	for i := 0; i < int(count); i++ {
		pSize := intPaperSize(pSizes[i])
		// ***
		pName := ([paperNameSize]uint16)((pNames[i]))
		n := syscall.UTF16ToString(pName[:])
		fmt.Printf("%s: %v\n", n, pSize)
		// ***
		ps := fyne.NewSize(float32(pSize.w), float32(pSize.h))
		paperSize := stdPaperSizes.findPaperSizeFromWindowsPaperSize(ps)
		if paperSize != nil {
			p.pSizes.add(*paperSize)
		} else {
			pName := ([paperNameSize]uint16)((pNames[i]))
			n := syscall.UTF16ToString(pName[:])
			p.pSizes.add(*newPaperSize(n, n, dmPaperNone, ps.Width*10, ps.Height*10))
			fmt.Printf("Added paper size %s: %v\n", n, ps)
		}
	}
	return nil
}

// PortName returns the port name of the printer.
func (p *Printer) PortName() string {
	return p.printerInfo2.PortName()
}
