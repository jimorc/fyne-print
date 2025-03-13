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
	printerInfo2 PrinterInfo2
	pSizes       paperSizes
}

// NewPrinter creates a Printer struct based on information provided in the PrinterInfo2 argument.
func NewPrinter(pInfo2 *PrinterInfo2) *Printer {
	p := &Printer{printerInfo2: *pInfo2}
	return p
}

// Comment returns the comment set in the printer properties.
func (p *Printer) Comment() string {
	return p.printerInfo2.Comment()
}

// Location returns the location set in the printer properties.
func (p *Printer) Location() string {
	return p.printerInfo2.Location()
}

// Name returns the name of the printer.
func (p *Printer) Name() string {
	return p.printerInfo2.Name()
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

	if err != syscall.Errno(0) {
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
	if err != syscall.Errno(0) {
		fyne.LogError("Error retrieving paper sizes", err)
		eStr := fmt.Sprintf("error retrieving paper sizes: %s", err.Error())
		return errors.New(eStr)
	}
	p.pSizes.empty()
	for i := 0; i < int(count); i++ {
		pSize := intPaperSize(pSizes[i])
		ps := fyne.NewSize(float32(pSize.w), float32(pSize.h))
		paperSize := stdPaperSizes.findPaperSizeFromWindowsPaperSize(ps)
		if paperSize != nil {
			p.pSizes.add(*paperSize)
		} else {
			pName := ([paperNameSize]uint16)((pNames[i]))
			n := syscall.UTF16ToString(pName[:])
			p.pSizes.add(newPaperSize(n, n, ps.Width*10, ps.Height*10))
			fmt.Printf("Added paper size %s: %v\n", n, ps)
		}
	}
	return nil
}

// PortName returns the port name of the printer.
func (p *Printer) PortName() string {
	return p.printerInfo2.PortName()
}

func (p *Printer) paperNames() []string {
	var names []string
	for _, ps := range p.pSizes.sizes {
		names = append(names, ps.name())
	}
	return names
}
