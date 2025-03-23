//go:build !windows

package print

/*
// Printers contains a slice of Printer objects.
type Printers struct {
	Printers []Printer
}

// DefaultPrinter returns the name of default printer
// in the Printers object, or nil if no default printer.
func (p *Printers) DefaultPrinter() *Printer {
	groups, err := getResponseGroups(goipp.OpCupsGetDefault,
		localCupsURI, "printer-name")
	if err != nil {
		fyne.LogError("Error getting default CUPS printer", err)
		return nil
	}
	for _, group := range *groups {
		if group.Tag == goipp.TagPrinterGroup {
			pr := newPrinter(group)
			return p.getPrinterByName(pr.Name())
		}
	}
	return nil
}

// getNames retrieves the list of all printer names.
func (p *Printers) getNames() []string {
	names := []string{}
	for _, printer := range p.Printers {
		names = append(names, printer.Name())
	}
	return names
}

// getPrinterByName returns the the printer within the printers
// object that has the specified name.
//
// Params:
//
//	name is the name of the printer to find.
//
// Returns:
//
//	The matching printer within the Printers object, or nil if not found.
func (p *Printers) getPrinterByName(name string) *Printer {
	for i, pr := range p.Printers {
		if pr.Name() == name {
			return &p.Printers[i]
		}
	}
	return nil
}

// addPrinter adds a Printer object to the Printers object.
func (p *Printers) addPrinter(pr *Printer) {
	p.Printers = append(p.Printers, *pr)
}

// newPrinters contacts the CUPs system to retrieve printer info.
// It then creates a Printer object for every printer for which
// info was returned.
func NewPrinters() (*Printers, error) {
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
*/
