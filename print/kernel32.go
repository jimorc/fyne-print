//go:build windows

package print

//#define UNICODE
//#include "windows.h"
import "C"

import (
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procLCIDToLocaleName = modkernel32.NewProc("LCIDToLocaleName")
)

// LCIDToLocaleName retrieves the locale name (e.g. 'en-US') for the specified
// locale ID.
func LCIDToLocaleName(locale C.LCID, name *uint16, size int, flags uint32) int {
	r1, _, err := procLCIDToLocaleName.Call(
		uintptr(locale),
		uintptr(unsafe.Pointer(name)),
		uintptr(size),
		uintptr(flags))
	if r1 == 0 {
		fyne.LogError("LCIDToLocaleName failed", err)
	}
	return int(r1)
}
