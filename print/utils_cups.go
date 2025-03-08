//go:build !windows

package print

import (
	"bytes"
	"net/http"
	"os"

	"github.com/OpenPrinting/goipp"
)

const localCupsURI = "http://localhost:631"

// createGroupsFromMessage parses a goipp.Message and returns a slice of all
// goipp.Group objects in the message.
func createGroupsFromMessage(message *goipp.Message) (*[]goipp.Group, error) {
	groups := &[]goipp.Group{}
	// this print is temporary
	message.Print(os.Stdout, true)
	for _, group := range message.Groups {
		*groups = append(*groups, group)
	}
	return groups, nil
}

// generateRequest creates and encodes an ipp request based on the arguments to the function.
func GenerateRequest(op goipp.Op, printerUri string, attributes string) ([]byte, error) {
	m := goipp.NewRequest(goipp.DefaultVersion, op, 1)
	m.Operation.Add(goipp.MakeAttribute("attributes-charset",
		goipp.TagCharset, goipp.String("utf-8")))
	// Always use en-us as language. This is the default for CUPS. Translations
	// are provided elsewhere.
	m.Operation.Add(goipp.MakeAttribute("attributes-natural-language",
		goipp.TagLanguage, goipp.String("en-us")))
	m.Operation.Add(goipp.MakeAttribute("printer-uri",
		goipp.TagURI, goipp.String(printerUri)))
	m.Operation.Add(goipp.MakeAttribute("requested-attributes",
		goipp.TagKeyword, goipp.String(attributes)))

	return m.EncodeBytes()
}

// getResponseGroups posts a message to the specified URI, retrieves the response, and
// returns a slice of all top-level groups in the response.
func getResponseGroups(op goipp.Op, uri string, attributes string) (*[]goipp.Group, error) {
	request, err := GenerateRequest(op, uri, attributes)
	if err != nil {
		return &[]goipp.Group{}, err
	}
	response, err := http.Post(uri, goipp.ContentType, bytes.NewBuffer(request))
	if err != nil {
		return &[]goipp.Group{}, err
	}
	var msg goipp.Message
	err = msg.Decode(response.Body)
	if err != nil {
		return &[]goipp.Group{}, err
	}

	groups, err := createGroupsFromMessage(&msg)
	return groups, err
}
