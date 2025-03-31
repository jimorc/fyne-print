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

// PrinterInfo2 specifies detailed printer information.
type PrinterInfo2 C.PRINTER_INFO_2W

// ServerName returns the printer server name.
func (pi2 *PrinterInfo2) ServerName() string {
	if pi2.pServerName == nil || *pi2.pServerName == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pServerName)))
}

// PrinterName returns the printer name.
func (pi2 *PrinterInfo2) PrinterName() string {
	if pi2.pPrinterName == nil || *pi2.pPrinterName == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pPrinterName)))
}

// ShareName returns the printer's share name.
func (pi2 *PrinterInfo2) ShareName() string {
	if pi2.pShareName == nil || *pi2.pShareName == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pShareName)))
}

// PortName returns the printer's port name.
func (pi2 *PrinterInfo2) PortName() string {
	if pi2.pPortName == nil || *pi2.pPortName == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pPortName)))
}

// DriverName returns the printer's driver name.
func (pi2 *PrinterInfo2) DriverName() string {
	if pi2.pDriverName == nil || *pi2.pDriverName == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pDriverName)))
}

// Comment returns the comment associated with the printer.
func (pi2 *PrinterInfo2) Comment() string {
	if pi2.pComment == nil || *pi2.pComment == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pComment)))
}

// Location returns the printer's location as set during its configuration.
func (pi2 *PrinterInfo2) Location() string {
	if pi2.pLocation == nil || *pi2.pLocation == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pLocation)))
}

// DevMode returns the printer's devMode object.
func (pi2 *PrinterInfo2) DevMode() *devMode {
	return (*devMode)(unsafe.Pointer(pi2.pDevMode))
}

// SepFile returns the name of the file used to create the separator page.
// This page separates print jobs sent to the printer.
func (pi2 *PrinterInfo2) SepFile() string {
	if pi2.pSepFile == nil || *pi2.pSepFile == 0 {
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pSepFile)))
}

// PrintProcessor returns the name of the print processor used by the printer.
func (pi2 *PrinterInfo2) PrintProcessor() string {
	if pi2.pPrintProcessor == nil || *pi2.pPrintProcessor == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pPrintProcessor)))
}

// Datatype returns the data type that is used to record the print job. You can use
// the EnumPrintProcessorDatatypes function to obtain a list of data types supported
// by a specific print processor.
func (pi2 PrinterInfo2) Datatype() string {
	if pi2.pDatatype == nil || *pi2.pDatatype == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pDatatype)))
}

// Parameters returns the default print-processor parameters.
func (pi2 *PrinterInfo2) Parameters() string {
	if pi2.pParameters == nil || *pi2.pParameters == 0 {
		return "(none)"
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(pi2.pParameters)))
}

// SecurityDescriptor returns the security descriptor associated with the printer.
// This may be nil.
func (pi2 *PrinterInfo2) SecurityDescriptor() *securityDescriptor {
	if pi2.pSecurityDescriptor == nil {
		return nil
	}
	return (*securityDescriptor)(unsafe.Pointer(pi2.pSecurityDescriptor))
}

// Attrs returns the printer attributes.
// See http://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-2
// for values.
func (pi2 *PrinterInfo2) Attrs() attributes {
	return attributes((pi2.Attributes))
}

// JobPriority returns the priority value that the spooler uses to route print jobs.
func (pi2 *PrinterInfo2) JobPriority() uint32 {
	return uint32(pi2.Priority)
}

// DefPriority returns the default priority assigned to each print job.
func (pi2 *PrinterInfo2) DefPriority() uint32 {
	return uint32(pi2.DefaultPriority)
}

// PrinterStartTime returns the earliest time at which a printer will print
// a job. The format is (hour):(minute).
func (pi2 *PrinterInfo2) PrinterStartTime() string {
	start := uint32(pi2.StartTime)
	s := fmt.Sprintf("%d:%02d", start/60, start%60)
	return s
}

// PrinterUntilTime returns the latest time at which a printer will print
// a job. The format is (hour):(minute).
func (pi2 *PrinterInfo2) PrinterUntilTime() string {
	until := uint32(pi2.UntilTime)
	s := fmt.Sprintf("%d:%02d", until/60, until%60)
	return s
}

// PrinterStatus returns the printer status as a string slice.
// The following statuses are not defined by cgo, so they cannot
func (pi2 *PrinterInfo2) PrinterStatus() printerStatus {
	return printerStatus(pi2.Status)
}

