package print

import (
	"fmt"

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

// getPrinterIndexByName returns the index of the printer within the printers
// object that has the specified name.
//
// Params:
//
//	name is the name of the printer to find.
//
// Returns:
//
//	index of the matching printer within the Printers object, or -1 on error.
//	error if matching printer is not found, or nil if found.
func (p *Printers) getPrinterIndexByName(name string) (int, error) {
	for i, pr := range p.printers {
		if pr.Name() == name {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Printer with name: \"%s\" not found", name)
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
