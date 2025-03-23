//go:build !windows

package print

/*
import (
	"testing"

	"github.com/OpenPrinting/goipp"
	"github.com/stretchr/testify/assert"
)

func TestNewPrintersFromGroups(t *testing.T) {
	msg := createTestPrintersResponse()
	groups, err := createGroupsFromMessage(msg)
	assert.Nil(t, err)

	p := newPrintersFromGroups(groups)
	assert.Equal(t, "Printer1", p.Printers[0].Name())
}

func TestPrinters_getPrinterByName(t *testing.T) {
	tests := []struct {
		name          string
		printerNames  []string
		targetName    string
		expectedFound bool
	}{
		{
			name:          "Printer found",
			printerNames:  []string{"Printer1", "Printer2", "Printer3"},
			targetName:    "Printer2",
			expectedFound: true,
		},
		{
			name:          "Printer not found",
			printerNames:  []string{"Printer1", "Printer2", "Printer3"},
			targetName:    "Printer4",
			expectedFound: false,
		},
		{
			name:          "Empty printers list",
			printerNames:  []string{},
			targetName:    "Printer1",
			expectedFound: false,
		},
		{
			name:          "Single printer",
			printerNames:  []string{"Printer1"},
			targetName:    "Printer1",
			expectedFound: true,
		},
		{
			name:          "Single printer not found",
			printerNames:  []string{"Printer1"},
			targetName:    "Printer2",
			expectedFound: false,
		},
		{
			name:          "Case sensitive",
			printerNames:  []string{"printer1", "Printer2"},
			targetName:    "Printer1",
			expectedFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printers := &Printers{Printers: []Printer{}}
			for _, printerName := range tt.printerNames {
				// We use a fake group here. We don't care about the ipp data
				// in this test.
				printers.addPrinter(newPrinter(goipp.Group{Attrs: goipp.Attributes{goipp.MakeAttribute("printer-name", goipp.TagName, goipp.String(printerName))}}))
			}

			foundPrinter := printers.getPrinterByName(tt.targetName)

			if tt.expectedFound {
				assert.NotNil(t, foundPrinter)
				assert.Equal(t, tt.targetName, foundPrinter.Name())
			} else {
				assert.Nil(t, foundPrinter)
			}
		})
	}
}*/
