//go:build windows

package print

import (
	"fmt"
	"strings"
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
	dmPaperSize        dmPaperSize
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

// string returns the PrinterDevMode struct as a string.
func (m *PrinterDevMode) string() string {
	var s strings.Builder
	s.WriteString("PrinterDevMode:\n")
	s.WriteString(fmt.Sprintf("    Device Name: %s\n", m.DeviceName()))
	s.WriteString(fmt.Sprintf("    SpecVersion: %d\n", m.dmSpecVersion))
	s.WriteString(fmt.Sprintf("    Driver Version: %d\n", m.dmDriverVersion))
	s.WriteString(fmt.Sprintf("    Size: %d\n", m.dmSize))
	s.WriteString(fmt.Sprintf("    Driver Extra: %d\n", m.dmDriverExtra))
	s.WriteString(fmt.Sprintf("    Fields: %d\n", m.dmFields))
	s.WriteString(fmt.Sprintf("    Orientation: %d\n", m.dmOrientation))
	s.WriteString(fmt.Sprintf("    PaperSize: %d\n", m.dmPaperSize))
	s.WriteString(fmt.Sprintf("    Paper Length: %d\n", m.dmPaperLength))
	s.WriteString(fmt.Sprintf("    Paper Width: %d\n", m.dmPaperWidth))
	s.WriteString(fmt.Sprintf("    Scale: %d\n", m.dmScale))
	s.WriteString(fmt.Sprintf("    Copies: %d\n", m.dmCopies))
	s.WriteString(fmt.Sprintf("    Default Source: %d\n", m.dmDefaultSource))
	s.WriteString(fmt.Sprintf("    Print Quality: %d\n", m.dmPrintQuality))
	s.WriteString(fmt.Sprintf("    Color: %d\n", m.dmColor))
	s.WriteString(fmt.Sprintf("    Duplex: %d\n", m.dmDuplex))
	s.WriteString(fmt.Sprintf("    Y-Resolution: %d\n", m.dmYResolution))
	s.WriteString(fmt.Sprintf("    TT Option: %d\n", m.dmTTOption))
	s.WriteString(fmt.Sprintf("    Collate: %d\n", m.dmCollate))
	s.WriteString(fmt.Sprintf("    Form Name: %s\n", m.FormName()))
	s.WriteString(fmt.Sprintf("    Logical Pixels: %d\n", m.dmLogPixels))
	s.WriteString(fmt.Sprintf("    Bits Per Pel: %d\n", m.dmBitsPerPel))
	s.WriteString(fmt.Sprintf("    Pels Width: %d\n", m.dmPelsWidth))
	s.WriteString(fmt.Sprintf("    Pels Height: %d\n", m.dmPelsHeight))
	s.WriteString(fmt.Sprintf("    Where NUP is done: %d\n", m.dmNup))
	s.WriteString(fmt.Sprintf("    Display Frequency: %d\n", m.dmDisplayFrequency))
	s.WriteString(fmt.Sprintf("    ICM Method: %d\n", m.dmICMMethod))
	s.WriteString(fmt.Sprintf("    ICM Intent: %d\n", m.dmICMIntent))
	s.WriteString(fmt.Sprintf("    Media Type: %d\n", m.dmMediaType))
	s.WriteString(fmt.Sprintf("    Dither Type: %d\n", m.dmDitherType))
	s.WriteString(fmt.Sprintf("    Panning Width: %d\n", m.dmPanningWidth))
	s.WriteString(fmt.Sprintf("    Panning Height: %d\n", m.dmPanningHeight))

	return s.String()
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
