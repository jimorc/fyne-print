package print

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPrintersFromGroups(t *testing.T) {
	msg := createGetPrintersResponse()
	groups, err := createGroupsFromMessage(msg)
	assert.Nil(t, err)

	p := newPrintersFromGroups(groups)
	assert.Equal(t, "Printer1", p.printers[0].Name())
}
