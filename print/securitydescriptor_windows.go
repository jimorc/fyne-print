package print

//#include "windows.h"
import "C"

// securityDescriptor contains security information related to a printer.
// See https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-security_descriptor for
// information of the various fields.
//
// A securityDescriptor object is required, but values are currently not used.
type securityDescriptor C.SECURITY_DESCRIPTOR
