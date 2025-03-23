//go:build !windows

package print

// #include "cups/cups.h"
import "C"

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

// Instance retrieves the printer Instance from the CUPS desination object
// associated with the Printer.
// This may be an empty string.
func (p *Printer) Instance() string {
	return C.GoString(p.dest.instance)
}

// Name retrieves the printer Name from the CUPS destination object associated
// with the Printer.
func (p *Printer) Name() string {
	return C.GoString(p.dest.name)
}
