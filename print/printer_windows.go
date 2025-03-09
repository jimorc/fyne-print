//go:build windows

package print

// Printer is a struct containing information related to the printer.
type Printer struct {
	printerInfo2 PrinterInfo2
}

//NewPrinter creates a Printer struct based on information provided in the PrinterInfo2 argument.
func NewPrinter(pInfo2 *PrinterInfo2) *Printer {
	p := &Printer{printerInfo2: *pInfo2}
	return p
}

//Name returns the name of the printer.
func (p *Printer) Name() string {
	return p.printerInfo2.Name()
}

func (p *Printer) Location() string {
	return p.printerInfo2.Location()
}
