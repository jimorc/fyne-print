package print

import (
	"testing"

	"github.com/OpenPrinting/goipp"
	"github.com/stretchr/testify/assert"
)

func TestNewPrinter(t *testing.T) {
	var pr *Printer
	msg := createTestPrintersResponse()
	groups, err := createGroupsFromMessage(msg)
	assert.Nil(t, err)
	for _, group := range *groups {
		if group.Tag != goipp.TagPrinterGroup {
			continue
		}
		pr = newPrinter(group)
		break
	}
	assert.NotNil(t, pr)
	assert.Equal(t, "Printer1", pr.Name())

}
