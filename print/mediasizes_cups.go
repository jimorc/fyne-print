//go:build !windows

package print

import "strings"

// MediaSizes is a slice of MediaSize objects
type MediaSizes []MediaSize

// Add appends a MediaSize object.
func (m *MediaSizes) Add(s MediaSize) {
	*m = append(*m, s)
}

// AsString converts the MediaSize objects to a string (for printing)
func (m MediaSizes) AsString() string {
	var s strings.Builder
	s.WriteString("MediaSizes:\n")
	for _, ms := range m {
		s.WriteString(ms.AsString())
	}
	return s.String()
}
