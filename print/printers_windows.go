package print

import (
	"fmt"
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
)

// Printers contains a slice of Printer structs. This typically contains a slice of all
// printers available on the computer.
type Printers struct {
	Printers []Printer
}

// NewPrinters retrieves a Printers struct containing printers configured on the computing device.
func NewPrinters() (*Printers, error) {
	flags := PRINTER_ENUM_LOCAL |
		PRINTER_ENUM_CONNECTIONS
		// buffer is a slice of bytes that will contain an array of PrinterInfo1 structs
	var buffer = make([]byte, 1)
	// size of *info1 in bytes
	var info2Size uint32 = 0
	// number of bytes needed in *info1
	var info2Needed uint32 = 0
	// number of info1 structs returned.
	var info2Count uint32 = 0

	_, err := enumPrinters(flags,
		"",
		2,
		&buffer[0],
		info2Size,
		&info2Needed,
		&info2Count)

	if err != syscall.ERROR_INSUFFICIENT_BUFFER {
		return nil, err
	}
	info2Size = info2Needed
	buffer = make([]byte, info2Size)
	_, err = enumPrinters(flags,
		"",
		2,
		&buffer[0],
		info2Size,
		&info2Needed,
		&info2Count)
	if err != syscall.Errno(0) {
		return nil, err
	}
	pInfo2 := (*[1024]PrinterInfo2)(unsafe.Pointer(&buffer[0]))[:info2Count:info2Count]
	printers := &Printers{}
	for _, info2 := range pInfo2 {
		fmt.Println(info2.string())
		p := newPrinter(&info2)
		printers.Add(p)
	}
	return printers, nil
}

// Add adds a printer to the Printers struct.
func (p *Printers) Add(printer *Printer) {
	p.Printers = append(p.Printers, *printer)
}

func (p *Printers) getPrinterByName(name string) *Printer {
	for _, printer := range p.Printers {
		if printer.Name() == name {
			return &printer
		}
	}
	return nil
}

// DefaultPrinter returns the system's default Printer, or nil on error or no default.
func (p *Printers) DefaultPrinter() *Printer {
	buf := make([]uint16, 3)
	bufN := uint32(len(buf))
	err := getDefaultPrinter(&buf[0], &bufN)
	if err != syscall.ERROR_INSUFFICIENT_BUFFER {
		fyne.LogError("Error getting default printer", err)
		return nil
	}
	buf = make([]uint16, bufN)
	err = getDefaultPrinter(&buf[0], &bufN)
	if err != syscall.Errno(0) {
		fyne.LogError("Error getting default printer", err)
		return nil
	}
	for _, printer := range p.Printers {
		if printer.Name() == StringFromUTF16(&buf[0]) {
			return &printer
		}
	}
	return nil
}

// getNames retrieves the names of all printers in the Printers struct.
func (p *Printers) getNames() []string {
	var names []string
	for _, printer := range p.Printers {
		names = append(names, printer.Name())
	}
	return names
}
