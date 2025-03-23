//go:build !windows

package print

// #include "cups/cups.h"
import "C"
import "fmt"

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
	p := &Printer{}
	c := C.cupsCopyDest(dest, 1, &p.dest)
	fmt.Printf("newPrinter: c=%d\n", int(c))
	return p
}

// Close frees any CUPS memory allocations for the Printer.
func (p *Printer) Close() {
	C.cupsFreeDests(1, p.dest)
	p.dest = nil
}
