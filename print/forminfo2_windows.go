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

// formInfo2 contains information about a localizable printer (media) form.
type formInfo2 C.FORM_INFO_2W

// flags retrieves the formInfo2 object's Flags field
func (f *formInfo2) flags() formInfo2Flags {
	return formInfo2Flags(f.Flags)
}

// formName returns the formInfo2 object's form name.
func (f *formInfo2) formName() string {
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(f.pName)))
}

// size returns the form size in thousandths of a mm.
func (f *formInfo2) size() formSize {
	return formSize(f.Size)
}

// imageableArea returns the imageable are of a formInfo2 object.
func (f *formInfo2) imageableArea() imageableArea {
	return imageableArea(f.ImageableArea)
}

// keyWord returns the formInfo2 object's keyWord.
// This keyWord is used to identify the form in all locales.
// Unfortunately, not all forms have keyWord values.
func (f *formInfo2) keyWord() string {
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(f.pKeyword)))
}

// stringType returns the formInfo2's string type.
func (f *formInfo2) stringType() stringType {
	return stringType(f.StringType)
}

// muiDLL returns the formInfo2's localized resource DLL that contains the
// localized display name.
func (f *formInfo2) muiDLL() string {
	if f.stringType().isMuidll() {
		return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(f.pMuiDll)))
	}
	return ""
}

// muiResourceId returns the resource ID of the formInfo2's display name.
func (f *formInfo2) muiResourceId() uint32 {
	return (uint32(f.dwResourceId))
}

// String returns a string representation of the formInfo2 object.
func (f *formInfo2) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("Form Type: %s\n", f.flags().String()))
	s.WriteString(fmt.Sprintf("Form Name: %s\n", f.formName()))
	s.WriteString(fmt.Sprintf("Form Size: %s\n", f.size().String()))
	s.WriteString(fmt.Sprintf("Imageable Area: %s\n", f.imageableArea().String()))
	s.WriteString(fmt.Sprintf("Keyword: %s\n", f.keyWord()))
	s.WriteString(fmt.Sprintf("String Type: %s\n", f.stringType().String()))
	if f.stringType().isMuidll() {
		s.WriteString(fmt.Sprintf("MUI DLL: %s\n", f.muiDLL()))
		s.WriteString(fmt.Sprintf("MUI Resource ID: %d\n", f.muiResourceId()))
	}
	return s.String()
}

// formInfo2Flags represents the Form field of the formInfo2 object.
type formInfo2Flags uint32

// String returns a string repesentation of the formInfo2Flags object.
func (f formInfo2Flags) String() string {
	switch f {
	case C.FORM_USER:
		return "User Form"
	case C.FORM_BUILTIN:
		return "Builtin Form"
	case C.FORM_PRINTER:
		return "Printer Form"
	default:
		err := fmt.Errorf("unknown form type: %d", f)
		fyne.LogError("Invalid FormInfo setting: %s", err)
		return "Unknown form type"
	}
}

// formSize represents the Size field of a formInfo2 object.
type formSize C.SIZEL

// String returns a string representation of a formSize object.
func (f formSize) String() string {
	m := f.inMM()
	i := f.inInches()
	return fmt.Sprintf("%.3f x %.3f mm (%.3f x %.3f in)", m.Width, m.Height, i.Width, i.Height)
}

// inMM returns the size of a formSize object in mm.
func (f formSize) inMM() fyne.Size {
	return fyne.NewSize(float32(f.cx)/1000, float32(f.cy)/1000)
}

// imInches returns the size of a formSize object in inches.
func (f formSize) inInches() fyne.Size {
	return fyne.NewSize(float32(f.cx)/25400, float32(f.cy)/25400)
}

// imageableArea is the imageable area of a formInfo2 object.
type imageableArea C.RECTL

// String returns a string representation of an imageableArea object.
func (i imageableArea) String() string {
	return fmt.Sprintf("(%.3f, %.3f) mm to (%.3f, %.3f) mm",
		(float32(i.left))/1000, (float32(i.top))/1000,
		(float32(i.right))/1000, (float32(i.bottom))/1000)
}

// stringType is the string type of a formInfo2 object.
type stringType uint32

// the following values represent the string types. They are defined here
// because they are not defined in the mingw64 winspool.h file.
const (
	STRING_NONE     stringType = 0x00000001
	STRING_MUIDLL   stringType = 0x00000002
	STRING_LANGPAIR stringType = 0x00000004
)

// String returns a string representation of a stringType object.
func (st stringType) String() string {
	var s strings.Builder
	if st == 0 {
		return "Unknown value"
	}
	if st&STRING_NONE != 0 {
		s.WriteString("None")
	}
	if st&STRING_MUIDLL != 0 {
		if s.Len() > 0 {
			s.WriteString(" | ")
		}
		s.WriteString("MuiDLL")
	}
	if st&STRING_LANGPAIR != 0 {
		if s.Len() > 0 {
			s.WriteString(" | ")
		}
		s.WriteString("LangPair")
	}
	return s.String()
}

// isMuidll returns true if the stringType contains STRING_MUIDLL.
func (st stringType) isMuidll() bool {
	return st&STRING_MUIDLL != 0
}

// isLangPain returns true if the stringType contains STRING_LANGPAIR.
func (st stringType) isLangPair() bool {
	return st&STRING_LANGPAIR != 0
}
