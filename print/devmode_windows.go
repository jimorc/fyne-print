package print

//#define UNICODE
//#include "windows.h"
import "C"

import (
	"fmt"
	"strings"
	"unsafe"

	"fyne.io/fyne/v2"
	"golang.org/x/sys/windows"
)

// devMode is the Win32 DEVMODEW struct
type devMode C.DEVMODEW

//The following methods return the DEVMODEW fields with the "dm" prefix removed.

// DeviceName returns the printer's device name.
func (d *devMode) DeviceName() string {
	if d.dmDeviceName[0] == 0 {
		return "(none)"
	}
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
func (d *devMode) Orientation() orientation {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[unsafe.Sizeof(d.anon0) / 2]uint16)(p)[0:1]
	return orientation(pSlice[0])
}

// PaperSize returns the printer's media size as one of the DMPAPER_xxx
// values defined in wingdi.h. If PaperSize returns 0, then the media size
// is specified by the PaperLength and PaperWidth methods.
func (d *devMode) PaperSize() paperSize {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[unsafe.Sizeof(d.anon0) / 2]uint16)(p)[1:2]
	return paperSize(pSlice[0])
}

// PaperLength returns the printer's media size length in 1/10ths of a mm. This
// value is valid only if PaperSize return 0.
func (d *devMode) PaperLength() float32 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[unsafe.Sizeof(d.anon0)]uint16)(p)[2:3]
	return float32(pSlice[0]) / 10.
}

// PaperWidth returns the printer's media size width in 1/10ths of a mm. This
// value is valid only if PaperSize return 0.
func (d *devMode) PaperWidth() float32 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[unsafe.Sizeof(d.anon0)]uint16)(p)[3:4]
	return float32(pSlice[0]) / 10.
}

// Scale returns the percentage by which the image is to be scaled for printing.
// The image's media size is scaled to the physical page by the factor of Scale()/100.
func (d *devMode) Scale() uint16 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[unsafe.Sizeof(d.anon0)]uint16)(p)[4:5]
	return pSlice[0]
}

// Copies returns the number of copies to be printed if the driver supports
// multiple copies
func (d *devMode) Copies() uint16 {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[unsafe.Sizeof(d.anon0)]uint16)(p)[5:6]
	return pSlice[0]
}

// DefaultSource returns the printer's default input bin. This must be one of
// // the DMBIN-prefixed constants defined in wingdi.h. If the specified constant
// is DMBIN_FORMSOURCE, the input bin should be selected automatically.
func (d *devMode) DefaultSource() defaultSource {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[unsafe.Sizeof(d.anon0)]uint16)(p)[6:7]
	return defaultSource(pSlice[0])
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
func (d *devMode) PrintQuality() printQuality {
	p := unsafe.Pointer(&d.anon0[0])
	pSlice := (*[unsafe.Sizeof(d.anon0)]int16)(p)[7:8]
	return printQuality(pSlice[0])
}

// Color indicates whether a color printer should print in color or monochrome mode.
// The returned value is one of DMCOLOR_COLOR or DMCOLOR_MONOCHROME.
func (d *devMode) Color() color {
	return color(d.dmColor)
}

// Duplex specifies duplex (double-sided) printing for duplex-capable printers.
// The returned value is one of DMDUP_HORIZONTAL, DM_DUP_SIMPLEX, and DMDUP_VERTICAL.
func (d *devMode) Duplex() duplex {
	return duplex(d.dmDuplex)
}

// YResolution returns the y resolution of the printer in DPI. This value is valid only
// if PrintQuality() returns a positive value.
func (d *devMode) YResolution() uint16 {
	return uint16(d.dmYResolution)
}

// TTOPtion specifies how TrueType fonts should be printed. The returned value will be
// one of the DMTT_ prefixed values defined in wingdi.h.
func (d *devMode) TTOption() ttOption {
	return ttOption(d.dmTTOption)
}

