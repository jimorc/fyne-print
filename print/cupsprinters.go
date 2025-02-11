package print

import (
	"fyne.io/fyne/v2"
	"github.com/OpenPrinting/goipp"
)

// Printers contains a slice of Printer objects.
type Printers struct {
	printers []Printer
}

// getNames retrieves the list of all printer names.
func (p *Printers) getNames() []string {
	names := []string{}
	for _, printer := range p.printers {
		names = append(names, printer.Name())
	}
	return names
}

// addPrinter adds a Printer object to the Printers object.
func (p *Printers) addPrinter(pr *Printer) {
	p.printers = append(p.printers, *pr)
}

// newPrinters contacts the CUPs system to retrieve printer info.
// It then creates a Printer object for every printer for which
// info was returned.
func newPrinters() (*Printers, error) {
	groups, err := getResponseGroups(goipp.OpCupsGetPrinters,
		localCupsURI, "all")
	if err != nil {
		fyne.LogError("Error getting CUPS printers", err)
		return &Printers{}, err
	}
	p := newPrintersFromGroups(groups)
	return p, nil
}

// newPrintersFromGroups creates a Printers object for the each printer
// specified in an IPP printer group.
func newPrintersFromGroups(groups *[]goipp.Group) *Printers {
	p := &Printers{}
	for _, group := range *groups {
		if group.Tag == goipp.TagPrinterGroup {
			pr := newPrinter(group)
			p.addPrinter(pr)
		}
	}
	return p
}
