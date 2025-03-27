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
func (d *devMode) Fields() devModeFields {
	return devModeFields(d.dmFields)
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
	f := d.Fields()
	s.WriteString("DevMode:\n")
	s.WriteString(fmt.Sprintf("    Device Name: %s\n", d.DeviceName()))
	s.WriteString(fmt.Sprintf("    Spec Version: %d\n", d.SpecVersion()))
	s.WriteString(fmt.Sprintf("    Driver Version: %d\n", d.DriverVersion()))
	s.WriteString(fmt.Sprintf("    Size: %d bytes\n", d.Size()))
	s.WriteString(fmt.Sprintf("    Driver Extra: %d bytes\n", d.DriverExtra()))
	s.WriteString(prepend("    ", d.Fields().String()))
	if f.orientationSet() {
		s.WriteString(fmt.Sprintf("    Orientation: %d\n", d.Orientation()))
	}
	if f.paperSizeSet() {
		s.WriteString(fmt.Sprintf("    Paper Size: %d\n", d.PaperSize()))
	}
	if f.paperLengthSet() {
		s.WriteString(fmt.Sprintf("    Paper Length: %d\n", d.PaperLength()))
	}
	if f.paperWidthSet() {
		s.WriteString(fmt.Sprintf("    Paper Width: %d\n", d.PaperWidth()))
	}
	if f.scaleSet() {
		s.WriteString(fmt.Sprintf("    Scale: %d\n", d.Scale()))
	}
	if f.copiesSet() {
		s.WriteString(fmt.Sprintf("    Copies: %d\n", d.Copies()))
	}
	if f.defaultSourceSet() {
		s.WriteString(fmt.Sprintf("    Default Source: %d\n", d.DefaultSource()))
	}
	if f.printQualitySet() {
		s.WriteString(fmt.Sprintf("    Print Quality: %d\n", d.PrintQuality()))
	}
	if f.colorSet() {
		s.WriteString(fmt.Sprintf("    Color: %d\n", d.Color()))
	}
	if f.duplexSet() {
		s.WriteString(fmt.Sprintf("    Duplex: %d\n", d.Duplex()))
	}
	if f.yResolutionSet() {
		s.WriteString(fmt.Sprintf("    Y Resolution: %d\n", d.YResolution()))
	}
	if f.ttOptionSet() {
		s.WriteString(fmt.Sprintf("    TT Option: %d\n", d.TTOption()))
	}
	if f.collateSet() {
		s.WriteString(fmt.Sprintf("    Collate: %d\n", d.Collate()))
	}
	if f.formNameSet() {
		s.WriteString(fmt.Sprintf("    Form Name: %s\n", d.FormName()))
	}
	if f.nupSet() {
		s.WriteString(fmt.Sprintf("    Nup: %d\n", d.Nup()))
	}
	if f.icmMethodSet() {
		s.WriteString(fmt.Sprintf("    ICM Method: %d\n", d.ICMMethod()))
	}
	if f.icmIntentSet() {
		s.WriteString(fmt.Sprintf("    ICM Intent: %d\n", d.ICMIntent()))
	}
	if f.mediaTypeSet() {
		s.WriteString(fmt.Sprintf("    Media Type: %d\n", d.MediaType()))
	}
	if f.ditherTypeSet() {
		s.WriteString(fmt.Sprintf("    Dither Type: %d\n", d.DitherType()))
	}
	return s.String()
}

// devModeFields is the Fields value from the devMode object.
type devModeFields uint32

// orientationSet returns true if the DM_ORIENTATION bit is set.
func (f devModeFields) orientationSet() bool {
	return f&C.DM_ORIENTATION != 0
}

// paperSizeSet returns true if the DM_PAPERSIZE bit is set.
func (f devModeFields) paperSizeSet() bool {
	return f&C.DM_PAPERSIZE != 0
}

// paperLengthSet returns true if the DM_PAPERLENGTH bit is set.
func (f devModeFields) paperLengthSet() bool {
	return f&C.DM_PAPERLENGTH != 0
}

// paperWidthSet returns true if the DM_PAPERWIDTH bit is set.
func (f devModeFields) paperWidthSet() bool {
	return f&C.DM_PAPERWIDTH != 0
}

// scaleSet returns true if the DM_SCALE bit is set.
func (f devModeFields) scaleSet() bool {
	return f&C.DM_SCALE != 0
}