// Collate indicates whether multiple prints should be collated. Valid values are
// DMCOLLATE_TRUE and DM_COLLATE_FALSE.
func (d *devMode) Collate() collate {
	return collate(d.dmCollate)
}

// FormName returns the name of the form (media size) to use. This is a name
// that is returned by the Printer.EnumForms method.
func (d *devMode) FormName() string {
	if d.dmFormName[0] == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(&d.dmFormName[0])))
}

// Nup specifies whether the printer handles N-up printing (printing multiple pages
// on a single physical page). The returned value is one of the following:
// DMNUP_SYSTEM: The print system handles N-up printing.
// DMNUP_ONEUP: The print system does not handle "N-up" printing. An application can
// set dmNup to DMNUP_ONEUP if it intends to carry out "N-up" printing on its own.
func (d *devMode) Nup() nup {
	p := unsafe.Pointer(&d.anon1[0])
	pSlice := (*[1 << 30]uint32)(p)[0:4]
	return nup(pSlice[0])
}

// ICMMethod specifies how ICM processing should be performed. Valid values and
// their meanings are a follows:
// DMICMMETHOD_NONE: ICM is disabled.
// DMICMMETHOD_SYSTEM: ICM is handled by the print system.
// DMICMMETHOD_DRIVER: ICM is handled by the printer driver.
// DMICMMETHOD_DEVICE: ICM is handled by the printer device.
func (d *devMode) ICMMethod() icmMethod {
	return icmMethod(d.dmICMMethod)
}

// ICMIntent specifies the ICM Intent. Valud values and meanings are as follows:
// DMICM_SATURATE: Maximize color saturation.
// DMICM_CONTRAST: Maximize color contrast.
// DMICM_COLORIMETRIC: Use specific color metric.
// DMICM_ABS_COLORIMETRIC: Use specific color metric.
func (d *devMode) ICMIntent() icmIntent {
	return icmIntent(d.dmICMIntent)
}

// MediaType returns the media type. Valid values and meanings are as follows:
// DMMEDIA_STANDARD: Standard paper.
// DMMEDIA_TRANSPARENCY: Transparency.
// DMMEDIA_GLOSSY: Glossy paper.
func (d *devMode) MediaType() mediaType {
	return mediaType(d.dmMediaType)
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
		s.WriteString(fmt.Sprintf("    Orientation: %s\n", d.Orientation().String()))
	}
	if f.paperSizeSet() {
		s.WriteString(fmt.Sprintf("    Paper Size: %s\n", d.PaperSize().String()))
	}
	if f.paperLengthSet() {
		s.WriteString(fmt.Sprintf("    Paper Length: %.1f mm\n", d.PaperLength()))
	}
	if f.paperWidthSet() {
		s.WriteString(fmt.Sprintf("    Paper Width: %.1f mm\n", d.PaperWidth()))
	}
	if f.scaleSet() {
		s.WriteString(fmt.Sprintf("    Scale: %d%%\n", d.Scale()))
	}
	if f.copiesSet() {
		s.WriteString(fmt.Sprintf("    Copies: %d\n", d.Copies()))
	}
	if f.defaultSourceSet() {
		s.WriteString(fmt.Sprintf("    Default Source: %s\n", d.DefaultSource().String()))
	}
	if f.printQualitySet() {
		s.WriteString(fmt.Sprintf("    Print Quality: %s\n", d.PrintQuality().String()))
	}
	if f.colorSet() {
		s.WriteString(fmt.Sprintf("    Color: %s\n", d.Color().String()))
	}
	if f.duplexSet() {
		s.WriteString(fmt.Sprintf("    Duplex: %s\n", d.Duplex().String()))
	}
	if f.yResolutionSet() {
		s.WriteString(fmt.Sprintf("    Y Resolution: %d dpi\n", d.YResolution()))
	}
	if f.ttOptionSet() {
		s.WriteString(fmt.Sprintf("    TT Option: %s\n", d.TTOption().String()))
	}
	if f.collateSet() {
		s.WriteString(fmt.Sprintf("    Collate: %s\n", d.Collate().String()))
	}
	if f.formNameSet() {
		s.WriteString(fmt.Sprintf("    Form Name: %s\n", d.FormName()))
	}
	if f.nupSet() {
		s.WriteString(fmt.Sprintf("    Nup: %s\n", d.Nup().String()))
	}
	if f.icmMethodSet() {
		s.WriteString(fmt.Sprintf("    ICM Method: %s\n", d.ICMMethod().String()))
	}
	if f.icmIntentSet() {
		s.WriteString(fmt.Sprintf("    ICM Intent: %s\n", d.ICMIntent().String()))
	}
	if f.mediaTypeSet() {
		s.WriteString(fmt.Sprintf("    Media Type: %s\n", d.MediaType().String()))
	}
	if f.ditherTypeSet() {
		s.WriteString(fmt.Sprintf("    Dither Type: %d\n", d.DitherType()))
	}
	return s.String()
}

