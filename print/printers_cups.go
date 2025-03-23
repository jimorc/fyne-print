//go:build !windows

package print

// #include "cups/cups.h"
import "C"
import "unsafe"

// Printers is the object containing all of the CUPS Printer objects.
type Printers struct {
	dests    *C.cups_dest_t
	nDests   C.int
	Printers []Printer
}

// NewPrinters generates Printer objects for each CUPS printer.
func NewPrinters() *Printers {
	ps := &Printers{}
	var dests *C.cups_dest_t
	ps.nDests = C.cupsGetDests(&ps.dests)
	pDest := unsafe.Pointer(ps.dests)

	for i := 0; i < int(ps.nDests); i++ {
		d := (*C.cups_dest_t)(unsafe.Pointer(uintptr(pDest) + uintptr(i)*unsafe.Sizeof(*dests)))
		pr := newPrinter(d)
		ps.Printers = append(ps.Printers, *pr)
	}
	return ps
}

// Close frees CUPS memory assigned to each Printer.
func (p *Printers) Close() {
	for _, pr := range p.Printers {
		pr.Close()
	}
	C.cupsFreeDests(p.nDests, p.dests)
	p.Printers = nil
}
