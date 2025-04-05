//go:build windows

package print

import (
	"fmt"
	"strings"
)

// MediaSize contains media size info for Windows paper sizes.
type MediaSize struct {
	name            string
	paperSize       paperSize
	width           float32
	height          float32
	imageableWidth  float32
	imageableHeight float32
	offsetX         float32
	offsetY         float32
	hasData         bool
}

// newMediaSize creates a partially-constructed MediaSize Windows object.
// setData should be called to fill in the rest of the fields.
func newMediaSize(name string, size paperSize) MediaSize {
	return MediaSize{
		name:      name,
		paperSize: size,
		hasData:   false,
	}
}

// newMediaSizeWithData creates a fully-constructed MediaSize Windows object.
func newMediaSizeWithData(name string, size paperSize, width, height, imageableWidth,
	imageableHeight, offsetX, offsetY float32) MediaSize {
	return MediaSize{
		name:            name,
		paperSize:       size,
		width:           width,
		height:          height,
		imageableWidth:  imageableWidth,
		imageableHeight: imageableHeight,
		offsetX:         offsetX,
		offsetY:         offsetY,
		hasData:         true,
	}
}

// setData sets size and margins data for a Windows MediaSize object
func (ms *MediaSize) setData(width, height, imageableWidth, imageableHeight, offsetX, offsetY float32) {
	ms.width = width
	ms.height = height
	ms.imageableWidth = imageableWidth
	ms.imageableHeight = imageableHeight
	ms.offsetX = offsetX
	ms.offsetY = offsetY
	ms.hasData = true
}

// Width returns the width of a MediaSize object.
func (ms *MediaSize) Width() float32 {
	return float32(ms.width)
}

// Height returns the height of a MediaSize object.
func (ms *MediaSize) Height() float32 {
	return float32(ms.height)
}

// LocalName returns the local name of a MediaSize object.
func (ms *MediaSize) LocalName() string {
	return ms.name
}

// Margins returns the margins of a MediaSize object.
func (ms *MediaSize) Margins() Margins {
	m := Margins{}
	m.left = ms.offsetX
	m.right = ms.width - (ms.offsetX + ms.imageableWidth)
	m.top = ms.offsetY
	m.bottom = ms.height - (ms.offsetY + ms.imageableHeight)
	return m
}

// PaperSize returns the paperSize value of a MediaSize object.
// paperSize is one of the dmPaperSize values that are returned by
// calling DeviceCapabilities requesting DC_PAPERS values.
func (ms *MediaSize) PaperSize() paperSize {
	return ms.paperSize
}

// String converts the MediaSize object to a string (for printing)
func (ms *MediaSize) String() string {
	var s strings.Builder
	s.WriteString("MediaSize:\n")
	s.WriteString(fmt.Sprintf("    Media Name: %s\n", ms.LocalName()))
	s.WriteString(fmt.Sprintf("    PaperSize: %d, %s\n", ms.paperSize, ms.paperSize.String()))
	if ms.hasData {
		s.WriteString(fmt.Sprintf("    Width: %.2f\n", ms.Width()))
		s.WriteString(fmt.Sprintf("    Height: %.2f\n", ms.Height()))
		s.WriteString(ms.Margins().String())
	} else {
		s.WriteString("    No data\n")
	}
	return s.String()
}