// devModeFields is the dmFields value from the devMode object.
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

// orientation is the dmPrientation field from the devMode struct.
type orientation uint16

// String outputs the paper orientation.
func (o orientation) String() string {
	switch o {
	case C.DMORIENT_PORTRAIT:
		return "DMORIENT_PORTRAIT"
	case C.DMORIENT_LANDSCAPE:
		return "DMORIENT_LANDSCAPE"
	default:
		return fmt.Sprintf("Unknown orientation: %d", o)
	}
}

// paperSize is the dmPaperSize field from the devMode struct.
type paperSize uint16

// String outputs the paper size.
func (p paperSize) String() string {
	switch p {
	case C.DMPAPER_LETTER:
		return "DMPAPER_LETTER"
	case C.DMPAPER_LETTERSMALL:
		return "DMPAPER_LETTERSMALL"
	case C.DMPAPER_TABLOID:
		return "DMPAPER_TABLOID"
	case C.DMPAPER_LEDGER:
		return "DMPAPER_LEDGER"
	case C.DMPAPER_LEGAL:
		return "DMPAPER_LEGAL"
	case C.DMPAPER_STATEMENT:
		return "DMPAPER_STATEMENT"
	case C.DMPAPER_EXECUTIVE:
		return "DMPAPER_EXECUTIVE"
	case C.DMPAPER_A3:
		return "DMPAPER_A3"
	case C.DMPAPER_A4:
		return "DMPAPER_A4"
	case C.DMPAPER_A4SMALL:
		return "DMPAPER_A4SMALL"
	case C.DMPAPER_A5:
		return "DMPAPER_A5"
	case C.DMPAPER_B4:
		return "DMPAPER_B4"
	case C.DMPAPER_B5:
		return "DMPAPER_B5"
	case C.DMPAPER_FOLIO:
		return "DMPAPER_FOLIO"
	case C.DMPAPER_QUARTO:
		return "DMPAPER_QUARTO"
	case C.DMPAPER_10X14:
		return "DMPAPER_10X14"
	case C.DMPAPER_11X17:
		return "DMPAPER_11X17"
	case C.DMPAPER_NOTE:
		return "DMPAPER_NOTE"
	case C.DMPAPER_ENV_9:
		return "DMPAPER_ENV_9"
	case C.DMPAPER_ENV_10:
		return "DMPAPER_ENV_10"
	case C.DMPAPER_ENV_11:
		return "DMPAPER_ENV_11"
	case C.DMPAPER_ENV_12:
		return "DMPAPER_ENV_12"
	case C.DMPAPER_ENV_14:
		return "DMPAPER_ENV_14"
	case C.DMPAPER_CSHEET:
		return "DMPAPER_CSHEET"
	case C.DMPAPER_DSHEET:
		return "DMPAPER_DSHEET"
	case C.DMPAPER_ESHEET:
		return "DMPAPER_ESHEET"
	case C.DMPAPER_ENV_DL:
		return "DMPAPER_ENV_DL"
	case C.DMPAPER_ENV_C5:
		return "DMPAPER_ENV_C5"
	case C.DMPAPER_ENV_C3:
		return "DMPAPER_ENV_C3"
	case C.DMPAPER_ENV_C4:
		return "DMPAPER_ENV_C4"
	case C.DMPAPER_ENV_C6:
		return "DMPAPER_ENV_C6"
	case C.DMPAPER_ENV_C65:
		return "DMPAPER_ENV_C65"
	case C.DMPAPER_ENV_B4:
		return "DMPAPER_ENV_B4"
	case C.DMPAPER_ENV_B5:
		return "DMPAPER_ENV_B5"
	case C.DMPAPER_ENV_ITALY:
		return "DMPAPER_ENV_ITALY"
	case C.DMPAPER_ENV_MONARCH:
		return "DMPAPER_ENV_MONARCH"
	case C.DMPAPER_ENV_PERSONAL:
		return "DMPAPER_ENV_PERSONAL"
	case C.DMPAPER_FANFOLD_US:
		return "DMPAPER_FANFOLD_US"
	case C.DMPAPER_FANFOLD_STD_GERMAN:
		return "DMPAPER_FANFOLD_STD_GERMAN"
	case C.DMPAPER_FANFOLD_LGL_GERMAN:
		return "DMPAPER_FANFOLD_LGL_GERMAN"
	case C.DMPAPER_ISO_B4:
		return "DMPAPER_ISO_B4"
	case C.DMPAPER_JAPANESE_POSTCARD:
		return "DMPAPER_JAPANESE_POSTCARD"
	case C.DMPAPER_9X11:
		return "DMPAPER_9X11"
	case C.DMPAPER_10X11:
		return "DMPAPER_10X11"
	case C.DMPAPER_15X11:
		return "DMPAPER_15X11"
	case C.DMPAPER_ENV_INVITE:
		return "DMPAPER_ENV_INVITE"
	case C.DMPAPER_RESERVED_48:
		return "DMPAPER_RESERVED_48"
	case C.DMPAPER_RESERVED_49:
		return "DMPAPER_RESERVED_49"
	case C.DMPAPER_LETTER_EXTRA:
		return "DMPAPER_LETTER_EXTRA"
	case C.DMPAPER_LEGAL_EXTRA:
		return "DMPAPER_LEGAL_EXTRA"
	case C.DMPAPER_TABLOID_EXTRA:
		return "DMPAPER_TABLOID_EXTRA"
	case C.DMPAPER_A4_EXTRA:
		return "DMPAPER_A4_EXTRA"
	case C.DMPAPER_LETTER_TRANSVERSE:
		return "DMPAPER_LETTER_TRANSVERSE"
	case C.DMPAPER_A4_TRANSVERSE:
		return "DMPAPER_A4_TRANSVERSE"
	case C.DMPAPER_LETTER_EXTRA_TRANSVERSE:
		return "DMPAPER_LETTER_EXTRA_TRANSVERSE"
	case C.DMPAPER_A_PLUS:
		return "DMPAPER_A_PLUS"
	case C.DMPAPER_B_PLUS:
		return "DMPAPER_B_PLUS"
	case C.DMPAPER_LETTER_PLUS:
		return "DMPAPER_LETTER_PLUS"
	case C.DMPAPER_A4_PLUS:
		return "DMPAPER_A4_PLUS"
	case C.DMPAPER_A5_TRANSVERSE:
		return "DMPAPER_A5_TRANSVERSE"
	case C.DMPAPER_B5_TRANSVERSE:
		return "DMPAPER_B5_TRANSVERSE"
	case C.DMPAPER_A3_EXTRA:
		return "DMPAPER_A3_EXTRA"
	case C.DMPAPER_DBL_JAPANESE_POSTCARD:
		return "DMPAPER_DBL_JAPANESE_POSTCARD"
	case C.DMPAPER_A6:
		return "DMPAPER_A6"
	case C.DMPAPER_JENV_KAKU2:
		return "DMPAPER_JENV_KAKU2"
	case C.DMPAPER_JENV_KAKU3:
		return "DMPAPER_JENV_KAKU3"
	case C.DMPAPER_JENV_CHOU3:
		return "DMPAPER_JENV_CHOU3"
	case C.DMPAPER_JENV_CHOU4:
		return "DMPAPER_JENV_CHOU4"
	case C.DMPAPER_LETTER_ROTATED:
		return "DMPAPER_LETTER_ROTATED"
	case C.DMPAPER_A3_ROTATED:
		return "DMPAPER_A3_ROTATED"
	case C.DMPAPER_A4_ROTATED:
		return "DMPAPER_A4_ROTATED"
	case C.DMPAPER_A5_ROTATED:
		return "DMPAPER_A5_ROTATED"
	case C.DMPAPER_B4_JIS_ROTATED:
		return "DMPAPER_B4_JIS_ROTATED"
	case C.DMPAPER_B5_JIS_ROTATED:
		return "DMPAPER_B5_JIS_ROTATED"
	case C.DMPAPER_JAPANESE_POSTCARD_ROTATED:
		return "DMPAPER_JAPANESE_POSTCARD_ROTATED"
	case C.DMPAPER_DBL_JAPANESE_POSTCARD_ROTATED:
		return "DMPAPER_DBL_JAPANESE_POSTCARD_ROTATED"
	case C.DMPAPER_A6_ROTATED:
		return "DMPAPER_A6_ROTATED"
	case C.DMPAPER_JENV_KAKU2_ROTATED:
		return "DMPAPER_JENV_KAKU2_ROTATED"
	case C.DMPAPER_JENV_KAKU3_ROTATED:
		return "DMPAPER_JENV_KAKU3_ROTATED"
	case C.DMPAPER_JENV_CHOU3_ROTATED:
		return "DMPAPER_JENV_CHOU3_ROTATED"
	case C.DMPAPER_JENV_CHOU4_ROTATED:
		return "DMPAPER_JENV_CHOU4_ROTATED"
	case C.DMPAPER_B6_JIS:
		return "DMPAPER_B6_JIS"
	case C.DMPAPER_B6_JIS_ROTATED:
		return "DMPAPER_B6_JIS_ROTATED"
	case C.DMPAPER_12X11:
		return "DMPAPER_12X11"
	case C.DMPAPER_JENV_YOU4:
		return "DMPAPER_JENV_YOU4"
	case C.DMPAPER_JENV_YOU4_ROTATED:
		return "DMPAPER_JENV_YOU4_ROTATED"
	case C.DMPAPER_P16K:
		return "DMPAPER_P16K"
	case C.DMPAPER_P32K:
		return "DMPAPER_P32K"
	case C.DMPAPER_P32KBIG:
		return "DMPAPER_P32KBIG"
	case C.DMPAPER_PENV_1:
		return "DMPAPER_PENV_1"
	case C.DMPAPER_PENV_2:
		return "DMPAPER_PENV_2"
	case C.DMPAPER_PENV_3:
		return "DMPAPER_PENV_3"
	case C.DMPAPER_PENV_4:
		return "DMPAPER_PENV_4"
	case C.DMPAPER_PENV_5:
		return "DMPAPER_PENV_5"
	case C.DMPAPER_PENV_6:
		return "DMPAPER_PENV_6"
	case C.DMPAPER_PENV_7:
		return "DMPAPER_PENV_7"
	case C.DMPAPER_PENV_8:
		return "DMPAPER_PENV_8"
	case C.DMPAPER_PENV_9:
		return "DMPAPER_PENV_9"
	case C.DMPAPER_PENV_10:
		return "DMPAPER_PENV_10"
	case C.DMPAPER_P16K_ROTATED:
		return "DMPAPER_P16K_ROTATED"
	case C.DMPAPER_P32K_ROTATED:
		return "DMPAPER_P32K_ROTATED"
	case C.DMPAPER_P32KBIG_ROTATED:
		return "DMPAPER_P32KBIG_ROTATED"
	case C.DMPAPER_PENV_1_ROTATED:
		return "DMPAPER_PENV_1_ROTATED"
	case C.DMPAPER_PENV_2_ROTATED:
		return "DMPAPER_PENV_2_ROTATED"
	case C.DMPAPER_PENV_3_ROTATED:
		return "DMPAPER_PENV_3_ROTATED"
	case C.DMPAPER_PENV_4_ROTATED:
		return "DMPAPER_PENV_4_ROTATED"
	case C.DMPAPER_PENV_5_ROTATED:
		return "DMPAPER_PENV_5_ROTATED"
	case C.DMPAPER_PENV_6_ROTATED:
		return "DMPAPER_PENV_6_ROTATED"
	case C.DMPAPER_PENV_7_ROTATED:
		return "DMPAPER_PENV_7_ROTATED"
	case C.DMPAPER_PENV_8_ROTATED:
		return "DMPAPER_PENV_8_ROTATED"
	case C.DMPAPER_PENV_9_ROTATED:
		return "DMPAPER_PENV_9_ROTATED"
	case C.DMPAPER_PENV_10_ROTATED:
		return "DMPAPER_PENV_10_ROTATED"
	default:
		if p >= C.DMPAPER_USER {
			return "DMPAPER_USER Defined"
		}
		return fmt.Sprintf("Unknown paper size: %d", p)
	}
}

