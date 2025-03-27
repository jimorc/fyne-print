package print

//#define UNICODE
//#include "windows.h"
import "C"

import (
	"fmt"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

// devMode is the Win32 DEVMODEW struct
type devMode C.DEVMODEW

//The following methods return the DEVMODEW fields with the "dm" prefix removed.

// DeviceName returns the printer's device name.
func (d *devMode) DeviceName() string {
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(&d.dmDeviceName[0])))
}

// SpecVersion returns the DEVMODE version number.
func (d *devMode) SpecVersion() uint16 {
	return uint16(d.dmSpecVersion)
}

// DriverVersion returns the printer's driver version number assigned by the
// printer driver developer.
func (d *devMode) DriverVersion() uint16 {
	return uint16(d.dmDriverVersion)
}

// Size returns the size in bytes of the public DEVMODEW struct, not including
// any private, driver-specified members identified by the DriverExtra member.
func (d *devMode) Size() uint16 {
	return uint16(d.dmSize)
}

// DriverExtra returns the number of bytes of private driver data that follow the
// public structure members.
func (d *devMode) DriverExtra() uint16 {
	return uint16(d.dmDriverExtra)
}

// Fields specifies bit flags identifying which of the following DEVMODEW members
// are in use. For example, the DM_ORIENTATION flag is set when the dmOrientation
// member contains valid data. The DM_XXX flags are defined in wingdi.h.
func (d *devMode) Fields() uint32 {
	return uint32(d.dmFields)
}

// Orientation returns the printer's paper orientation. Valid values are
// C.DMORIENTPORTRAIT and C.DMORIENT_LANDSCAPE.
func (d *devMode) Orientation() uint16 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[1 << 30]uint16)(p)[0:2]
	return pSlice[0]
}

// PaperSize returns the printer's media size as one of the DMPAPER_xxx
// values defined in wingdi.h. If PaperSize returns 0, then the media size
// is specified by the PaperLength and PaperWidth methods.
func (d *devMode) PaperSize() uint16 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[1 << 30]uint16)(p)[2:4]
	return pSlice[0]
}

// PaperLength returns the printer's media size length in 1/10ths of a mm. This
// value is valid only if PaperSize return 0.
func (d *devMode) PaperLength() uint16 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[1 << 30]uint16)(p)[4:6]
	return pSlice[0]
}

// PaperWidth returns the printer's media size width in 1/10ths of a mm. This
// value is valid only if PaperSize return 0.
func (d *devMode) PaperWidth() uint16 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[1 << 30]uint16)(p)[6:8]
	return pSlice[0]
}

// Scale returns the percentage by which the image is to be scaled for printing.
// The image's media size is scaled to the physical page by the factor of Scale()/100.
func (d *devMode) Scale() uint16 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[1 << 30]uint16)(p)[8:10]
	return pSlice[0]
}

// Copies returns the number of copies to be printed if the driver supports
// multiple copies
func (d *devMode) Copies() uint16 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[1 << 30]uint16)(p)[10:12]
	return pSlice[0]
}

// DefaultSource returns the printer's default input bin. This must be one of
// // the DMBIN-prefixed constants defined in wingdi.h. If the specified constant
// is DMBIN_FORMSOURCE, the input bin should be selected automatically.
func (d *devMode) DefaultSource() uint16 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[1 << 30]uint16)(p)[12:14]
	return pSlice[0]
}

// PrintQuality returns the printer resolution. The following negative constant
// values are defined in wingdi.h:
//
// DMRES_HIGH
// DMRES_MEDIUM
// DMRES_LOW
// DMRES_DRAFT
// If a positive value is specified, it represents the number of dots per inch (DPI)
// for the x resolution, and the y resolution is specified by YResolution.
func (d *devMode) PrintQuality() uint16 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[1 << 30]uint16)(p)[14:16]
	return pSlice[0]
}

// Color indicates whether a color printer should print in color or monochrome mode.
// The returned value is one of DMCOLOR_COLOR or DMCOLOR_MONOCHROME.
func (d *devMode) Color() uint16 {
	return uint16(d.dmColor)
}

// Duplex specifies duplex (double-sided) printing for duplex-capable printers.
// The returned value is one of DMDUP_HORIZONTAL, DM_DUP_SIMPLEX, and DMDUP_VERTICAL.
func (d *devMode) Duplex() uint16 {
	return uint16(d.dmDuplex)
}

// YResolution returns the y resolution of the printer in DPI. This value is valid only
// if PrintQuality() returns a positive value.
func (d *devMode) YResolution() uint16 {
	return uint16(d.dmYResolution)
}

// TTOPtion specifies how TrueType fonts should be printed. The returned value will be
// one of the DMTT_ prefixed values defined in wingdi.h.
func (d *devMode) TTOption() uint16 {
	return uint16(d.dmTTOption)
}

// Collate indicates whether multiple prints should be collated. Valid values are
// DMCOLLATE_TRUE and DM_COLLATE_FALSE.
func (d *devMode) Collate() uint16 {
	return uint16(d.dmCollate)
}

// FormName returns the name of the form (media size) to use. This is a name
// that is returned by the Printer.EnumForms method.
func (d *devMode) FormName() string {
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(&d.dmFormName[0])))
}

