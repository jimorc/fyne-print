package print

//#define UNICODE
//#include "windows.h"
import "C"
import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// PrinterDefault specifies the default data type, environment, initialization data,
// and access rights for a printer.
// A PrinterDefault object is passed into the OpenPrinter function to define the access
// rights for the printer.
type PrinterDefaults C.PRINTER_DEFAULTSW

func newPrinterDefaults(dType string,
	dm *devMode,
	access C.ACCESS_MASK) *PrinterDefaults {
	dt, _ := syscall.UTF16FromString(dType)
	return &PrinterDefaults{pDatatype: C.LPWSTR(unsafe.Pointer(&dt[0])),
		pDevMode: C.LPDEVMODEW(unsafe.Pointer(dm)), DesiredAccess: access}
}

// )
// Datatype returns the default data type for the printer.
func (pd *PrinterDefaults) Datatype() string {
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pd.pDatatype)))
}

// DevMode returns the printer's devMode object.
func (pd *PrinterDefaults) DevMode() *devMode {
	return (*devMode)(unsafe.Pointer(pd.pDevMode))
}

// Access returns the desired access rights for the printer as a string slice.
// The OpenPrinter function uses the access member to set access rights to the printer.
// These rights can affect the operation of the SetPrinter and DeletePrinter
// functions. The access rights can be one of the following:
//
// PRINTER_ACCESS_ADMINISTER: To perform administrative tasks, such as those provided
// by SetPrinter.
// PRINTER_ACCESS_USE: To perform basic printing operations.
func (pd *PrinterDefaults) Access() []string {
	access := uint32(pd.DesiredAccess)
	s := []string{}
	if access&C.PRINTER_ACCESS_ADMINISTER != 0 {
		s = append(s, "Administer Access")
	}
	if access&C.PRINTER_ACCESS_USE != 0 {
		s = append(s, "Use Access")
	}
	return s
}

// String returns a string representation of the PrinterDefaults struct.
func (pd *PrinterDefaults) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintln("PrinterDefaults:"))
	s.WriteString(fmt.Sprintf("    Datatype: %s\n", pd.Datatype()))
	s.WriteString(prepend("    ", pd.DevMode().String()))
	s.WriteString("    Desired Access:\n")
	for _, a := range pd.Access() {
		s.WriteString(fmt.Sprintf("        %s\n", a))
	}

	return s.String()
}