// QueuedJobs returns the number of jobs queued for the printer.
func (pi2 *PrinterInfo2) QueuedJobs() uint32 {
	return uint32(pi2.cJobs)
}

// AveragePpm returns the average number of pages per minute that the printer can print.
func (pi2 *PrinterInfo2) AveragePpm() uint32 {
	return uint32(pi2.AveragePPM)
}

// String returns a string representation of the PrinterInfo2 struct.
func (pi2 *PrinterInfo2) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintln("PrinterInfo2:"))
	s.WriteString(fmt.Sprintf("    Server Name: %s\n", pi2.ServerName()))
	s.WriteString(fmt.Sprintf("    Printer Name: %s\n", pi2.PrinterName()))
	s.WriteString(fmt.Sprintf("    Share Name: %s\n", pi2.ShareName()))
	s.WriteString(fmt.Sprintf("    Port Name: %s\n", pi2.PortName()))
	s.WriteString(fmt.Sprintf("    Driver Name: %s\n", pi2.DriverName()))
	s.WriteString(fmt.Sprintf("    Comment: %s\n", pi2.Comment()))
	s.WriteString(fmt.Sprintf("    Location: %s\n", pi2.Location()))
	s.WriteString(prepend("    ", pi2.DevMode().String()))
	s.WriteString(fmt.Sprintf("    Sep File: %s\n", pi2.SepFile()))
	s.WriteString(fmt.Sprintf("    Print Processor: %s\n", pi2.PrintProcessor()))
	s.WriteString(fmt.Sprintf("    Data Type: %s\n", pi2.Datatype()))
	s.WriteString(fmt.Sprintf("    Parameters: %s\n", pi2.Parameters()))
	s.WriteString("    Attributes:\n")
	s.WriteString(prepend("    ", pi2.Attrs().String()))
	s.WriteString(fmt.Sprintf("    Priority: %d\n", pi2.JobPriority()))
	s.WriteString(fmt.Sprintf("    Default Priority: %d\n", pi2.DefPriority()))
	s.WriteString(fmt.Sprintf("    Start Time: %s\n", pi2.PrinterStartTime()))
	s.WriteString(fmt.Sprintf("    Until Time: %s\n", pi2.PrinterUntilTime()))
	s.WriteString("    Status:\n")
	s.WriteString(prepend("    ", pi2.PrinterStatus().String()))
	s.WriteString(fmt.Sprintf("    Jobs: %d\n", pi2.QueuedJobs()))
	s.WriteString(fmt.Sprintf("    Average PPMs: %d\n", pi2.AveragePpm()))
	return s.String()
}

// attributes specify the printer attributes.
type attributes uint32

// String outputs the attributes as a string.
func (a attributes) String() string {
	var s strings.Builder
	if a == 0 {
		s.WriteString("    None\n")
		return s.String()
	}
	if a&C.PRINTER_ATTRIBUTE_QUEUED != 0 {
		s.WriteString("    Queued\n")
	}
	if a&C.PRINTER_ATTRIBUTE_DIRECT != 0 {
		s.WriteString("    Direct\n")
	}
	if a&C.PRINTER_ATTRIBUTE_DEFAULT != 0 {
		s.WriteString("    Default\n")
	}
	if a&C.PRINTER_ATTRIBUTE_SHARED != 0 {
		s.WriteString("    Shared\n")
	}
	if a&C.PRINTER_ATTRIBUTE_NETWORK != 0 {
		s.WriteString("    Network\n")
	}
	if a&C.PRINTER_ATTRIBUTE_HIDDEN != 0 {
		s.WriteString("    Hidden\n")
	}
	if a&C.PRINTER_ATTRIBUTE_LOCAL != 0 {
		s.WriteString("    Local\n")
	}
	if a&C.PRINTER_ATTRIBUTE_ENABLE_DEVQ != 0 {
		s.WriteString("    Enable DevQ\n")
	}
	if a&C.PRINTER_ATTRIBUTE_KEEPPRINTEDJOBS != 0 {
		s.WriteString("    Keep Printed Jobs\n")
	}
	if a&C.PRINTER_ATTRIBUTE_DO_COMPLETE_FIRST != 0 {
		s.WriteString("    Do Complete First\n")
	}
	if a&C.PRINTER_ATTRIBUTE_WORK_OFFLINE != 0 {
		s.WriteString("    Work Offline\n")
	}
	if a&C.PRINTER_ATTRIBUTE_ENABLE_BIDI != 0 {
		s.WriteString("    Enable BiDirectional\n")
	}
	if a&C.PRINTER_ATTRIBUTE_RAW_ONLY != 0 {
		s.WriteString("    Raw Only\n")
	}
	if a&C.PRINTER_ATTRIBUTE_PUBLISHED != 0 {
		s.WriteString("    Published\n")
	}
	if a&C.PRINTER_ATTRIBUTE_FAX != 0 {
		s.WriteString("    Fax\n")
	}
	if a&C.PRINTER_ATTRIBUTE_TS != 0 {
		s.WriteString("    TS\n")
	}
	return s.String()
}

