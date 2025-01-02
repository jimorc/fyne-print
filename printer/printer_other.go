package printer

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/OpenPrinting/goipp"
)

const uri = "http://localhost:631"

// Printer contains a number of printer properties. This struct will make interacting with the
// printer easier as functionality is added.
type Printer struct {
	Name     string
	Location string
	Model    string
}

// newPrinter creates a new Printer object from goipp.Group data.
func newPrinter(ippGroup *goipp.Group) *Printer {
	p := &Printer{}
	for _, attr := range ippGroup.Attrs {
		if attr.Name == "printer-name" {
			p.Name = attr.Values.String()
		}
		if attr.Name == "printer-location" {
			p.Location = attr.Values.String()
		}
		if attr.Name == "printer-make-and-model" {
			p.Model = attr.Values.String()
		}
	}
	return p
}

// Printers contains a slice of Printer objects that represent the printers available on a system.
type Printers struct {
	Printers []Printer
}

// NewPrinters creates a new Printers object containing all of the printers available at that moment
// on the system.
func NewPrinters() (*Printers, error) {
	p := &Printers{}
	groups, err := getPrinterGroups()
	if err != nil {
		return p, err
	}
	for _, group := range *groups {
		printer := newPrinter(&group)
		p.AddPrinter(printer)
	}
	return p, nil
}

// Clear clears all of the Printer objects from the Printers object
func (pd *Printers) Clear() {
	pd.Printers = []Printer{}
}

// AddPrinter adds the specified Printer object to the Printers object
func (pd *Printers) AddPrinter(pr *Printer) {
	pd.Printers = append(pd.Printers, *pr)
}

// GetNames retrieves the list of all printer names.
func (pd *Printers) GetNames() []string {
	names := []string{}
	for _, printer := range pd.Printers {
		names = append(names, printer.Name)
	}
	return names
}

// GetPrinterIndexByName determines the index of the printer specified by its name
// in the Printers object.
func (pd *Printers) GetPrinterIndexByName(name string) (int, error) {
	for i, pr := range pd.Printers {
		if pr.Name == name {
			return i, nil
		}
	}
	return 0, errors.New("Printer not found")
}

// GetDefaultPrinter retrieves a Printer object representing the default CUPS printer,
// or nil on error.
func (pd *Printers) GetDefaultPrinter() (int, error) {
	p := &Printer{}
	request, err := makeGetDefaultPrinterRequest()
	if err != nil {
		return 0, err
	}
	response, err := http.Post(uri, goipp.ContentType, bytes.NewBuffer(request))
	if err != nil {
		return 0, err
	}

	var respMsg goipp.Message
	err = respMsg.Decode(response.Body)
	if err != nil {
		return 0, err
	}

	for _, group := range respMsg.Groups {
		if group.Tag == goipp.TagPrinterGroup {
			p = newPrinter(&group)
		}
	}
	return pd.GetPrinterIndexByName(p.Name)
}

func getPrinterGroups() (*[]goipp.Group, error) {
	groups := &[]goipp.Group{}
	request, err := makeGetPrintersRequest()
	if err != nil {
		return groups, err
	}
	response, err := http.Post(uri, goipp.ContentType, bytes.NewBuffer(request))
	if err != nil {
		return groups, err
	}

	var respMsg goipp.Message
	err = respMsg.Decode(response.Body)
	if err != nil {
		return groups, err
	}

	for _, group := range respMsg.Groups {
		if group.Tag == goipp.TagPrinterGroup {
			*groups = append(*groups, group)
		}
	}
	return groups, nil
}

func makeGetPrintersRequest() ([]byte, error) {
	m := goipp.NewRequest(goipp.DefaultVersion, goipp.OpCupsGetPrinters, 1)
	m.Operation.Add(goipp.MakeAttribute("attributes-charset",
		goipp.TagCharset, goipp.String("utf-8")))
	m.Operation.Add(goipp.MakeAttribute("attributes-natural-language",
		goipp.TagLanguage, goipp.String("en-us")))
	m.Operation.Add(goipp.MakeAttribute("printer-uri",
		goipp.TagURI, goipp.String(uri)))
	m.Operation.Add(goipp.MakeAttribute("requested-attributes",
		goipp.TagKeyword, goipp.String("all")))

	return m.EncodeBytes()
}

func makeGetDefaultPrinterRequest() ([]byte, error) {
	m := goipp.NewRequest(goipp.DefaultVersion, goipp.OpCupsGetDefault, 1)
	m.Operation.Add(goipp.MakeAttribute("attributes-charset",
		goipp.TagCharset, goipp.String("utf-8")))
	m.Operation.Add(goipp.MakeAttribute("attributes-natural-language",
		goipp.TagLanguage, goipp.String("en-us")))
	m.Operation.Add(goipp.MakeAttribute("printer-uri",
		goipp.TagURI, goipp.String(uri)))
	m.Operation.Add(goipp.MakeAttribute("requested-attributes",
		goipp.TagKeyword, goipp.String("all")))

	return m.EncodeBytes()

}
