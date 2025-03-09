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

// Comment returns the comment set in the printer properties.
func (p *Printer) Comment() string {
	return p.printerInfo2.Comment()
}

//Name returns the name of the printer.
func (p *Printer) Name() string {
	return p.printerInfo2.Name()
}

// Location returns the location set in the printer properties.
func (p *Printer) Location() string {
	return p.printerInfo2.Location()
}
