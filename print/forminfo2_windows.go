package print

//#define UNICODE
//#include "windows.h"
import "C"
import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
)

// formInfo2 contains information about a localizable printer (media) form.
type formInfo2 C.FORM_INFO_2W

// flags retrieves the formInfo2 object's Flags field
func (f *formInfo2) flags() formInfo2Flags {
	return formInfo2Flags(f.Flags)
}

// String returns a string representation of the formInfo2 object
func (f *formInfo2) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("Form Type: %s\n", f.flags().String()))
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
