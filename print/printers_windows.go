package print

//#define UNICODE
//#include "windows.h"
import "C"

import (
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
)

// Printers contains a slice of Printer objects. This is normally all of the
// printers available on the computer.
type Printers []Printer

func NewPrinters() *Printers {
	var flags uint32 = C.PRINTER_ENUM_LOCAL |
		C.PRINTER_ENUM_CONNECTIONS
		// buffer is a slice of bytes that will contain an array of PrinterInfo1 structs
	var buffer = make([]byte, 1)
	// size of *info2 in bytes
	var info2Size uint32 = 0
	// number of bytes needed in *info1
	var info2Needed uint32 = 0
	// number of info1 structs returned.
	var info2Count uint32 = 0

	_, err := enumPrinters(flags,
		"",
		2,
		&buffer[0],
		info2Size,
		&info2Needed,
		&info2Count)

	if err != syscall.ERROR_INSUFFICIENT_BUFFER {
		fyne.LogError("Error getting printers", err)
		return nil
	}
	info2Size = info2Needed
	buffer = make([]byte, info2Size)
	_, err = enumPrinters(flags,
		"",
		2,
		&buffer[0],
		info2Size,
		&info2Needed,
		&info2Count)
	if err != syscall.Errno(0) {
		fyne.LogError("Error getting printers", err)
		return nil
	}
	pInfo2 := (*[1024]PrinterInfo2)(unsafe.Pointer(&buffer[0]))[:info2Count:info2Count]
	printers := &Printers{}
	for _, info2 := range pInfo2 {
		p := newPrinter(&info2)
		printers.Add(p)
	}
	return printers
}

// Add adds a printer to the Printers struct.
func (p *Printers) Add(printer *Printer) {
	*p = append(*p, *printer)
}

// ClosePrinters calls Printer.close for all printers.
func (p *Printers) ClosePrinters() {
	for _, printer := range *p {
		printer.close()
	}
}
