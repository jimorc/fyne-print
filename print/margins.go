package print

import (
	"fmt"
	"strings"
)

// Margins is a struct containing the margins for each side of the media.
type Margins struct {
	left   float32
	right  float32
	top    float32
	bottom float32
}

// String converts the Margins object to a string.
func (m *Margins) String() string {
	var s strings.Builder
	s.WriteString("Margins:")
	s.WriteString(fmt.Sprintf("    top: %.2f", m.top))
	s.WriteString(fmt.Sprintf("    bottom: %.2f", m.bottom))
	s.WriteString(fmt.Sprintf("    left: %.2f", m.left))
	s.WriteString(fmt.Sprintf("    right: %.2f\n", m.right))
	return s.String()
}
