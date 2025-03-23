//go:build !windows

package print

// #include "cups/cups.h"
import "C"
import "unsafe"

// Printer represents a CUPS printer
type Printer struct {
	dest *C.cups_dest_t
}

// newPrinter creates a new Printer object.
//
// Params:
//
//	dest is the CUPS destination for the printer.
func newPrinter(dest *C.cups_dest_t) *Printer {
	p := &Printer{dest: dest}
	return p
}

// Close frees any CUPS memory allocations for the Printer.
func (p *Printer) Close() {
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
