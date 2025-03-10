package print

import (
	"fmt"
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
)

// Printer is a struct containing information related to the printer.
type Printer struct {
	printerInfo2 PrinterInfo2
	papNames     []string
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

func (p *Printer) paperNames() []string {
	if len(p.papNames) > 0 {
		return p.papNames
	}
	// 256 is more than enough for all paper names retrieved from any Windows printer driver.
	var pNames [maxNumPaperSizes][paperNameSize]uint16
	count, err := deviceCapabilities(
		p.Name(),
		p.PortName(),
		dcPaperNames,
		uintptr(unsafe.Pointer(&pNames)),
		p.printerInfo2.DevMode)

	if err != syscall.Errno(0) {
		fyne.LogError("Error retrieving paper names", err)
		return []string{}
	}
	var paperNames []string
	for i := 0; i < int(count); i++ {
		pName := ([paperNameSize]uint16)((pNames[i]))
		paperNames = append(paperNames, syscall.UTF16ToString(pName[:]))
		fmt.Printf("%s\n", syscall.UTF16ToString(pName[:]))
	}
	p.papNames = paperNames
	return paperNames
}

// PortName returns the port name of the printer.
func (p *Printer) PortName() string {
	return p.printerInfo2.PortName()
}
