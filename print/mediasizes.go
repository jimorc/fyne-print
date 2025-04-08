package print

import (
	"strings"
)

// MediaSizes is a slice of MediaSize objects
type MediaSizes []MediaSize

// Add appends a MediaSize object.
func (m *MediaSizes) Add(s MediaSize) {
	*m = append(*m, s)
}

// String converts the MediaSize objects to a string (for printing)
func (m *MediaSizes) String() string {
	var s strings.Builder
	s.WriteString("MediaSizes:\n")
	for _, ms := range *m {
		s.WriteString(ms.String())
	}
	return s.String()
}
