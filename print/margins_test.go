package print

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMargins_String(t *testing.T) {
	m := Margins{}
	s := m.String()
	assert.Equal(t, "Margins:    top: 0.00    bottom: 0.00    left: 0.00    right: 0.00\n", s)

	m = Margins{top: 1.0, bottom: 2.0, left: 3.0, right: 4.0}
	assert.Equal(t, "Margins:    top: 1.00    bottom: 2.00    left: 3.00    right: 4.00\n", m.String())
}
