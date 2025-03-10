//go:build windows

package print

import (
	"syscall"
	"unsafe"
)

// Printer enumeration constants. One of more of these may be used to enumerate printers configured on the system.
const (
	PRINTER_ENUM_DEFAULT     uint32 = 0x1
	PRINTER_ENUM_LOCAL       uint32 = 0x2
	PRINTER_ENUM_CONNECTIONS uint32 = 0x4
	PRINTER_ENUM_FAVORITE    uint32 = 0x4
	PRINTER_ENUM_NAME        uint32 = 0x8
	PRINTER_ENUM_REMOTE      uint32 = 0x10
	PRINTER_ENUM_SHARED      uint32 = 0x20
	PRINTER_ENUM_NETWORK     uint32 = 0x40
)

var (
	modwinspool = syscall.NewLazyDLL("winspool.drv")

	procEnumPrinters      = modwinspool.NewProc("EnumPrintersW")
	procGetDefaultPrinter = modwinspool.NewProc("GetDefaultPrinterW")
)

// EnumPrinters enumerates available printers, print servers, domains, or print providers.
// See https://learn.microsoft.com/en-us/windows/win32/printdocs/enumprinters for information on the arguments.
//
// Returns:
//
//	A bool indicating if the function succeeded.
//	An error if the function failed.
func EnumPrinters(flags uint32,
	name string,
	level uint32,
	buf *byte,
	cbBuf uint32,
	needed *uint32,
	cReturned *uint32) (bool, error) {
	n, _ := syscall.UTF16FromString(name)
	r1, _, err := procEnumPrinters.Call(
		uintptr(flags),
		uintptr(unsafe.Pointer(&n[0])),
		uintptr(level),
		uintptr(unsafe.Pointer(buf)),
		uintptr(cbBuf),
		uintptr(unsafe.Pointer(needed)),
		uintptr(unsafe.Pointer(cReturned)))
	return r1 != 0, err
}

// getDefaultPrinter returns default printer information.
func getDefaultPrinter(buf *uint16, bufN *uint32) error {
	_, _, err := procGetDefaultPrinter.Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(bufN)))
	return err
}
