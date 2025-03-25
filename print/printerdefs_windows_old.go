package print

/*
import "syscall"

// printerAccessRights is the type for printer access rights. Imagine that!
type printerAccessRights uint32

const (
	// printerAccessAdminister allows performing administration tasks such as
	// those provided by SetPrinter.
	printerAccessAdminister printerAccessRights = 4
	// printerAccessUse allows performing basic printer functions.
	printerAccessUse printerAccessRights = 8
)

// printerDefs specifies the default data type, environment, initialization data,
// and access rights for a printer.
type printerDefs struct {
	defType *uint16
	devMode *PrinterDevMode
	access  printerAccessRights
}

// newPrinterDefaults creates a printerDefs structure based on the input values.
//
// Params:
//	dType specifies the default data type for the printer. This is often "RAW".
//	dm is a pointer to a printerDevMode struct that identifies the default environment
// and initialization data for the printer.
//	access specifies the access rights for the printer.
func newPrinterDefaults(dType string,
	dm *PrinterDevMode,
	access printerAccessRights) *printerDefs {
	dt, _ := syscall.UTF16FromString(dType)
	return &printerDefs{defType: &dt[0], devMode: dm, access: access}
}
*/