type printerStatus uint32

// String outputs the printer status as a string.
// PrinterStatus returns the printer status as a string slice.
// The following statuses are not defined by cgo, so they cannot
func (ps printerStatus) String() string {
	var status strings.Builder
	if ps == 0 {
		status.WriteString("    None")
		return status.String()
	}
	if ps&C.PRINTER_STATUS_PAUSED != 0 {
		status.WriteString("    Paused")
	}
	if ps&C.PRINTER_STATUS_ERROR != 0 {
		status.WriteString("    Error")
	}
	if ps&C.PRINTER_STATUS_PENDING_DELETION != 0 {
		status.WriteString("    Pending Deletion")
	}
	if ps&C.PRINTER_STATUS_PAPER_JAM != 0 {
		status.WriteString("    Paper Jam")
	}
	if ps&C.PRINTER_STATUS_PAPER_OUT != 0 {
		status.WriteString("    Paper Out")
	}
	if ps&C.PRINTER_STATUS_MANUAL_FEED != 0 {
		status.WriteString("    Manual Feed")
	}
	if ps&C.PRINTER_STATUS_PAPER_PROBLEM != 0 {
		status.WriteString("    Paper Problem")
	}
	if ps&C.PRINTER_STATUS_OFFLINE != 0 {
		status.WriteString("    Offline")
	}
	if ps&C.PRINTER_STATUS_IO_ACTIVE != 0 {
		status.WriteString("    Active")
	}
	if ps&C.PRINTER_STATUS_BUSY != 0 {
		status.WriteString("    Busy")
	}
	if ps&C.PRINTER_STATUS_PRINTING != 0 {
		status.WriteString("    Printing")
	}
	if ps&C.PRINTER_STATUS_OUTPUT_BIN_FULL != 0 {
		status.WriteString("    Output Bin Full")
	}
	if ps&C.PRINTER_STATUS_NOT_AVAILABLE != 0 {
		status.WriteString("    Not Available")
	}
	if ps&C.PRINTER_STATUS_WAITING != 0 {
		status.WriteString("    Waiting")
	}
	if ps&C.PRINTER_STATUS_PROCESSING != 0 {
		status.WriteString("    Processing")
	}
	if ps&C.PRINTER_STATUS_INITIALIZING != 0 {
		status.WriteString("    Initializing")
	}
	if ps&C.PRINTER_STATUS_WARMING_UP != 0 {
		status.WriteString("    Warming Up")
	}
	if ps&C.PRINTER_STATUS_TONER_LOW != 0 {
		status.WriteString("    Toner Low")
	}
	if ps&C.PRINTER_STATUS_NO_TONER != 0 {
		status.WriteString("    No Toner")
	}
	if ps&C.PRINTER_STATUS_PAGE_PUNT != 0 {
		status.WriteString("    Page Cannot Be Printed")
	}
	if ps&C.PRINTER_STATUS_USER_INTERVENTION != 0 {
		status.WriteString("    User Intervention Required")
	}
	if ps&C.PRINTER_STATUS_OUT_OF_MEMORY != 0 {
		status.WriteString("    Out of Memory")
	}
	if ps&C.PRINTER_STATUS_DOOR_OPEN != 0 {
		status.WriteString("    Door Open")
	}
	if ps&C.PRINTER_STATUS_SERVER_UNKNOWN != 0 {
		status.WriteString("    Server Unknown")
	}
	if ps&C.PRINTER_STATUS_POWER_SAVE != 0 {
		status.WriteString("    Power Save")
	}
	if ps&C.PRINTER_STATUS_SERVER_UNKNOWN != 0 {
		status.WriteString("    Server Unknown")
	}
	if ps&0x04000000 != 0 { // PRINTER_STATUS_DRIVER_UPDATE_NEEDED not defined by cgo
		status.WriteString("    Driver Update Needed")
	}
	return status.String()
}
