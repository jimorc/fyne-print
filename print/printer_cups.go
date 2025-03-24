//go:build !windows

package print

// #include "cups/cups.h"
import "C"
import (
	"fmt"
	"strconv"
	"unsafe"
)

// Printer represents a CUPS printer
type Printer struct {
	http  *C.http_t
	dest  *C.cups_dest_t
	dinfo *C.cups_dinfo_t
	caps  capabilities
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
	fmt.Printf("Media count: %d\n", mCount)
	for i := 0; i < int(mCount); i++ {
		var mSize C.cups_size_t

		res := C.cupsGetDestMediaByIndex(p.http, p.dest, p.dinfo, C.int(i),
			0, &mSize)

		fmt.Printf("MediaByIndex2 result: %d\n", res)
		if res == 1 {
			fmt.Printf("Media Name: %s\n", C.GoString(&mSize.media[0]))
			fmt.Printf("Media Size: %v\n", mSize)
		}
	}
	return p
}

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

// Name retrieves the printer Name from the CUPS destination object associated
// with the Printer.
func (p *Printer) Name() string {
	return C.GoString(p.dest.name)
}

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
