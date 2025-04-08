package print

import (
	"strings"
	"unsafe"
)

// printerInfo8 specifies the printer's devMode information.
type printerInfo8 struct {
	pDevMode *devMode
}

// DevMode returns the printer's devMode object.
func (pi8 *printerInfo8) DevMode() *devMode {
	return (*devMode)(unsafe.Pointer(pi8.pDevMode))
}

// String returns a string representation of the printerInfo8 object.
func (pi8 *printerInfo8) String() string {
	var s strings.Builder
	s.WriteString("printerInfo8:\n")
	s.WriteString(pi8.DevMode().String())
	return s.String()
}
