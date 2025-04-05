//go:build windows

package print

import (
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
	"golang.org/x/sys/windows"
)

// intPaperSize is the paper size as integer values. When retrieved using deviceCapabilities,
// the size is in tenths of a millimeter.
type intPaperSize struct {
	w int32
	h int32
}

const (
	maxNumPaperSizes = 256 // more than enough for all paper sizes returned from a Windows printer driver
	paperNameSize    = 64  // maximum size of a paper name.
)

// devCapIndex define the capabilities that may be queried for a printer
type devCapIndex uint16

const (
	dcFields           devCapIndex = 1  // DC_FIELDS
	dcPapers           devCapIndex = 2  // DC_PAPERS
	dcPaperSize        devCapIndex = 3  // DC_PAPERSIZE
	dcMinExtent        devCapIndex = 4  // DC_MINEXTENT
	dcMaxExtent        devCapIndex = 5  // DC_MAXEXTENT
	dcBins             devCapIndex = 6  // DC_BINS
	dcDuplex           devCapIndex = 7  // DC_DUPLEX
	dcSize             devCapIndex = 8  // DC_SIZE
	dcExtra            devCapIndex = 9  // DC_EXTRA
	dcVersion          devCapIndex = 10 // DC_VERSION
	dcDriver           devCapIndex = 11 // DC_DRIVER
	dcBinNames         devCapIndex = 12 // DC_BINNAMES
	dcEnumResolutions  devCapIndex = 13 // DC_ENUMRESOLUTIONS
	dcFileDependencies devCapIndex = 14 // DC_FILEDEPENDENCIES
	dcTruetype         devCapIndex = 15 // DC_TRUETYPE
	dcPaperNames       devCapIndex = 16 // DC_PAPERNAMES
	dcOrientation      devCapIndex = 17 // DC_ORIENTATION
	dcCopies           devCapIndex = 18 // DC_COPIES
	dcBinAdjust        devCapIndex = 19 // DC_BINADJUST
	dcEmfCompliant     devCapIndex = 20 // DC_EMF_COMPLIANT
	dcDatatypeProduced devCapIndex = 21 // DC_DATATYPE_PRODUCED
	dcCollate          devCapIndex = 22 // DC_COLLATE
	dcManufacturer     devCapIndex = 23 // DC_MANUFACTURER
	dcModel            devCapIndex = 24 // DC_MODEL
	dcPersonality      devCapIndex = 25 // DC_PERSONALITY
	dcPrintRate        devCapIndex = 26 // DC_PRINTRATE
	dcPrintRateUnit    devCapIndex = 27 // DC_PRINTRATEUNIT
	dcPrinterMem       devCapIndex = 28 // DC_PRINTERMEM
	dcMediaReady       devCapIndex = 29 // DC_MEDIAREADY
	dcStaple           devCapIndex = 30 // DC_STAPLE
	dcPrintRatePPM     devCapIndex = 31 // DC_PRINTRATEPPM
	dcColorDevice      devCapIndex = 32 // DC_COLORDEVICE
	dcNup              devCapIndex = 33 // DC_NUP
	dcMediaTypes       devCapIndex = 34 // DC_MEDIATYPES
	dcMediaTypeNames   devCapIndex = 35 // DC_MEDIATYPENAMES
)

// printRateUnit defines the unit for the value returned by querying
// the dcPrintRate.
type printRateUnit int32

const (
	pruPPM printRateUnit = 1 // PRINTRATEUNIT_PPM
	pruCPS printRateUnit = 2 // PRINTRATEUNIT_CPS
	pruLPM printRateUnit = 3 // PRINTRATEUNIT_LPM
	pruIPM printRateUnit = 4 // PRINTRATEUNIT_IPM
)

var (
	modgdi32 = syscall.NewLazyDLL("gdi32.dll")

	procCreateDC      = modgdi32.NewProc("CreateDCW")
	procDeleteDC      = modgdi32.NewProc("DeleteDC")
	procGetDeviceCaps = modgdi32.NewProc("GetDeviceCaps")
)

// createDC creates a device context for the named printer.
//
// Params:
//
//	prName is the name of the printer as held in the printer's PrinterInfo2 struct.
//
// Returns the printer's device context, or 0 on error.
func createDC(prName string) syscall.Handle {
	n, _ := windows.UTF16FromString(prName)
	r1, _, err := procCreateDC.Call(0, uintptr(unsafe.Pointer(&n[0])), 0, 0)
	if r1 == 0 {
		fyne.LogError("Error creating printer DC: ", err)
	}
	return syscall.Handle(r1)
}

// deleteDC deletes a device context. The dc should have been created using
// createDC. Do not call deleteDC for a dc that was created using the GetDC function.
//
// Params:
//
//	dc is the device context to delete.
func deleteDC(dc syscall.Handle) {
	procDeleteDC.Call(uintptr(dc))
}

func getDeviceCaps(dc syscall.Handle, item int32) int32 {
	r1, _, _ := procGetDeviceCaps.Call(uintptr(dc), uintptr(item))
	return int32(r1)
}
