//go:build !windows

package print

import "github.com/OpenPrinting/goipp"

// Printer contains a number of printer properties. This struct makes interacting
// with the printer easier.
type Printer struct {
	name     string
	location string
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
	}
	return pr
}

// Name retrieves the printer's name as returned by CUPS.
func (pr *Printer) Name() string {
	return pr.name
}

// Location retrieves the printer's location as returned by CUPS.
func (pr *Printer) Location() string {
	return pr.location
}