// defaultSource represents the default media source for the printer.
type defaultSource uint16

// String returns the default source as a string.
func (d defaultSource) String() string {
	switch d {
	case C.DMBIN_UPPER:
		return "DMBIN_UPPER"
	case C.DMBIN_LOWER:
		return "Lower"
	case C.DMBIN_MIDDLE:
		return "DMBIN_MIDDLE"
	case C.DMBIN_MANUAL:
		return "DMBIN_MANUAL"
	case C.DMBIN_ENVELOPE:
		return "DMBIN_ENVELOPE"
	case C.DMBIN_ENVMANUAL:
		return "DMBIN_ENVMANUAL"
	case C.DMBIN_AUTO:
		return "DMBIN_AUTO"
	case C.DMBIN_TRACTOR:
		return "DMBIN_TRACTOR"
	case C.DMBIN_SMALLFMT:
		return "DMBIN_SMALLFMT"
	case C.DMBIN_LARGEFMT:
		return "DMBIN_LARGEFMT"
	case C.DMBIN_LARGECAPACITY:
		return "DMBIN_LARGECAPACITY"
	case C.DMBIN_CASSETTE:
		return "DMBIN_CASSETTE"
	case C.DMBIN_FORMSOURCE:
		return "DMBIN_FORMSOURCE"
	default:
		return fmt.Sprintf("Unknown: %d", d)
	}
}

