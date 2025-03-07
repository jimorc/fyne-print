//go:build windows

package print

import (
	"fmt"
	"strings"
)

const (
	CCHDEVICENAME = 32
	CCHFORMNAME   = 32
)

type PrinterAttribute uint32

// Printer attributes.
// These attributes may be set in the attributes field of the PrinterInfo2 struct.
// For attributes below that are not preceeded by comments, see
// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-2 for their meanings.
const (
	printJobQueued      PrinterAttribute = 0x00000001 // PRINTER_ATTRIBUTE_QUEUED
	printJobDirect      PrinterAttribute = 0x00000002 // PRINTER_ATTRIBUTE_DIRECT
	defaultPrinter      PrinterAttribute = 0x00000004 // PRINTER_ATTRIBUTE_DEFAULT
	sharedPrinter       PrinterAttribute = 0x00000008 // PRINTER_ATTRIBUTE_SHARED
	networkPrinter      PrinterAttribute = 0x00000010 // PRINTER_ATTRIBUTE_NETWORK
	hiddenPrinter       PrinterAttribute = 0x00000020 // PRINTER_ATTRIBUTE_HIDDEN
	localPrinter        PrinterAttribute = 0x00000040 // PRINTER_ATTRIBUTE_LOCAL
	enableDevQ          PrinterAttribute = 0x00000080 // PRINTER_ATTRIBUTE_ENABLE_DEVQ
	keepPrintedJobs     PrinterAttribute = 0x00000100 // PRINTER_ATTRIBUTE_KEEPPRINTEDJOBS
	doCompleteFirst     PrinterAttribute = 0x00000200 // PRINTER_ATTRIBUTE_DO_COMPLETE_FIRST
	workOffline         PrinterAttribute = 0x00000400 // PRINTER_ATTRIBUTE_WORK_OFFLINE
	enableBiDirectional PrinterAttribute = 0x00000800 // PRINTER_ATTRIBUTE_ENABLE_BIDI
	rawOnly             PrinterAttribute = 0x00001000 // PRINTER_ATTRIBUTE_RAW_ONLY
	published           PrinterAttribute = 0x00002000 // PRINTER_ATTRIBUTE_PUBLISHED
	fax                 PrinterAttribute = 0x00004000 // PRINTER_ATTRIBUTE_FAX
	ts                  PrinterAttribute = 0x00008000 // PRINTER_ATTRIBUTE_TS
	pushedUser          PrinterAttribute = 0x00020000 // PRINTER_ATTRIBUTE_PUSHED_USER
	pushedMachine       PrinterAttribute = 0x00040000 // PRINTER_ATTRIBUTE_PUSHED_MACHINE
	machine             PrinterAttribute = 0x00080000 // PRINTER_ATTRIBUTE_MACHINE
	friendlyName        PrinterAttribute = 0x00100000 // PRINTER_ATTRIBUTE_FRIENDLY_NAME

	//
	// If the redirected TS printer is installed with generic
	// TS printer driver (TSPRINT.dll) then this attribute is set
	// by the UMRDP service and passed on to the spooler
	//
	tsGenericDriver PrinterAttribute = 0x00200000 // PRINTER_ATTRIBUTE_TS_GENERIC_DRIVER
	//
	// Defines a bit allowing a local queue to be installed as a
	// "per-user" queue that is only usable/enumerable by that user.
	//
	perUser PrinterAttribute = 0x00400000 // PRINTER_ATTRIBUTE_PER_USER

	//
	// Defines a bit indicating that a queue is an enterprise
	// cloud print queue.
	//
	enterpriseCloud PrinterAttribute = 0x00800000 // PRINTER_ATTRIBUTE_ENTERPRISE_CLOUD
)

