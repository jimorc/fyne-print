//go:build windows

package print

import (
	"syscall"
	"testing"
)

// Helper function to create a UTF16 pointer from a string (for test convenience)
func StringToUTF16Ptr(s string) *uint16 {
	ptr, _ := syscall.UTF16PtrFromString(s)
	return ptr
}

func TestPrinterInfo2_Attributes(t *testing.T) {
	pi2 := &PrinterInfo2{attributes: printJobQueued | printJobDirect}
	if pi2.Attributes() != printJobQueued|printJobDirect {
		t.Errorf("Expected attributes %d, but got %d", printJobQueued|printJobDirect, pi2.Attributes())
	}
}

func TestPrinterInfo2_Comment(t *testing.T) {
	comment := "Test Comment"
	pi2 := &PrinterInfo2{comment: StringToUTF16Ptr(comment)}
	if pi2.Comment() != comment {
		t.Errorf("Expected comment '%s', but got '%s'", comment, pi2.Comment())
	}
}

func TestPrinterInfo2_DataType(t *testing.T) {
	dataType := "RAW"
	pi2 := &PrinterInfo2{dataType: StringToUTF16Ptr(dataType)}
	if pi2.DataType() != dataType {
		t.Errorf("Expected data type '%s', but got '%s'", dataType, pi2.DataType())
	}
}

func TestPrinterInfo2_DriverName(t *testing.T) {
	driverName := "Test Driver"
	pi2 := &PrinterInfo2{driverName: StringToUTF16Ptr(driverName)}
	if pi2.DriverName() != driverName {
		t.Errorf("Expected driver name '%s', but got '%s'", driverName, pi2.DriverName())
	}
}

func TestPrinterInfo2_Location(t *testing.T) {
	location := "Test Location"
	pi2 := &PrinterInfo2{location: StringToUTF16Ptr(location)}
	if pi2.Location() != location {
		t.Errorf("Expected location '%s', but got '%s'", location, pi2.Location())
	}
}

func TestPrinterInfo2_Name(t *testing.T) {
	name := "Test Printer"
	pi2 := &PrinterInfo2{printerName: StringToUTF16Ptr(name)}
	if pi2.Name() != name {
		t.Errorf("Expected name '%s', but got '%s'", name, pi2.Name())
	}
}

func TestPrinterInfo2_Parameters(t *testing.T) {
	params := "Test Parameters"
	pi2 := &PrinterInfo2{parameters: StringToUTF16Ptr(params)}
	if pi2.Parameters() != params {
		t.Errorf("Expected parameters '%s', but got '%s'", params, pi2.Parameters())
	}
}

func TestPrinterInfo2_PortName(t *testing.T) {
	portName := "Test Port"
	pi2 := &PrinterInfo2{portName: StringToUTF16Ptr(portName)}
	if pi2.PortName() != portName {
		t.Errorf("Expected port name '%s', but got '%s'", portName, pi2.PortName())
	}
}

func TestPrinterInfo2_PrintProcessor(t *testing.T) {
	processor := "Test Processor"
	pi2 := &PrinterInfo2{printProcessor: StringToUTF16Ptr(processor)}
	if pi2.PrintProcessor() != processor {
		t.Errorf("Expected print processor '%s', but got '%s'", processor, pi2.PrintProcessor())
	}
}

func TestPrinterInfo2_SeparatorFile(t *testing.T) {
	sepFile := "Test Sep File"
	pi2 := &PrinterInfo2{sepFile: StringToUTF16Ptr(sepFile)}
	if pi2.SeparatorFile() != sepFile {
		t.Errorf("Expected separator file '%s', but got '%s'", sepFile, pi2.SeparatorFile())
	}
}

func TestPrinterInfo2_ServerName(t *testing.T) {
	serverName := "Test Server"
	pi2 := &PrinterInfo2{serverName: StringToUTF16Ptr(serverName)}
	if pi2.ServerName() != serverName {
		t.Errorf("Expected server name '%s', but got '%s'", serverName, pi2.ServerName())
	}

	pi2 = &PrinterInfo2{serverName: nil}
	if pi2.ServerName() != "nil" {
		t.Errorf("Expected server name 'nil', but got '%s'", pi2.ServerName())
	}
}

