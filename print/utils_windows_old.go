package print

/*
import (
	"syscall"
	"unsafe"
)

// maxUTF16Len is the maximum length of a UTF16 slice.
const maxUTF16Len = 1024

// StringFromUTF16 converts a pointer to a slice of UTF16 characters to a string.
// The maximum length of the UTF16 slice is maxUTF16Len as that should be more than enough for any
// UTF16 string returned from win32 printer functions.
func StringFromUTF16(utf16 *uint16) string {
	p := unsafe.Pointer(utf16)
	pSlice := (*[1 << 30]uint16)(p)[0:maxUTF16Len]

	return syscall.UTF16ToString(pSlice)

}
*/
