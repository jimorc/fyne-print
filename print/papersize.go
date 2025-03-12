package print

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/lang"
)

// Margin is the paper margin. That is, the margins that will not be printed.
type Margin struct {
	top    float32
	bottom float32
	left   float32
	right  float32
}

// NewMargin creates a Margin struct for the values input.
func NewMargin(top, bottom, left, right float32) Margin {
	return Margin{top: top, bottom: bottom, left: left, right: right}
}

// Papersize defines the size of the paper.
type PaperSize struct {
	psName string
	name   string
	width  float32
	height float32
}

// newPaperSize creates a PaperSize struct for the values input.
//
// Params:
//
//	psName is the name of the paper size. This would normally be the IPP media size name
//
// registered with IANA. If the media size is not registered, it may be some other name.
//
//	name is the English name of the media size. It is translated to the system language if the
//
// translation exists.
//
//	width is the width in 100ths of a mm for the media in portrait mode.
//	height is the height in 100ths of a mm for the media in portrait mode.
func newPaperSize(psName string, name string, width float32, height float32) PaperSize {
	ps := PaperSize{psName: psName, width: width, height: height}
	ps.name = lang.L(name)
	return ps
}
