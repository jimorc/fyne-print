package printer

import (
	"bytes"
	"net/http"

	"github.com/OpenPrinting/goipp"
)

const uri = "http://localhost:631"

func GetPrinters() ([]string, error) {
	request, err := makeGetPrintersRequest()
	if err != nil {
		return nil, err
	}
	response, err := http.Post(uri, goipp.ContentType, bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}

	var respMsg goipp.Message
	err = respMsg.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	var printerNames []string
	//	respMsg.Print(os.Stdout, false)
	for _, group := range respMsg.Groups {
		if group.Tag == goipp.TagPrinterGroup {
			for _, attr := range group.Attrs {
				if attr.Name == "printer-name" {
					printerNames = append(printerNames, attr.Values.String())
				}
			}
		}
	}
	return printerNames, nil
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
