//go:build windows

package print

import (
	"fmt"
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
)

/*
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
*/
var (
	modwinspool = syscall.NewLazyDLL("winspool.drv")

	procClosePrinter       = modwinspool.NewProc("ClosePrinter")
	procDeviceCapabilities = modwinspool.NewProc("DeviceCapabilitiesW")
	procEnumForms          = modwinspool.NewProc("EnumFormsW")
	procEnumPrinters       = modwinspool.NewProc("EnumPrintersW")
	procGetDefaultPrinter  = modwinspool.NewProc("GetDefaultPrinterW")
	procOpenPrinter        = modwinspool.NewProc("OpenPrinterW")
)

// closePrinter closes the printer.
//
// Params:
//
//	printerHandle is the handle of the printer to close.
func closePrinter(printerHandle syscall.Handle) {
	r1, _, err := procClosePrinter.Call(uintptr(printerHandle))
	if r1 == 0 {
		fyne.LogError("Failed to close printer", err)
	}
}

/*
// devicvCapabilities retrieves the capabilities of a printer driver.
// See https://learn.microsoft.com/en-us/windows/win32/printdocs/documentproperties
// for information on the arguments.
//
// The return value depends on the the capability being requested.
// A return value of zero generally indicates that, while the function completed successfully,
// there was some type of failure, such as a capability that is not supported. For more details,
// see the descriptions for the fwCapability values.
//
// If the function returns -1, this may mean either that the capability is not supported or
// there was a general function failure.
func deviceCapabilities(name string,

		port string,
		capability devCapIndex,
		output uintptr,
		devMode *PrinterDevMode) (int32, error) {
		n, _ := syscall.UTF16FromString(name)
		p, _ := syscall.UTF16FromString(port)
		r1, _, err := procDeviceCapabilities.Call(
			uintptr(unsafe.Pointer(&n[0])),
			uintptr(unsafe.Pointer(&p[0])),
			uintptr(capability),
			output,
			uintptr(unsafe.Pointer(devMode)))
		return int32(r1), err
	}
*/
// enumPrinters enumerates available printers, print servers, domains, or print providers.
// See https://learn.microsoft.com/en-us/windows/win32/printdocs/enumprinters for information
// on the arguments.
//
// Returns:
//
//	A bool indicating if the function succeeded.
//	An error if the function failed.
func enumPrinters(flags uint32,
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

/*
// getDefaultPrinter returns default printer information.

	func getDefaultPrinter(buf *uint16, bufN *uint32) error {
		_, _, err := procGetDefaultPrinter.Call(
			uintptr(unsafe.Pointer(buf)),
			uintptr(unsafe.Pointer(bufN)))
		return err
	}
*/
func openPrinter(pName string, printerDefs *PrinterDefaults) syscall.Handle {
	name, _ := syscall.UTF16FromString(pName)
	var prHandle syscall.Handle
	r0, _, err := procOpenPrinter.Call(
		uintptr(unsafe.Pointer(&name[0])),
		uintptr(unsafe.Pointer(&prHandle)),
		uintptr(unsafe.Pointer(printerDefs)))
	if r0 == 0 {
		eMsg := fmt.Sprintf("Failed to open printer %s", pName)
		fyne.LogError(eMsg, err)
		return syscall.Handle(0)
	}
	return prHandle
}

/*
enumForms returns the formInfo2 objects for a printer. This is the media size,
and localizable name of each form.
See https://learn.microsoft.com/en-us/windows/win32/printdocs/form-info-2 for information
on the arguments.
This function should be called twice, the first time with a zero or small buffer length
to determine the required buffer size, and the second time to retrieve all of the
formInfo2 objects. The retrieved formInfo2 slice should then be shortened to the
actual number of objects retrieved. This is necessary because the first call returns
a 'needed' size that is much larger than is actually needed.

Returns:
  - a slice of formInfo2 objects, or nil on error.
  - an error object if an error other than syscall.ERROR_INSUFFICIENT_BUFFER.
*/
func enumForms(printerHandle syscall.Handle, level uint32, forms []formInfo2,
	cbBuf uint32, needed *uint32, cReturned *uint32) ([]formInfo2, error) {
	r1, _, err := procEnumForms.Call(
		uintptr(printerHandle),
		uintptr(level),
		uintptr(unsafe.Pointer(&forms[0])),
		uintptr(cbBuf),
		uintptr(unsafe.Pointer(needed)),
		uintptr(unsafe.Pointer(cReturned)))
	if r1 == 0 && err != syscall.ERROR_INSUFFICIENT_BUFFER {
		fyne.LogError("Failed to enumerate forms", err)
		return nil, err
	}
	return forms, nil
}
