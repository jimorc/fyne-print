package print

import (
	"fyne.io/fyne/v2"
	"github.com/OpenPrinting/goipp"
)

type Printers struct {
	printers []Printer
}

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
			p.AddPrinter(pr)
		}
	}
	return p, err
}

func (p *Printers) AddPrinter(pr *Printer) {
	p.printers = append(p.printers, *pr)
}
