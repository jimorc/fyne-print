package print

//#define UNICODE
//#include "windows.h"
import "C"
import (
	"fmt"
	"strings"
	"syscall"
)

// Printer is a struct that allows access to a printer.
type Printer struct {
	pi2    PrinterInfo2
	handle syscall.Handle
}

// newPrinter creates a Printer struct based on information provided in the PrinterInfo2 argument.
func newPrinter(pInfo2 *PrinterInfo2) *Printer {
	p := &Printer{pi2: *pInfo2}
	printerDefs := newPrinterDefaults("RAW", pInfo2.DevMode(),
		C.PRINTER_ACCESS_USE)

	prHandle := openPrinter(p.pi2.PrinterName(), printerDefs)
	p.handle = prHandle
	return p
}

// String returns a string representation of the Printer struct.
func (pr *Printer) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("    Handle: %d\n", pr.handle))
	s.WriteString(prepend("    ", pr.pi2.String()))
	return s.String()
}