// Printer status values.
// See https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-2
// for the meaning of each value.
const (
	PRINTER_STATUS_PAUSED               = 0x00000001
	PRINTER_STATUS_ERROR                = 0x00000002
	PRINTER_STATUS_PENDING_DELETION     = 0x00000004
	PRINTER_STATUS_PAPER_JAM            = 0x00000008
	PRINTER_STATUS_PAPER_OUT            = 0x00000010
	PRINTER_STATUS_MANUAL_FEED          = 0x00000020
	PRINTER_STATUS_PAPER_PROBLEM        = 0x00000040
	PRINTER_STATUS_OFFLINE              = 0x00000080
	PRINTER_STATUS_IO_ACTIVE            = 0x00000100
	PRINTER_STATUS_BUSY                 = 0x00000200
	PRINTER_STATUS_PRINTING             = 0x00000400
	PRINTER_STATUS_OUTPUT_BIN_FULL      = 0x00000800
	PRINTER_STATUS_NOT_AVAILABLE        = 0x00001000
	PRINTER_STATUS_WAITING              = 0x00002000
	PRINTER_STATUS_PROCESSING           = 0x00004000
	PRINTER_STATUS_INITIALIZING         = 0x00008000
	PRINTER_STATUS_WARMING_UP           = 0x00010000
	PRINTER_STATUS_TONER_LOW            = 0x00020000
	PRINTER_STATUS_NO_TONER             = 0x00040000
	PRINTER_STATUS_PAGE_PUNT            = 0x00080000
	PRINTER_STATUS_USER_INTERVENTION    = 0x00100000
	PRINTER_STATUS_OUT_OF_MEMORY        = 0x00200000
	PRINTER_STATUS_DOOR_OPEN            = 0x00400000
	PRINTER_STATUS_SERVER_UNKNOWN       = 0x00800000
	PRINTER_STATUS_POWER_SAVE           = 0x01000000
	PRINTER_STATUS_SERVER_OFFLINE       = 0x02000000
	PRINTER_STATUS_DRIVER_UPDATE_NEEDED = 0x04000000
)

// PrinterInfo2 struct specifies detailed printer information.
// See https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-2 for the meanings
// of each field.
type PrinterInfo2 struct {
	serverName      *uint16
	printerName     *uint16
	shareName       *uint16
	portName        *uint16
	driverName      *uint16
	comment         *uint16
	location        *uint16
	DevMode         *PrinterDevMode
	sepFile         *uint16
	printProcessor  *uint16
	dataType        *uint16
	parameters      *uint16
	secDescriptor   *SecurityDescriptor
	attributes      PrinterAttribute
	priority        uint32
	defaultPriority uint32
	startTime       uint32
	untilTime       uint32
	status          uint32
	cJobs           uint32
	averagePPMs     uint32
}

// Attributes returns the printer attributes. It can be any combination of the bits defined
// in the constants beginning with PRINTER_ATTRIBUTE_ defined above.
func (pi2 *PrinterInfo2) Attributes() PrinterAttribute {
	return pi2.attributes
}

// Comment returns the comment associated with the printer.
func (pi2 *PrinterInfo2) Comment() string {
	return StringFromUTF16(pi2.comment)
}

// DataType returns the value of the dataType field. This is typically "RAW".
func (pi2 *PrinterInfo2) DataType() string {
	return StringFromUTF16(pi2.dataType)
}

// DriverName returns the name of the printer's driver.
func (pi2 *PrinterInfo2) DriverName() string {
	return StringFromUTF16(pi2.driverName)
}

// Location returns the location string for the printer.
func (pi2 *PrinterInfo2) Location() string {
	return StringFromUTF16(pi2.location)
}

// Name returns the name of the printer.
func (pi2 *PrinterInfo2) Name() string {
	return StringFromUTF16(pi2.printerName)
}

// Parameters returns the default print processor parameters.
func (pi2 *PrinterInfo2) Parameters() string {
	return StringFromUTF16(pi2.parameters)
}

// PortName returns the name of port(s) associated with the printer.
func (pi2 *PrinterInfo2) PortName() string {
	return StringFromUTF16(pi2.portName)
}

// PrintProcessor returns the name of the print processor used by the printer.
func (pi2 *PrinterInfo2) PrintProcessor() string {
	return StringFromUTF16(pi2.printProcessor)
}

// SepFile returns the name of the file used to create the separator page.
func (pi2 *PrinterInfo2) SeparatorFile() string {
	return StringFromUTF16(pi2.sepFile)
}

// ServerName returns the name of the print server associated with the printer.
func (pi2 *PrinterInfo2) ServerName() string {
	if pi2.serverName == nil {
		return "nil"
	} else {
		return StringFromUTF16(pi2.serverName)
	}
}

func (pi2 *PrinterInfo2) Status() uint32 {
	return pi2.status
}