// Nup specifies whether the printer handles N-up printing (printing multiple pages
// on a single physical page). The returned value is one of the following:
// DMNUP_SYSTEM: The print system handles N-up printing.
// DMNUP_ONEUP: The print system does not handle "N-up" printing. An application can
// set dmNup to DMNUP_ONEUP if it intends to carry out "N-up" printing on its own.
func (d *devMode) Nup() uint32 {
	p := unsafe.Pointer(&d.anon1[0])
	pSlice := (*[1 << 30]uint32)(p)[0:4]
	return pSlice[0]
}

// ICMMethod specifies how ICM processing can.should be performed. Valid values and
// their meanings are a follows:
// DMICMMETHOD_NONE: ICM is disabled.
// DMICMMETHOD_SYSTEM: ICM is handled by the print system.
// DMICMMETHOD_DRIVER: ICM is handled by the printer driver.
// DMICMMETHOD_DEVICE: ICM is handled by the printer device.
func (d *devMode) ICMMethod() uint32 {
	return uint32(d.dmICMMethod)
}

// ICMIntent specifies the ICM Intent. Valud values and meanings are as follows:
// DMICM_SATURATE: Maximize color saturation.
// DMICM_CONTRAST: Maximize color contrast.
// DMICM_COLORIMETRIC: Use specific color metric.
// DMICM_ABS_COLORIMETRIC: Use specific color metric.
func (d *devMode) ICMIntent() uint32 {
	return uint32(d.dmICMIntent)
}

// MediaType returns the media type. Valid values and meanings are as follows:
// DMMEDIA_STANDARD: Standard paper.
// DMMEDIA_TRANSPARENCY: Transparency.
// DMMEDIA_GLOSSY: Glossy paper.
func (d *devMode) MediaType() uint32 {
	return uint32(d.dmMediaType)
}

// DitherType specifies how dithering is performed. Valid values and meanings
// are as follows:
// DMDITHER_NONE: No dithering.
// DMDITHER_COARSE: Dither with a coarse brush.
// DMDITHER_FINE: Dither with a fine brush.
// DMDITHER_LINEART: LineArt dithering.
// DMDITHER_ERRORDIFFUSION: LineArt dithering.
// DMDITHER_RESERVED6: LineArt dithering.
// DMDITHER_RESERVED7: LineArt dithering.
// DMDITHER_RESERVED8: LineArt dithering.
// DMDITHER_RESERVED9: LineArt dithering.
// DMDITHER_GRAYSCALE: Device does grayscaling.
func (d *devMode) DitherType() uint32 {
	return uint32(d.dmDitherType)
}

func (d *devMode) String() string {
	var s strings.Builder
	s.WriteString("DevMode:\n")
	s.WriteString(fmt.Sprintf("    Device Name: %s\n", d.DeviceName()))
	s.WriteString(fmt.Sprintf("    Spec Version: %d\n", d.SpecVersion()))
	s.WriteString(fmt.Sprintf("    Driver Version: %d\n", d.DriverVersion()))
	s.WriteString(fmt.Sprintf("    Size: %d\n", d.Size()))
	s.WriteString(fmt.Sprintf("    Driver Extra: %d\n", d.DriverExtra()))
	s.WriteString(fmt.Sprintf("    Fields: %d\n", d.Fields()))
	s.WriteString(fmt.Sprintf("    Orientation: %d\n", d.Orientation()))
	s.WriteString(fmt.Sprintf("    Paper Size: %d\n", d.PaperSize()))
	s.WriteString(fmt.Sprintf("    Paper Length: %d\n", d.PaperLength()))
	s.WriteString(fmt.Sprintf("    Paper Width: %d\n", d.PaperWidth()))
	s.WriteString(fmt.Sprintf("    Scale: %d\n", d.Scale()))
	s.WriteString(fmt.Sprintf("    Copies: %d\n", d.Copies()))
	s.WriteString(fmt.Sprintf("    Default Source: %d\n", d.DefaultSource()))
	s.WriteString(fmt.Sprintf("    Print Quality: %d\n", d.PrintQuality()))
	s.WriteString(fmt.Sprintf("    Color: %d\n", d.Color()))
	s.WriteString(fmt.Sprintf("    Duplex: %d\n", d.Duplex()))
	s.WriteString(fmt.Sprintf("    Y Resolution: %d\n", d.YResolution()))
	s.WriteString(fmt.Sprintf("    TT Option: %d\n", d.TTOption()))
	s.WriteString(fmt.Sprintf("    Collate: %d\n", d.Collate()))
	s.WriteString(fmt.Sprintf("    Form Name: %s\n", d.FormName()))
	s.WriteString(fmt.Sprintf("    Nup: %d\n", d.Nup()))
	s.WriteString(fmt.Sprintf("    ICM Method: %d\n", d.ICMMethod()))
	s.WriteString(fmt.Sprintf("    ICM Intent: %d\n", d.ICMIntent()))
	s.WriteString(fmt.Sprintf("    Media Type: %d\n", d.MediaType()))
	s.WriteString(fmt.Sprintf("    Dither Type: %d\n", d.DitherType()))
	return s.String()
}
