//go:build !windows

package print

//#cgo LDFLAGS: -lcups
//#include "cups/cups.h"
import "C"

func cupsDests() (C.int, *C.cups_dest_t) {

	var dests *C.cups_dest_t
	nDests := C.cupsGetDests(&dests)

	return nDests, dests

}

func cupsGetErrorString() string {
	return C.GoString(C.cupsLastErrorString())
}