// positionSet returns true if the DM_POSITION bit is set.
func (f devModeFields) positionSet() bool {
	return f&C.DM_POSITION != 0
}

// nupSet returns true if the DM_NUP bit is set.
func (f devModeFields) nupSet() bool {
	return f&C.DM_NUP != 0
}

// displayOrientationSet returns true if the DM_DISPLAYORIENTATION bit is set.
func (f devModeFields) displayOrientationSet() bool {
	return f&C.DM_DISPLAYORIENTATION != 0
}

// copiesSet returns true if the DM_COPIES bit is set.
func (f devModeFields) copiesSet() bool {
	return f&C.DM_COPIES != 0
}

// defaultSourceSet returns true if the DM_DEFAULTSOURCE bit is set.
func (f devModeFields) defaultSourceSet() bool {
	return f&C.DM_DEFAULTSOURCE != 0
}

// printQualitySet returns true if the DM_PRINTQUALITY bit is set.
func (f devModeFields) printQualitySet() bool {
	return f&C.DM_PRINTQUALITY != 0
}

// colorSet returns true if the DM_COLOR bit is set.
func (f devModeFields) colorSet() bool {
	return f&C.DM_COLOR != 0
}

// duplexSet returns true if the DM_DUPLEX bit is set.
func (f devModeFields) duplexSet() bool {
	return f&C.DM_DUPLEX != 0
}

// yResolutionSet returns true if the DM_YRESOLUTION but is set.
func (f devModeFields) yResolutionSet() bool {
	return f&C.DM_YRESOLUTION != 0
}

// ttOptionSet returns true if the DM_TTOPTION bit is set.
func (f devModeFields) ttOptionSet() bool {
	return f&C.DM_TTOPTION != 0
}

// collateSet returns true if the DM_COLLATE bit is set.
func (f devModeFields) collateSet() bool {
	return f&C.DM_COLLATE != 0
}

// formNameSet returns true if the DM_FORMNAME bit is set.
func (f devModeFields) formNameSet() bool {
	return f&C.DM_FORMNAME != 0
}

// logPixelsSet returns true if the DM_LOGPIXELS bit is set.
func (f devModeFields) logPixelsSet() bool {
	return f&C.DM_LOGPIXELS != 0
}

// bitsPerPixelSet returns true if the DM_BITSPERPEL bit is set.
func (f devModeFields) bitsPerPixelSet() bool {
	return f&C.DM_BITSPERPEL != 0
}

// pelWidthSet returns true if the DM_PELSWIDTH bit is set.
func (f devModeFields) pelWidthSet() bool {
	return f&C.DM_PELSWIDTH != 0
}

// pelHeightSet returns true if the DM_PELSHEIGHT bit is set.
func (f devModeFields) pelHeightSet() bool {
	return f&C.DM_PELSHEIGHT != 0
}

// displayFlagsSet returns true if the DM_DISPLAYFLAGS bit is set.
func (f devModeFields) displayFlagsSet() bool {
	return f&C.DM_DISPLAYFLAGS != 0
}

// displayFrequencySet returns true if the DM_DISPLAYFREQUENCY bit is set.
func (f devModeFields) displayFrequencySet() bool {
	return f&C.DM_DISPLAYFREQUENCY != 0
}

// icmMethodSet returns true if the DM_ICMMETHOD bit is set.
func (f devModeFields) icmMethodSet() bool {
	return f&C.DM_ICMMETHOD != 0
}

// icmIntentSet returns true if the DM_ICMINTENT bit is set.
func (f devModeFields) icmIntentSet() bool {
	return f&C.DM_ICMINTENT != 0
}

// mediaTypeSet returns true if the DM_MEDIATYPE bit is set.
func (f devModeFields) mediaTypeSet() bool {
	return f&C.DM_MEDIATYPE != 0
}

// ditherTypeSet returns true if the DM_DITHERTYPE bit is set.
func (f devModeFields) ditherTypeSet() bool {
	return f&C.DM_DITHERTYPE != 0
}

// panningWidthSet returns true if the DM_PANNINGWIDTH bit is set.
func (f devModeFields) panningWidthSet() bool {
	return f&C.DM_PANNINGWIDTH != 0
}