// ShareName returns a string identifying the share point for the printer. The share point is
// only used if the PRINTER_ATTRIBUTE_SHARED constant was set for the Attributes member.
func (pi2 *PrinterInfo2) ShareName() string {
	if pi2.shareName == nil {
		return "nil"
	}
	return StringFromUTF16(pi2.shareName)
}

// Print prints out various values in the PrinterInfo2 struct.
func (pi2 *PrinterInfo2) Print() {
	fmt.Println("PrinterInfo2")
	fmt.Printf("    Server Name: %s\n", pi2.ServerName())
	fmt.Printf("    Printer Name: %s\n", pi2.Name())
	fmt.Printf("    Share Name: %s\n", pi2.ShareName())
	fmt.Printf("    Port Name: %s\n", pi2.PortName())
	fmt.Printf("    Driver Name: %s\n", pi2.DriverName())
	fmt.Printf("    Comment: %s\n", pi2.Comment())
	fmt.Printf("    Location: %s\n", pi2.Location())
	pi2.DevMode.Print()
	fmt.Printf("    Sep File: %s\n", pi2.SeparatorFile())
	fmt.Printf("    Print Processor: %s\n", pi2.PrintProcessor())
	fmt.Printf("    Data Type: %s\n", pi2.DataType())
	fmt.Printf("    Parameters: %s\n", pi2.Parameters())
	fmt.Println("    Attributes:")
	fmt.Println(parsePrinterAttributes(pi2.Attributes()))
	fmt.Printf("    Priority: %d\n", pi2.priority)
	fmt.Printf("    Default Priority: %d\n", pi2.defaultPriority)
	fmt.Printf("    Start Time: %d\n", pi2.startTime)
	fmt.Printf("    Until Time: %d\n", pi2.untilTime)
	fmt.Println("    Status:")
	fmt.Println(parsePrinterStatus(pi2.Status()))
	fmt.Printf("    Jobs: %d\n", pi2.cJobs)
	fmt.Printf("    Average PPMs: %d\n", pi2.averagePPMs)
}

func parsePrinterAttributes(attrs PrinterAttribute) string {
	var pAttrs strings.Builder
	if attrs&printJobQueued != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_QUEUED ")
	}
	if attrs&printJobDirect != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_DIRECT ")
	}
	if attrs&defaultPrinter != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_DEFAULT ")
	}
	if attrs&sharedPrinter != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_SHARED ")
	}
	if attrs&networkPrinter != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_NETWORK ")
	}
	if attrs&hiddenPrinter != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_HIDDEN ")
	}
	if attrs&localPrinter != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_LOCAL ")
	}
	if attrs&enableDevQ != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_ENABLE_DEVQ ")
	}
	if attrs&keepPrintedJobs != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_KEEPPRINTEDJOBS ")
	}
	if attrs&doCompleteFirst != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_DO_COMPLETE_FIRST ")
	}
	if attrs&workOffline != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_WORK_OFFLINE ")
	}
	if attrs&enableBiDirectional != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_ENABLE_BIDI ")
	}
	if attrs&rawOnly != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_RAW_ONLY ")
	}
	if attrs&published != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_PUBLISHED ")
	}
	if attrs&fax != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_FAX ")
	}
	if attrs&ts != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_TS ")
	}
	if attrs&pushedUser != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_PUSHED_USER ")
	}
	if attrs&pushedMachine != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_PUSHED_MACHINE ")
	}
	if attrs&machine != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_MACHINE ")
	}
	if attrs&friendlyName != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_FRIENDLY_NAME ")
	}
	if attrs&tsGenericDriver != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_TS_GENERIC_DRIVER ")
	}
	if attrs&perUser != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_PER_USER ")
	}
	if attrs&enterpriseCloud != 0 {
		pAttrs.WriteString("    PRINTER_ATTRIBUTE_ENTERPRISE_CLOUD ")
	}
	return pAttrs.String()
}

