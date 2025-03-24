//go:build !windows

package print

// #include "cups/cups.h"
import "C"
import (
	"errors"
	"strconv"
	"unsafe"

	"fyne.io/fyne/v2"
)

// Printer represents a CUPS printer
type Printer struct {
	http       *C.http_t
	dest       *C.cups_dest_t
	dinfo      *C.cups_dinfo_t
	caps       capabilities
	mediaSizes MediaSizes
}

// newPrinter creates a new Printer object.
//
// Params:
//
//	dest is the CUPS destination for the printer.
func newPrinter(dest *C.cups_dest_t) *Printer {
	p := &Printer{dest: dest}
	opts := p.Options()
	caps := opts["printer-type"]
	if len(caps) > 0 {
		c, _ := strconv.Atoi(caps)
		p.caps = capabilities(uint32(c))
	}
	p.http = C.cupsConnectDest(p.dest, C.CUPS_DEST_FLAGS_NONE,
		2000, nil, nil, 0, nil, nil)
	p.dinfo = C.cupsCopyDestInfo(p.http, p.dest)

	mCount := C.cupsGetDestMediaCount(p.http, p.dest, p.dinfo, 0)
	for i := 0; i < int(mCount); i++ {
		var mSize C.cups_size_t

		res := C.cupsGetDestMediaByIndex(p.http, p.dest, p.dinfo, C.int(i),
			0, &mSize)
		if res == 0 {
			e := C.cupsLastErrorString()
			fyne.LogError("Error getting media size", errors.New(C.GoString(e)))
			continue
		}
		s := newMediaSize(&mSize, p)
		p.mediaSizes.Add(s)
	}
	return p
}

// AddMediaSize adds a MediaSize object to the printer object.
func (p *Printer) AddMediaSize(s *MediaSize) {
	p.mediaSizes.Add(*s)
}

// Capabilities retrieves the printer's capabilities value.
func (p *Printer) Capabilities() capabilities {
	return p.caps
}

// Close frees any CUPS memory allocations for the Printer.
func (p *Printer) Close() {
	C.httpClose(p.http)
	C.cupsFreeDestInfo(p.dinfo)
	p.dinfo = nil
	p.http = nil
	p.dest = nil
}

// Instance retrieves the printer Instance from the CUPS destination object
// associated with the Printer.
// This may be an empty string.
func (p *Printer) Instance() string {
	return C.GoString(p.dest.instance)
}

// IsDefault returns whether the printer's CUPS destination object indicates that
// this printer is the default.
func (p *Printer) IsDefault() bool {
	return !(p.dest.is_default == 0)
}

// MediaSizes returns the media sizes for the printer.
func (p *Printer) MediaSizes() MediaSizes {
	return p.mediaSizes
}

// Name retrieves the printer Name from the CUPS destination object associated
// with the Printer.
func (p *Printer) Name() string {
	return C.GoString(p.dest.name)
}

// Options retrieves a map containing the options values as retrieved as
// part of the printer's dest value.
func (p *Printer) Options() map[string]string {
	options := make(map[string]string)
	oPtr := uintptr(unsafe.Pointer(p.dest.options))
	for i := 0; i < int(p.dest.num_options); i++ {
		// use of unsafe.Pointer on next line OK.
		opt := (*C.cups_option_t)(unsafe.Pointer(oPtr))
		options[C.GoString(opt.name)] = C.GoString(opt.value)
		oPtr += uintptr(unsafe.Sizeof(opt.name)) + uintptr(unsafe.Sizeof(opt.value))
	}
	return options
}