// printQuality defines the print quality, either as set string or as dpi.
type printQuality int16

// String returns the print quality value as a string.
func (p printQuality) String() string {
	switch p {
	case C.DMRES_DRAFT:
		return "Draft"
	case C.DMRES_LOW:
		return "Low"
	case C.DMRES_MEDIUM:
		return "Medium"
	case C.DMRES_HIGH:
		return "High"
	default:
		if p > 0 {
			return fmt.Sprintf("%d dpi", p)
		} else {
			return fmt.Sprintf("Unknown value: %d", p)
		}
	}
}

// color defines the color setting (color or monochrome) for the printer.
type color uint16

// String returns the color value as a string.
func (c color) String() string {
	switch c {
	case C.DMCOLOR_MONOCHROME:
		return "Monochrome"
	case C.DMCOLOR_COLOR:
		return "Color"
	default:
		return fmt.Sprintf("Unknown value: %d", c)
	}
}

// duplex defines the duplex setting for the printer.
type duplex uint16

// String returns the duplex value as a string.
func (d duplex) String() string {
	switch d {
	case C.DMDUP_SIMPLEX:
		return "Simplex"
	case C.DMDUP_VERTICAL:
		return "Long Edge Binding"
	case C.DMDUP_HORIZONTAL:
		return "Short Edge Binding"
	default:
		err := fmt.Errorf("unknown duplex value: %d", d)
		fyne.LogError("Invalid DevMode setting: ", err)
		return "Invalid value"
	}
}

