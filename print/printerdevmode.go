//go:build windows

package print

import (
	"fmt"
	"syscall"
)

// PrinterDevMode specifies the characteristics of the Win32 printer devices.
// See https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew for the meaning
// of each field.
type PrinterDevMode struct {
	dmDeviceName       [CCHDEVICENAME]uint16
	dmSpecVersion      uint16
	dmDriverVersion    uint16
	dmSize             uint16
	dmDriverExtra      uint16
	dmFields           uint32
	dmOrientation      int16
	dmPaperSize        int16
	dmPaperLength      int16
	dmPaperWidth       int16
	dmScale            int16
	dmCopies           int16
	dmDefaultSource    int16
	dmPrintQuality     int16
	dmColor            int16
	dmDuplex           int16
	dmYResolution      int16
	dmTTOption         int16
	dmCollate          int16
	dmFormName         [CCHFORMNAME]uint16
	dmLogPixels        uint16
	dmBitsPerPel       uint32
	dmPelsWidth        uint32
	dmPelsHeight       uint32
	dmNup              uint32
	dmDisplayFrequency uint32
	dmICMMethod        uint32
	dmICMIntent        uint32
	dmMediaType        uint32
	dmDitherType       uint32
	dmReserved1        uint32
	dmReserved2        uint32
	dmPanningWidth     uint32
	dmPanningHeight    uint32
}

// Copies retrieves the number of copies specified.
func (m *PrinterDevMode) Copies() int16 {
	return m.dmCopies
}

// DeviceName retrieves the name of the printer device.
func (m *PrinterDevMode) DeviceName() string {
	return syscall.UTF16ToString(m.dmDeviceName[:])
}

// FormName specifies the name of the form (e.g. "Letter" or "Legal") to use.
// This must be a name that is retrieved from the Printer.EnumForms method
func (m *PrinterDevMode) FormName() string {
	return syscall.UTF16ToString(m.dmFormName[:])
}

// Print prints out the values of the PrinterDevMode struct.
func (m *PrinterDevMode) Print() {
	fmt.Println("PrinterDevMode:")
	fmt.Printf("    Device Name: %s\n", m.DeviceName())
	fmt.Printf("    SpecVersion: %d\n", m.dmSpecVersion)
	fmt.Printf("    Driver Version: %d\n", m.dmDriverVersion)
	fmt.Printf("    Size: %d\n", m.dmSize)
	fmt.Printf("    Driver Extra: %d\n", m.dmDriverExtra)
	fmt.Printf("    Fields: %d\n", m.dmFields)
	fmt.Printf("    Orientation: %d\n", m.dmOrientation)
	fmt.Printf("    PaperSize: %d\n", m.dmPaperSize)
	fmt.Printf("    Paper Length: %d\n", m.dmPaperLength)
	fmt.Printf("    Paper Width: %d\n", m.dmPaperWidth)
	fmt.Printf("    Scale: %d\n", m.dmScale)
	fmt.Printf("    Copies: %d\n", m.dmCopies)
	fmt.Printf("    Default Source: %d\n", m.dmDefaultSource)
	fmt.Printf("    Print Quality: %d\n", m.dmPrintQuality)
	fmt.Printf("    Color: %d\n", m.dmColor)
	fmt.Printf("    Duplex: %d\n", m.dmDuplex)
	fmt.Printf("    Y-Resolution: %d\n", m.dmYResolution)
	fmt.Printf("    TT Option: %d\n", m.dmTTOption)
	fmt.Printf("    Collate: %d\n", m.dmCollate)
	fmt.Printf("    Form Name: %s\n", m.FormName())
	fmt.Printf("    Logical Pixels: %d\n", m.dmLogPixels)
	fmt.Printf("    Bits Per Pel: %d\n", m.dmBitsPerPel)
	fmt.Printf("    Pels Width: %d\n", m.dmPelsWidth)
	fmt.Printf("    Pels Height: %d\n", m.dmPelsHeight)
	fmt.Printf("    Where NUP is done: %d\n", m.dmNup)
	fmt.Printf("    Display Frequency: %d\n", m.dmDisplayFrequency)
	fmt.Printf("    ICM Method: %d\n", m.dmICMMethod)
	fmt.Printf("    ICM Intent: %d\n", m.dmICMIntent)
	fmt.Printf("    Media Type: %d\n", m.dmMediaType)
	fmt.Printf("    Dither Type: %d\n", m.dmDitherType)
	fmt.Printf("    Panning Width: %d\n", m.dmPanningWidth)
	fmt.Printf("    Panning Height: %d\n", m.dmPanningHeight)
}

// Copy performs a deep copy of the PrinterDevMode struct.
func CopyDM(src *PrinterDevMode) *PrinterDevMode {
	dst := &PrinterDevMode{
		dmSpecVersion:      src.dmSpecVersion,
		dmDriverVersion:    src.dmDriverVersion,
		dmSize:             src.dmSize,
		dmDriverExtra:      src.dmDriverExtra,
		dmFields:           src.dmFields,
		dmOrientation:      src.dmOrientation,
		dmPaperSize:        src.dmPaperSize,
		dmPaperLength:      src.dmPaperLength,
		dmPaperWidth:       src.dmPaperWidth,
		dmScale:            src.dmScale,
		dmCopies:           src.dmCopies,
		dmDefaultSource:    src.dmDefaultSource,
		dmPrintQuality:     src.dmPrintQuality,
		dmColor:            src.dmColor,
		dmDuplex:           src.dmDuplex,
		dmYResolution:      src.dmYResolution,
		dmTTOption:         src.dmTTOption,
		dmCollate:          src.dmCollate,
		dmLogPixels:        src.dmLogPixels,
		dmBitsPerPel:       src.dmBitsPerPel,
		dmPelsWidth:        src.dmPelsWidth,
		dmPelsHeight:       src.dmPelsHeight,
		dmNup:              src.dmNup,
		dmDisplayFrequency: src.dmDisplayFrequency,
		dmICMMethod:        src.dmICMMethod,
		dmICMIntent:        src.dmICMIntent,
		dmMediaType:        src.dmMediaType,
		dmDitherType:       src.dmDitherType,
		dmReserved1:        src.dmReserved1,
		dmReserved2:        src.dmReserved2,
		dmPanningWidth:     src.dmPanningWidth,
		dmPanningHeight:    src.dmPanningHeight,
	}
	copy(dst.dmDeviceName[:], src.dmDeviceName[:])
	copy(dst.dmFormName[:], src.dmFormName[:])

	return dst
}
