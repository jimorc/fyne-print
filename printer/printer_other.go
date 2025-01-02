package printer

import (
	"bytes"
	"net/http"

	"github.com/OpenPrinting/goipp"
)

const uri = "http://localhost:631"

type Printer struct {
	Name string
}

func newPrinter(ippGroup *goipp.Group) *Printer {
	p := &Printer{}
	for _, attr := range ippGroup.Attrs {
		if attr.Name == "printer-name" {
			p.Name = attr.Values.String()
		}
	}
	return p
}

type Printers struct {
	printers []Printer
}

func NewPrinters() (*Printers, error) {
	p := &Printers{}
	groups, err := getPrinterGroups()
	if err != nil {
		return p, err
	}
	for _, group := range *groups {
		printer := newPrinter(&group)
		p.printers = append(p.printers, *printer)
	}
	return p, nil
}

func (pd *Printers) GetNames() []string {
	names := []string{}
	for _, printer := range pd.printers {
		names = append(names, printer.Name)
	}
	return names
}

// GetDefaultPrinter retrieves the name of the default CUPS printer, or nil
func GetDefaultPrinter() (string, error) {
	request, err := makeGetDefaultPrinterRequest()
	if err != nil {
		return "", err
	}
	response, err := http.Post(uri, goipp.ContentType, bytes.NewBuffer(request))
	if err != nil {
		return "", err
	}

	var respMsg goipp.Message
	err = respMsg.Decode(response.Body)
	if err != nil {
		return "", err
	}

	for _, attr := range respMsg.Printer {
		if attr.Name == "printer-name" {
			return attr.Values.String(), nil
		}
	}

	return "", nil
}

// GetPrinters retrieves a slice containing the names of available CUPS printers.
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