// ttOPtion defines how TrueType fonts should be printed.
type ttOption int16

// String returns the ttOption value as a string.
func (t ttOption) String() string {
	switch t {
	case C.DMTT_BITMAP:
		return "Bitmap"
	case C.DMTT_DOWNLOAD:
		return "Download"
	case C.DMTT_SUBDEV:
		return "Substitute printer fonts"
	case C.DMTT_DOWNLOAD_OUTLINE:
		return "Download as outline soft fonts"
	default:
		err := fmt.Errorf("unknown ttOption value: %d", t)
		fyne.LogError("Invalid DevMode setting: ", err)
		return "Invalid value"
	}
}

// collatae defines whether collation should be used when printing multiple copies.
type collate int16

// String returns the collate value as a string.
func (c collate) String() string {
	switch c {
	case C.DMCOLLATE_TRUE:
		return "Collate"
	case C.DMCOLLATE_FALSE:
		return "Don't Collate"
	default:
		err := fmt.Errorf("unknown collate value: %d", c)
		fyne.LogError("Invalid DevMode setting: ", err)
		return "Invalid value"
	}
}

// nup specifies where N-up processing is done.
type nup uint32

// String returns the NUP value as a string.
func (n nup) String() string {
	switch n {
	case C.DMNUP_SYSTEM:
		return "Print spooler does"
	case C.DMNUP_ONEUP:
		return "Application does"
	default:
		err := fmt.Errorf("unknown nup value: %d", n)
		fyne.LogError("Invalid DevMode setting: ", err)
		return "Invalid value"
	}
}

