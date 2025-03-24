//go:build !windows

package print

//#include "cups/cups.h"
import "C"
import (
	"fmt"
	"strings"
)

// Margins is a struct containing the margins for each side of the media.
type Margins struct {
	left   int
	right  int
	top    int
	bottom int
}

// AsString converts the Margins object to a string.
func (m *Margins) AsString() string {
	var s strings.Builder
	s.WriteString("Margins:")
	s.WriteString(fmt.Sprintf("    top: %d", m.top))
	s.WriteString(fmt.Sprintf("    bottom: %d", m.bottom))
	s.WriteString(fmt.Sprintf("    left: %d", m.left))
	s.WriteString(fmt.Sprintf("    right: %d\n", m.right))
	return s.String()
}

// MediaSize contains the PWG name, localized name, width, length, and margins for
// the media.
type MediaSize struct {
	size      *C.cups_size_t
	localName string
}

// newMediaSize creates a MediaSize object from the CUPS media size struct.
// The localized media size name is retrieved.
func newMediaSize(cupsSize *C.cups_size_t, pr *Printer) MediaSize {
	s := MediaSize{size: cupsSize}
	s.localName = C.GoString(C.cupsLocalizeDestMedia(pr.http, pr.dest, pr.dinfo,
		0, cupsSize))
	return s
}

// AsString converts the MediaSize object to a string (for printing).
func (s *MediaSize) AsString() string {
	var b strings.Builder
	b.WriteString("MediaSize:")
	b.WriteString(fmt.Sprintf("    Media Name: %s\n", s.MediaName()))
	b.WriteString(fmt.Sprintf("    Local Name: %s\n", s.LocalName()))
	b.WriteString(fmt.Sprintf("    Width: %d\n", s.Width()))
	b.WriteString(fmt.Sprintf("    Length: %d\n", s.Length()))
	b.WriteString(s.Margins().AsString())
	return b.String()
}

// LocalName retrieves the localized name for the media size.
func (s *MediaSize) LocalName() string {
	return s.localName
}

// MediaName retrieves the media name for the media size. This is usually
// the PWG name.
func (s *MediaSize) MediaName() string {
	return C.GoString(&s.size.media[0])
}

// Width retrieves the width of the media size.
func (s *MediaSize) Width() int {
	return int(s.size.width)
}

// Length retrieves the length of the media size.
func (s *MediaSize) Length() int {
	return int(s.size.length)
}

// Margins retrieves the margins for the media size.
func (s *MediaSize) Margins() *Margins {
	return &Margins{left: int(s.size.left), right: int(s.size.right),
		top: int(s.size.top), bottom: int(s.size.bottom)}
}
