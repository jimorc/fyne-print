package print

import (
	"fmt"
	"math"

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
	psN     string
	n       string
	winSize dmPaperSize
	w       float32
	h       float32
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
//	wSize is the Windows paper size. This is one of the values defined in papersizes.go.
//
//	width is the width in 100ths of a mm for the media in portrait mode.
//	height is the height in 100ths of a mm for the media in portrait mode.
func newPaperSize(psName string, name string, wSize dmPaperSize, width float32, height float32) PaperSize {
	ps := PaperSize{psN: psName, winSize: wSize, w: width, h: height}
	ps.n = lang.L(name)
	return ps
}

// Height retrieves the height in 100ths of a mm of the PaperSize object.
func (ps PaperSize) height() float32 {
	return ps.h
}

// Name retrieves the translated paper size common name.
func (ps PaperSize) name() string {
	return ps.n
}

// PSName retrieves the paper size name. For paper sizes that are IANA registered
// IPP media sizes, it is that size. For non-registered media sizes, it would either
// be a name specified by the printer, or a custom name.
func (ps PaperSize) psName() string {
	return ps.psN
}

// SizeInMM returns the size of the paper in mm.
func (ps PaperSize) SizeInMM() fyne.Size {
	return fyne.NewSize(ps.w/100, ps.h/100)
}

// SizeInInches returns the size of the paper in inches.
func (ps PaperSize) SizeInInches() fyne.Size {
	return fyne.NewSize(ps.w/2540, ps.h/2540)
}

// getWidthr retrieves the width of the paper in 100ths of a mm
func (ps PaperSize) width() float32 {
	return ps.w
}

// paperNameFromDimensions creates a paper size name from its dimensions.
func paperNameFromDimensions(w, h float32) string {
	wFormat := createFormat(w / 100.)
	hFormat := createFormat(h / 100.)
	format := fmt.Sprintf("%sx%smm", wFormat, hFormat)
	name := fmt.Sprintf(format, math.Round(float64(w))/100., math.Round(float64(h))/100.)
	return name
}

// createFormat is a helper function that creates the format for a paper size dimension.
func createFormat(dim float32) string {
	dim = float32(math.Round(float64(dim*100.))) / 100.
	if float32(int32(dim)) == dim {
		return "%.0f"
	} else if float32(int32(dim*10)) == dim*10 {
		return "%.1f"
	}
	return "%.2f"
}