func parsePrinterStatus(status uint32) string {
	var pStatus strings.Builder
	if status&PRINTER_STATUS_PAUSED != 0 {
		pStatus.WriteString("    PRINTER_STATUS_PAUSED ")
	}
	if status&PRINTER_STATUS_ERROR != 0 {
		pStatus.WriteString("    PRINTER_STATUS_ERROR ")
	}
	if status&PRINTER_STATUS_PENDING_DELETION != 0 {
		pStatus.WriteString("    PRINTER_STATUS_PENDING_DELETION ")
	}
	if status&PRINTER_STATUS_PAPER_JAM != 0 {
		pStatus.WriteString("    PRINTER_STATUS_PAPER_JAM ")
	}
	if status&PRINTER_STATUS_PAPER_OUT != 0 {
		pStatus.WriteString("    PRINTER_STATUS_PAPER_OUT ")
	}
	if status&PRINTER_STATUS_MANUAL_FEED != 0 {
		pStatus.WriteString("    PRINTER_STATUS_MANUAL_FEED ")
	}
	if status&PRINTER_STATUS_PAPER_PROBLEM != 0 {
		pStatus.WriteString("    PRINTER_STATUS_PAPER_PROBLEM ")
	}
	if status&PRINTER_STATUS_OFFLINE != 0 {
		pStatus.WriteString("    PRINTER_STATUS_OFFLINE ")
	}
	if status&PRINTER_STATUS_IO_ACTIVE != 0 {
		pStatus.WriteString("    PRINTER_STATUS_IO_ACTIVE ")
	}
	if status&PRINTER_STATUS_BUSY != 0 {
		pStatus.WriteString("    PRINTER_STATUS_BUSY ")
	}
	if status&PRINTER_STATUS_PRINTING != 0 {
		pStatus.WriteString("    PRINTER_STATUS_PRINTING ")
	}
	if status&PRINTER_STATUS_OUTPUT_BIN_FULL != 0 {
		pStatus.WriteString("    PRINTER_STATUS_OUTPUT_BIN_FULL ")
	}
	if status&PRINTER_STATUS_NOT_AVAILABLE != 0 {
		pStatus.WriteString("    PRINTER_STATUS_NOT_AVAILABLE ")
	}
	if status&PRINTER_STATUS_WAITING != 0 {
		pStatus.WriteString("    PRINTER_STATUS_WAITING ")
	}
	if status&PRINTER_STATUS_PROCESSING != 0 {
		pStatus.WriteString("    PRINTER_STATUS_PROCESSING ")
	}
	if status&PRINTER_STATUS_INITIALIZING != 0 {
		pStatus.WriteString("    PRINTER_STATUS_INITIALIZING ")
	}
	if status&PRINTER_STATUS_WARMING_UP != 0 {
		pStatus.WriteString("    PRINTER_STATUS_WARMING_UP ")
	}
	if status&PRINTER_STATUS_TONER_LOW != 0 {
		pStatus.WriteString("    PRINTER_STATUS_TONER_LOW ")
	}
	if status&PRINTER_STATUS_NO_TONER != 0 {
		pStatus.WriteString("    PRINTER_STATUS_NO_TONER ")
	}
	if status&PRINTER_STATUS_PAGE_PUNT != 0 {
		pStatus.WriteString("    PRINTER_STATUS_PAGE_PUNT ")
	}
	if status&PRINTER_STATUS_USER_INTERVENTION != 0 {
		pStatus.WriteString("    PRINTER_STATUS_USER_INTERVENTION ")
	}
	if status&PRINTER_STATUS_OUT_OF_MEMORY != 0 {
		pStatus.WriteString("    PRINTER_STATUS_OUT_OF_MEMORY ")
	}
	if status&PRINTER_STATUS_DOOR_OPEN != 0 {
		pStatus.WriteString("    PRINTER_STATUS_DOOR_OPEN ")
	}
	if status&PRINTER_STATUS_SERVER_UNKNOWN != 0 {
		pStatus.WriteString("    PRINTER_STATUS_SERVER_UNKNOWN ")
	}
	if status&PRINTER_STATUS_POWER_SAVE != 0 {
		pStatus.WriteString("    PRINTER_STATUS_POWER_SAVE ")
	}
	if status&PRINTER_STATUS_SERVER_OFFLINE != 0 {
		pStatus.WriteString("    PRINTER_STATUS_SERVER_OFFLINE ")
	}
	if status&PRINTER_STATUS_DRIVER_UPDATE_NEEDED != 0 {
		pStatus.WriteString("    PRINTER_STATUS_DRIVER_UPDATE_NEEDED ")
	}
	return pStatus.String()
}
