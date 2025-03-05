//go:build windows

package print

import (
	"syscall"
	"unsafe"
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

	_, err := EnumPrinters(flags,
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
	_, err = EnumPrinters(flags,
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
		//		info2.Print()
		p := NewPrinter(&info2)
		printers.Add(p)
	}
	return printers, nil
}

// Add adds a printer to the Printers struct.
func (p *Printers) Add(printer *Printer) {
	p.Printers = append(p.Printers, *printer)
}

// getNames retrieves the names of all printers in the Printers struct.
func (p *Printers) getNames() []string {
	var names []string
	for _, printer := range p.Printers {
		names = append(names, printer.Name())
	}
	return names
}