// icmMethod specifies how ICM is handled.  For a non-ICM application,
// this member determines if ICM is enabled or disabled. For ICM
// applications, the system examines this member to determine how to
// handle ICM support.
type icmMethod uint32

// String returns the ICMMethod value as a string.
func (i icmMethod) String() string {
	switch i {
	case C.DMICMMETHOD_NONE:
		return "None"
	case C.DMICMMETHOD_SYSTEM:
		return "System"
	case C.DMICMMETHOD_DRIVER:
		return "Driver"
	case C.DMICMMETHOD_DEVICE:
		return "Device"
	default:
		err := fmt.Errorf("unknown icmMethod value: %d", i)
		fyne.LogError("Invalid DevMode setting: ", err)
		return "Invalid value"
	}
}

// icmIntent spcifies which color matching method, or intent, should
// be used by default.  This member is primarily for non-ICM applications.
// ICM applications can establish intents by using the ICM functions.
type icmIntent uint32

// String returns the ICMIntent value as a string.
func (i icmIntent) String() string {
	switch i {
	case C.DMICM_SATURATE:
		return "Saturate"
	case C.DMICM_CONTRAST:
		return "Contrast"
	case C.DMICM_COLORIMETRIC:
		return "Colorimetric"
	case C.DMICM_ABS_COLORIMETRIC:
		return "Absolute Colorimetric"
	default:
		err := fmt.Errorf("unknown icmIntent value: %d", i)
		fyne.LogError("Invalid DevMode setting: ", err)
		return "Invalid value"
	}
}

// mediaType specifies the type of media being printed on.
type mediaType uint32

// String returns the mediaType value as a string.
func (m mediaType) String() string {
	switch m {
	case C.DMMEDIA_STANDARD:
		return "Standard"
	case C.DMMEDIA_TRANSPARENCY:
		return "Transparency"
	case C.DMMEDIA_GLOSSY:
		return "Glossy"
	default:
		if m >= C.DMMEDIA_USER {
			return "User Defined"
		}
		err := fmt.Errorf("unknown mediaType value: %d", m)
		fyne.LogError("Invalid DevMode setting: ", err)
		return "Invalid value"
	}
}