func TestPrinterInfo2_ShareName(t *testing.T) {
	shareName := "Test Share"
	pi2 := &PrinterInfo2{shareName: StringToUTF16Ptr(shareName)}
	if pi2.ShareName() != shareName {
		t.Errorf("Expected share name '%s', but got '%s'", shareName, pi2.ShareName())
	}
	pi2 = &PrinterInfo2{shareName: nil}
	if pi2.ShareName() != "nil" {
		t.Errorf("Expected share name 'nil', but got '%s'", pi2.ShareName())
	}
}

func TestPrinterInfo2_Status(t *testing.T) {
	pi2 := &PrinterInfo2{status: psPaused | psError}
	if pi2.Status() != psPaused|psError {
		t.Errorf("Expected attributes %d, but got %d", psPaused|psError, pi2.Status())
	}
}

func TestParsePrinterAttributes(t *testing.T) {
	tests := []struct {
		name     string
		attrs    PrinterAttribute
		expected string
	}{
		{
			name:     "No attributes",
			attrs:    0,
			expected: "",
		},
		{
			name:     "Queued",
			attrs:    printJobQueued,
			expected: "    PRINTER_ATTRIBUTE_QUEUED ",
		},
		{
			name:     "Direct",
			attrs:    printJobDirect,
			expected: "    PRINTER_ATTRIBUTE_DIRECT ",
		},
		{
			name:     "Default",
			attrs:    defaultPrinter,
			expected: "    PRINTER_ATTRIBUTE_DEFAULT ",
		},
		{
			name:     "Shared",
			attrs:    sharedPrinter,
			expected: "    PRINTER_ATTRIBUTE_SHARED ",
		},
		{
			name:     "Network",
			attrs:    networkPrinter,
			expected: "    PRINTER_ATTRIBUTE_NETWORK ",
		},
		{
			name:     "Hidden",
			attrs:    hiddenPrinter,
			expected: "    PRINTER_ATTRIBUTE_HIDDEN ",
		},
		{
			name:     "Local",
			attrs:    localPrinter,
			expected: "    PRINTER_ATTRIBUTE_LOCAL ",
		},
		{
			name:     "Enable DevQ",
			attrs:    enableDevQ,
			expected: "    PRINTER_ATTRIBUTE_ENABLE_DEVQ ",
		},
		{
			name:     "Keep Printed Jobs",
			attrs:    keepPrintedJobs,
			expected: "    PRINTER_ATTRIBUTE_KEEPPRINTEDJOBS ",
		},
		{
			name:     "Do Complete First",
			attrs:    doCompleteFirst,
			expected: "    PRINTER_ATTRIBUTE_DO_COMPLETE_FIRST ",
		},
		{
			name:     "Work Offline",
			attrs:    workOffline,
			expected: "    PRINTER_ATTRIBUTE_WORK_OFFLINE ",
		},
		{
			name:     "Enable Bidi",
			attrs:    enableBiDirectional,
			expected: "    PRINTER_ATTRIBUTE_ENABLE_BIDI ",
		},
		{
			name:     "Raw Only",
			attrs:    rawOnly,
			expected: "    PRINTER_ATTRIBUTE_RAW_ONLY ",
		},
		{
			name:     "Published",
			attrs:    published,
			expected: "    PRINTER_ATTRIBUTE_PUBLISHED ",
		},
		{
			name:     "Fax",
			attrs:    fax,
			expected: "    PRINTER_ATTRIBUTE_FAX ",
		},
		{
			name:     "TS",
			attrs:    ts,
			expected: "    PRINTER_ATTRIBUTE_TS ",
		},
		{
			name:     "Pushed User",
			attrs:    pushedUser,
			expected: "    PRINTER_ATTRIBUTE_PUSHED_USER ",
		},
		{
			name:     "Pushed Machine",
			attrs:    pushedMachine,
			expected: "    PRINTER_ATTRIBUTE_PUSHED_MACHINE ",
		},
		{
			name:     "Machine",
			attrs:    machine,
			expected: "    PRINTER_ATTRIBUTE_MACHINE ",
		},
		{
			name:     "Friendly Name",
			attrs:    friendlyName,
			expected: "    PRINTER_ATTRIBUTE_FRIENDLY_NAME ",
		},
		{
			name:     "TS Generic Driver",
			attrs:    tsGenericDriver,
			expected: "    PRINTER_ATTRIBUTE_TS_GENERIC_DRIVER ",
		},
		{
			name:     "Per User",
			attrs:    perUser,
			expected: "    PRINTER_ATTRIBUTE_PER_USER ",
		},
		{
			name:     "Enterprise Cloud",
			attrs:    enterpriseCloud,
			expected: "    PRINTER_ATTRIBUTE_ENTERPRISE_CLOUD ",
		},
		{
			name:     "Multiple Attributes",
			attrs:    printJobQueued | printJobDirect | sharedPrinter,
			expected: "    PRINTER_ATTRIBUTE_QUEUED     PRINTER_ATTRIBUTE_DIRECT     PRINTER_ATTRIBUTE_SHARED ",
		},
		{
			name:     "All Attributes",
			attrs:    printJobQueued | printJobDirect | defaultPrinter | sharedPrinter | networkPrinter | hiddenPrinter | localPrinter | enableBiDirectional | keepPrintedJobs | doCompleteFirst | workOffline | enableBiDirectional | rawOnly | published | fax | ts | pushedUser | pushedMachine | machine | friendlyName | tsGenericDriver | perUser | enterpriseCloud,
			expected: "    PRINTER_ATTRIBUTE_QUEUED     PRINTER_ATTRIBUTE_DIRECT     PRINTER_ATTRIBUTE_DEFAULT     PRINTER_ATTRIBUTE_SHARED     PRINTER_ATTRIBUTE_NETWORK     PRINTER_ATTRIBUTE_HIDDEN     PRINTER_ATTRIBUTE_LOCAL     PRINTER_ATTRIBUTE_KEEPPRINTEDJOBS     PRINTER_ATTRIBUTE_DO_COMPLETE_FIRST     PRINTER_ATTRIBUTE_WORK_OFFLINE     PRINTER_ATTRIBUTE_ENABLE_BIDI     PRINTER_ATTRIBUTE_RAW_ONLY     PRINTER_ATTRIBUTE_PUBLISHED     PRINTER_ATTRIBUTE_FAX     PRINTER_ATTRIBUTE_TS     PRINTER_ATTRIBUTE_PUSHED_USER     PRINTER_ATTRIBUTE_PUSHED_MACHINE     PRINTER_ATTRIBUTE_MACHINE     PRINTER_ATTRIBUTE_FRIENDLY_NAME     PRINTER_ATTRIBUTE_TS_GENERIC_DRIVER     PRINTER_ATTRIBUTE_PER_USER     PRINTER_ATTRIBUTE_ENTERPRISE_CLOUD ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parsePrinterAttributes(tt.attrs)
			if result != tt.expected {
				t.Errorf("parsePrinterAttributes(%d) = '%s', expected '%s'", tt.attrs, result, tt.expected)
			}
		})
	}
}

