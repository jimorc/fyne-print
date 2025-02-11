package print

import (
	"fyne.io/fyne/v2"
	"github.com/OpenPrinting/goipp"
)

// Printers contains a slice of Printer objects.
type Printers struct {
	printers []Printer
}

// addPrinter adds a Printer object to the Printers object.
func (p *Printers) addPrinter(pr *Printer) {
	p.printers = append(p.printers, *pr)
}

// newPrinters contacts the CUPs system to retrieve printer info.
// It then creates a Printer object for every printer for which
// info was returned.
func newPrinters() (*Printers, error) {
	p := &Printers{}
	groups, err := getResponseGroups(goipp.OpCupsGetPrinters,
		localCupsURI, "all")
	if err != nil {
		fyne.LogError("Error getting CUPS printers", err)
	}
	for _, group := range *groups {
		if group.Tag == goipp.TagPrinterGroup {
			pr := newPrinter(group)
			p.addPrinter(pr)
		}
	}
	return p, err
}
