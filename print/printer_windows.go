package print

//#define UNICODE
//#include "windows.h"
import "C"
import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
)

// Printer is a struct that allows access to a printer.
type Printer struct {
	pi2    PrinterInfo2
	handle syscall.Handle
	forms  []formInfo2
}

// newPrinter creates a Printer struct based on information provided in the PrinterInfo2 argument.
func newPrinter(pInfo2 *PrinterInfo2) *Printer {
	p := &Printer{pi2: *pInfo2}
	printerDefs := newPrinterDefaults("RAW", pInfo2.DevMode(),
		C.PRINTER_ACCESS_USE)

	prHandle := openPrinter(p.pi2.PrinterName(), printerDefs)
	p.handle = prHandle
	p.forms = make([]formInfo2, 1)
	p.getMediaSizes()
	return p
}

// String returns a string representation of the Printer struct.
func (pr *Printer) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("    Handle: %d\n", pr.handle))
	s.WriteString(prepend("    ", pr.pi2.String()))
	s.WriteString(fmt.Sprintf("    Printer has %d Media Sizes:\n", len(pr.forms)))
	s.WriteString("    Media Sizes:\n")
	for i, f := range pr.forms {
		s.WriteString(fmt.Sprintf("        Media Size %d:\n", i))
		s.WriteString(prepend("          ", f.String()))
	}
	return s.String()
}

// getMediaSizes retrieves the printer's formInfo2 objects. This is all of the
// media sizes that the printer might support.
func (p *Printer) getMediaSizes() {
	var cbBuf uint32 = uint32(unsafe.Sizeof(formInfo2{}))
	var needed uint32
	var returned uint32
	_, err := enumForms(p.handle, 2, p.forms, 0, &needed, &returned)
	if err != nil && err != syscall.ERROR_INSUFFICIENT_BUFFER {
		fyne.LogError("Error getting media sizes: ", err)
		return
	}
	p.forms = make([]formInfo2, needed/uint32(unsafe.Sizeof(formInfo2{})))
	cbBuf = needed
	_, err = enumForms(p.handle, 2, p.forms, cbBuf, &needed, &returned)
	if err != nil {
		fyne.LogError("Error getting media sizes: ", err)
	}
	// The 'needed' parameter requested many more objects than are actually returned.
	// Therefore, the slice of objects should be reduced to the number actually returned.
	p.forms = p.forms[:returned]
}

// close cleans up Printer-related data such as the printer handle.
func (p *Printer) close() {
	closePrinter(p.handle)
	p.handle = 0
}