func TestPrinterInfo2_string(t *testing.T) {
	tests := []struct {
		name     string
		pi2      *PrinterInfo2
		expected string
	}{
		{
			name: "Basic Print",
			pi2: &PrinterInfo2{
				serverName:  StringToUTF16Ptr("Test Server"),
				printerName: StringToUTF16Ptr("Test Printer"),
				shareName:   StringToUTF16Ptr("Test Share"),
				portName:    StringToUTF16Ptr("Test Port"),
				driverName:  StringToUTF16Ptr("Test Driver"),
				comment:     StringToUTF16Ptr("Test Comment"),
				location:    StringToUTF16Ptr("Test Location"),
				DevMode: &PrinterDevMode{
					dmDeviceName: [32]uint16{0x41, 0x42, 0x43}, //"ABC"
				},
				sepFile:         StringToUTF16Ptr("Test Sep File"),
				printProcessor:  StringToUTF16Ptr("Test Processor"),
				dataType:        StringToUTF16Ptr("RAW"),
				parameters:      StringToUTF16Ptr("Test Parameters"),
				attributes:      printJobQueued | printJobDirect,
				priority:        1,
				defaultPriority: 1,
				startTime:       0,
				untilTime:       0,
				status:          psPaused | psError,
				cJobs:           1,
				averagePPMs:     1,
			},
			expected: `PrinterInfo2
    Server Name: Test Server
    Printer Name: Test Printer
    Share Name: Test Share
    Port Name: Test Port
    Driver Name: Test Driver
    Comment: Test Comment
    Location: Test Location
PrinterDevMode:
    Device Name: ABC
    SpecVersion: 0
    Driver Version: 0
    Size: 0
    Driver Extra: 0
    Fields: 0
    Orientation: 0
    PaperSize: 0
    Paper Length: 0
    Paper Width: 0
    Scale: 0
    Copies: 0
    Default Source: 0
    Print Quality: 0
    Color: 0
    Duplex: 0
    Y-Resolution: 0
    TT Option: 0
    Collate: 0
    Form Name: 
    Logical Pixels: 0
    Bits Per Pel: 0
    Pels Width: 0
    Pels Height: 0
    Where NUP is done: 0
    Display Frequency: 0
    ICM Method: 0
    ICM Intent: 0
    Media Type: 0
    Dither Type: 0
    Panning Width: 0
    Panning Height: 0
    Sep File: Test Sep File
    Print Processor: Test Processor
    Data Type: RAW
    Parameters: Test Parameters
    Attributes:
    PRINTER_ATTRIBUTE_QUEUED     PRINTER_ATTRIBUTE_DIRECT 
    Priority: 1
    Default Priority: 1
    Start Time: 0
    Until Time: 0
    Status:
    PRINTER_STATUS_PAUSED     PRINTER_STATUS_ERROR 
    Jobs: 1
    Average PPMs: 1
`,
		},
		{
			name: "Empty strings",
			pi2: &PrinterInfo2{
				serverName:      StringToUTF16Ptr(""),
				printerName:     StringToUTF16Ptr(""),
				shareName:       StringToUTF16Ptr(""),
				portName:        StringToUTF16Ptr(""),
				driverName:      StringToUTF16Ptr(""),
				comment:         StringToUTF16Ptr(""),
				location:        StringToUTF16Ptr(""),
				DevMode:         &PrinterDevMode{},
				sepFile:         StringToUTF16Ptr(""),
				printProcessor:  StringToUTF16Ptr(""),
				dataType:        StringToUTF16Ptr(""),
				parameters:      StringToUTF16Ptr(""),
				attributes:      0,
				priority:        0,
				defaultPriority: 0,
				startTime:       0,
				untilTime:       0,
				status:          0,
				cJobs:           0,
				averagePPMs:     0,
			},
			expected: `PrinterInfo2
    Server Name: 
    Printer Name: 
    Share Name: 
    Port Name: 
    Driver Name: 
    Comment: 
    Location: 
PrinterDevMode:
    Device Name: 
    SpecVersion: 0
    Driver Version: 0
    Size: 0
    Driver Extra: 0
    Fields: 0
    Orientation: 0
    PaperSize: 0
    Paper Length: 0
    Paper Width: 0
    Scale: 0
    Copies: 0
    Default Source: 0
    Print Quality: 0
    Color: 0
    Duplex: 0
    Y-Resolution: 0
    TT Option: 0
    Collate: 0
    Form Name: 
    Logical Pixels: 0
    Bits Per Pel: 0
    Pels Width: 0
    Pels Height: 0
    Where NUP is done: 0
    Display Frequency: 0
    ICM Method: 0
    ICM Intent: 0
    Media Type: 0
    Dither Type: 0
    Panning Width: 0
    Panning Height: 0
    Sep File: 
    Print Processor: 
    Data Type: 
    Parameters: 
    Attributes:

    Priority: 0
    Default Priority: 0
    Start Time: 0
    Until Time: 0
    Status:

    Jobs: 0
    Average PPMs: 0
`,
		},
		{
			name: "nil Server and Share Name",
			pi2: &PrinterInfo2{
				serverName:      nil,
				printerName:     StringToUTF16Ptr("Test Printer"),
				shareName:       nil,
				portName:        StringToUTF16Ptr("Test Port"),
				driverName:      StringToUTF16Ptr("Test Driver"),
				comment:         StringToUTF16Ptr("Test Comment"),
				location:        StringToUTF16Ptr("Test Location"),
				DevMode:         &PrinterDevMode{},
				sepFile:         StringToUTF16Ptr("Test Sep File"),
				printProcessor:  StringToUTF16Ptr("Test Processor"),
				dataType:        StringToUTF16Ptr("RAW"),
				parameters:      StringToUTF16Ptr("Test Parameters"),
				attributes:      networkPrinter,
				priority:        1,
				defaultPriority: 1,
				startTime:       0,
				untilTime:       0,
				status:          psOffline,
				cJobs:           1,
				averagePPMs:     1,
			},
			expected: `PrinterInfo2
    Server Name: nil
    Printer Name: Test Printer
    Share Name: nil
    Port Name: Test Port
    Driver Name: Test Driver
    Comment: Test Comment
    Location: Test Location
PrinterDevMode:
    Device Name: 
    SpecVersion: 0
    Driver Version: 0
    Size: 0
    Driver Extra: 0
    Fields: 0
    Orientation: 0
    PaperSize: 0
    Paper Length: 0
    Paper Width: 0
    Scale: 0
    Copies: 0
    Default Source: 0
    Print Quality: 0
    Color: 0
    Duplex: 0
    Y-Resolution: 0
    TT Option: 0
    Collate: 0
    Form Name: 
    Logical Pixels: 0
    Bits Per Pel: 0
    Pels Width: 0
    Pels Height: 0
    Where NUP is done: 0
    Display Frequency: 0
    ICM Method: 0
    ICM Intent: 0
    Media Type: 0
    Dither Type: 0
    Panning Width: 0
    Panning Height: 0
    Sep File: Test Sep File
    Print Processor: Test Processor
    Data Type: RAW
    Parameters: Test Parameters
    Attributes:
    PRINTER_ATTRIBUTE_NETWORK 
    Priority: 1
    Default Priority: 1
    Start Time: 0
    Until Time: 0
    Status:
    PRINTER_STATUS_OFFLINE 
    Jobs: 1
    Average PPMs: 1
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdout to a buffer
			//			old := os.Stdout
			//			r, w, _ := os.Pipe()
			//			os.Stdout = w

			// Call the Print method
			pi2String := tt.pi2.string()
			if pi2String != tt.expected {
				t.Errorf("string() output mismatch for %s", tt.name)
			}

			// Close the write end of the pipe and restore stdout
			/*			w.Close()
						//			os.Stdout = old

									// Read the captured output
									var buf bytes.Buffer
									if _, err := buf.ReadFrom(r); err != nil {
										t.Fatalf("Failed to read from buffer: %v", err)
									}
									output := buf.String()

									// Trim leading and trailing whitespace from both output and expected
									output = strings.TrimSpace(output)
									expected := strings.TrimSpace(tt.expected)

									if output != expected {
										fmt.Println("Actual:\n", output)
										fmt.Println("Expected:\n", expected)
										t.Errorf("string() output mismatch for %s", tt.name)
									} */
		})
	}
}
