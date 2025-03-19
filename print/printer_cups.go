//go:build !windows

package print

import "github.com/OpenPrinting/goipp"

// Printer contains a number of printer properties. This struct makes interacting
// with the printer easier.
type Printer struct {
	name         string
	location     string
	comment      string
	defPaperSize string
}

// newPrinter creates a Printer object based on CUPS ipp group info.
func newPrinter(ippGroup goipp.Group) *Printer {
	pr := &Printer{}
	for _, attr := range ippGroup.Attrs {
		if attr.Name == "printer-name" {
			pr.name = attr.Values.String()
		}
		if attr.Name == "printer-location" {
			pr.location = attr.Values.String()
		}
		if attr.Name == "printer-info" {
			pr.comment = attr.Values.String()
		}
		if attr.Name == "media-default" {
			pr.defPaperSize = attr.Values.String()
		}
	}
	return pr
}

func (pr *Printer) Comment() string {
	return pr.comment
}

// Location retrieves the printer's location as returned by CUPS.
func (pr *Printer) Location() string {
	return pr.location
}

// Name retrieves the printer's name as returned by CUPS.
func (pr *Printer) Name() string {
	return pr.name
}

// defaultPrinterSize returns the PaperSize corresponding to the default paper size for the printer.
func (p *Printer) defaultPaperSize() *PaperSize {
	return stdPaperSizes.findPaperSizeFromName(p.defPaperSize)
}

// paperSizes returns the paper sizes for the printer. The first time
// this method is called, the paper sizes are retrieved.
func (p *Printer) paperSizes() (paperSizes, error) {
	return paperSizes{}, nil
}