// panningHeightSet returns true if the DM_PANNINGHEIGHT bit is set.
func (f devModeFields) panningHeightSet() bool {
	return f&C.DM_PANNINGHEIGHT != 0
}

// displayFixedOutputSet returns true if the DM_DISPLAYFIXEDOUTPUT bit is set.
func (f devModeFields) displayFixedOutputSet() bool {
	return f&C.DM_DISPLAYFIXEDOUTPUT != 0
}

// String returns a string containing the devMode fields that are set.
func (f devModeFields) String() string {
	var s strings.Builder
	s.WriteString("Fields:\n")
	if f.orientationSet() {
		s.WriteString("    DM_ORIENTATION\n")
	}
	if f.paperSizeSet() {
		s.WriteString("    DM_PAPERSIZE\n")
	}
	if f.paperLengthSet() {
		s.WriteString("    DM_PAPERLENGTH\n")
	}
	if f&C.DM_PAPERWIDTH != 0 {
		s.WriteString("    DM_PAPERWIDTH\n")
	}
	if f&C.DM_SCALE != 0 {
		s.WriteString("    DM_SCALE\n")
	}
	if f&C.DM_POSITION != 0 {
		s.WriteString("    DM_POSITION\n")
	}
	if f&C.DM_NUP != 0 {
		s.WriteString("    DM_NUP\n")
	}
	if f&C.DM_DISPLAYORIENTATION != 0 {
		s.WriteString("    DM_DISPLAYORIENTATION\n")
	}
	if f&C.DM_COPIES != 0 {
		s.WriteString("    DM_COPIES\n")
	}
	if f&C.DM_DEFAULTSOURCE != 0 {
		s.WriteString("    DM_DEFAULTSOURCE\n")
	}
	if f&C.DM_PRINTQUALITY != 0 {
		s.WriteString("    DM_PRINTQUALITY\n")
	}
	if f&C.DM_COLOR != 0 {
		s.WriteString("    DM_COLOR\n")
	}
	if f&C.DM_DUPLEX != 0 {
		s.WriteString("    DM_DUPLEX\n")
	}
	if f&C.DM_YRESOLUTION != 0 {
		s.WriteString("    DM_YRESOLUTION\n")
	}
	if f&C.DM_TTOPTION != 0 {
		s.WriteString("    DM_TTOPTION\n")
	}
	if f&C.DM_COLLATE != 0 {
		s.WriteString("    DM_COLLATE\n")
	}
	if f&C.DM_FORMNAME != 0 {
		s.WriteString("    DM_FORMNAME\n")
	}
	if f&C.DM_LOGPIXELS != 0 {
		s.WriteString("    DM_LOGPIXELS\n")
	}
	if f&C.DM_BITSPERPEL != 0 {
		s.WriteString("    DM_BITSPERPEL\n")
	}
	if f&C.DM_PELSWIDTH != 0 {
		s.WriteString("    DM_PELSWIDTH\n")
	}
	if f&C.DM_PELSHEIGHT != 0 {
		s.WriteString("    DM_PELSHEIGHT\n")
	}
	if f&C.DM_DISPLAYFLAGS != 0 {
		s.WriteString("    DM_DISPLAYFLAGS\n")
	}
	if f&C.DM_DISPLAYFREQUENCY != 0 {
		s.WriteString("    DM_DISPLAYFREQUENCY\n")
	}
	if f&C.DM_ICMMETHOD != 0 {
		s.WriteString("    DM_ICMMETHOD\n")
	}
	if f&C.DM_ICMINTENT != 0 {
		s.WriteString("    DM_ICMINTENT\n")
	}
	if f&C.DM_MEDIATYPE != 0 {
		s.WriteString("    DM_MEDIATYPE\n")
	}
	if f&C.DM_DITHERTYPE != 0 {
		s.WriteString("    DM_DITHERTYPE\n")
	}
	if f&C.DM_PANNINGWIDTH != 0 {
		s.WriteString("    DM_PANNINGWIDTH\n")
	}
	if f&C.DM_PANNINGHEIGHT != 0 {
		s.WriteString("    DM_PANNINGHEIGHT\n")
	}
	if f&C.DM_DISPLAYFIXEDOUTPUT != 0 {
		s.WriteString("    DM_DISPLAYFIXEDOUTPUT\n")
	}
	return s.String()
}
