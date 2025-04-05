package print

//#define UNICODE
//#include "windows.h"
import "C"
import (
	"errors"
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
)

var allForms []formInfo2

// Printer is a struct that allows access to a printer.
type Printer struct {
	pi2        PrinterInfo2
	handle     syscall.Handle
	dc         syscall.Handle
	forms      []formInfo2
	mediaNames []string
	mediaSizes []C.POINTL
	papers     []uint16
}

// newPrinter creates a Printer struct based on information provided in the PrinterInfo2 argument.
func newPrinter(pInfo2 *PrinterInfo2) *Printer {
	p := &Printer{pi2: *pInfo2}
	printerDefs := newPrinterDefaults("RAW", pInfo2.DevMode(),
		C.PRINTER_ACCESS_USE)

	prHandle := openPrinter(p.pi2.PrinterName(), printerDefs)
	p.handle = prHandle
	p.dc = createDC(p.pi2.PrinterName())
	p.forms = make([]formInfo2, 1)
	p.getMediaSizes()
	p.getPaperNames()
	p.getPaperSizes()
	p.getPapers()
	return p
}

// String returns a string representation of the Printer struct.
func (pr *Printer) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("    Handle: %d\n", pr.handle))
	s.WriteString(fmt.Sprintf("    DC: %d\n", pr.dc))
	s.WriteString(prepend("    ", pr.pi2.String()))
	s.WriteString(fmt.Sprintf("    For current paper size, width = %d\n", getDeviceCaps(pr.dc, C.PHYSICALWIDTH)))
	s.WriteString(fmt.Sprintf("    For current paper size, height = %d\n", getDeviceCaps(pr.dc, C.PHYSICALHEIGHT)))
	s.WriteString(fmt.Sprintf("    For current paper size, left offset = %d\n", getDeviceCaps(pr.dc, C.PHYSICALOFFSETX)))
	s.WriteString(fmt.Sprintf("    For current paper size, top offset = %d\n", getDeviceCaps(pr.dc, C.PHYSICALOFFSETY)))
	s.WriteString(fmt.Sprintf("    For current paper size, printable width = %d\n", getDeviceCaps(pr.dc, C.HORZRES)))
	s.WriteString(fmt.Sprintf("    For current paper size, printable height = %d\n", getDeviceCaps(pr.dc, C.VERTRES)))
	s.WriteString(fmt.Sprintf("    Printer has %d Media Sizes:\n", len(pr.forms)))
	s.WriteString("    Paper Names:\n")
	for i, f := range pr.forms {
		s.WriteString(fmt.Sprintf("        Form %d:\n", i))
		s.WriteString(prepend("          ", f.String()))
	}

	for i, n := range pr.mediaNames {
		s.WriteString(fmt.Sprintf("        Paper Name: %d:\n", i))
		ss := fmt.Sprintf("%s: %dx%d: %d", n, int32(pr.mediaSizes[i].x), int32(pr.mediaSizes[i].y), pr.papers[i])
		s.WriteString(prepend("          ", ss))
	}
	/*	for i, sz := range pr.mediaSizes {
		s.WriteString(fmt.Sprintf("        Paper Size: %d:\n", i))
		ss := fmt.Sprintf("%dx%d", int32(sz.x), int32(sz.y))
		s.WriteString(prepend("          ", ss))
	}*/
	return s.String()
}

// getMediaSizes retrieves the printer's formInfo2 objects. This is all of the
// media sizes that the printer might support.
func (p *Printer) getMediaSizes() {
	allForms = make([]formInfo2, 1)
	var cbBuf uint32 = uint32(unsafe.Sizeof(formInfo2{}))
	var needed uint32
	var returned uint32
	_, err := enumForms(p.handle, 2, allForms, 0, &needed, &returned)
	if err != nil && err != syscall.ERROR_INSUFFICIENT_BUFFER {
		fyne.LogError("Error getting media sizes: ", err)
		return
	}
	allForms = make([]formInfo2, needed/uint32(unsafe.Sizeof(formInfo2{})))
	cbBuf = needed
	_, err = enumForms(p.handle, 2, allForms, cbBuf, &needed, &returned)
	if err != nil {
		fyne.LogError("Error getting media sizes: ", err)
	}
	// The 'needed' parameter requested many more objects than are actually returned.
	// Therefore, the slice of objects should be reduced to the number actually returned.
	allForms = allForms[:returned]
	for _, fi2 := range allForms {
		fi := make([]byte, 1)
		needed = 1
		getForm(p.handle, fi2.formName(), &(fi[0]), &needed)

		fi = make([]byte, needed)
		err = getForm(p.handle, fi2.formName(), &(fi[0]), &needed)
		if err == nil {
			form := *(*formInfo2)(unsafe.Pointer(&fi[0]))
			p.forms = append(p.forms, form)
		}
	}
}

func (p *Printer) getPaperNames() error {
	// get number of paper names
	num, err := deviceCapabilities(p.pi2.PrinterName(),
		p.pi2.PortName(),
		C.DC_PAPERNAMES,
		0,
		p.pi2.DevMode())
	if num == -1 {
		if err == syscall.Errno(0) {
			err = errors.New("function unsupported, or general error")
		}
		fyne.LogError("Error getting paper names", err)
		return err
	}
	if num > 0 {
		// get paper names
		names := make([][64]uint16, num)
		num2, err := deviceCapabilities(p.pi2.PrinterName(),
			p.pi2.PortName(),
			C.DC_PAPERNAMES,
			uintptr(unsafe.Pointer(&names[0][0])),
			p.pi2.DevMode())
		if num2 != num {
			numErr := fmt.Errorf("returned paper names count (%d) does not match number available (%d)",
				num2, num)
			fyne.LogError("Error getting paper names", numErr)
			return err
		}
		p.mediaNames = make([]string, num2)
		for i := 0; i < int(num2); i++ {
			p.mediaNames[i] = syscall.UTF16ToString(names[i][:])
		}
	}
	return nil
}

func (p *Printer) getPaperSizes() error {
	// get number of paper sizes
	num, err := deviceCapabilities(p.pi2.PrinterName(),
		p.pi2.PortName(),
		C.DC_PAPERSIZE,
		0,
		p.pi2.DevMode())
	if num == -1 {
		if err == syscall.Errno(0) {
			err = errors.New("function unsupported, or general error")
		}
		fyne.LogError("Error getting paper sizes", err)
		return err
	}
	if num > 0 {
		// get paper sizes
		p.mediaSizes = make([]C.POINTL, num)
		num2, err := deviceCapabilities(p.pi2.PrinterName(),
			p.pi2.PortName(),
			C.DC_PAPERSIZE,
			uintptr(unsafe.Pointer(&p.mediaSizes[0])),
			p.pi2.DevMode())
		if num2 != num {
			numErr := fmt.Errorf("returned paper sizes count (%d) does not match number available (%d)",
				num2, num)
			fyne.LogError("Error getting paper names", numErr)
			return err
		}
	}
	return nil
}

func (p *Printer) getPapers() error {
	// get number of paper sizes
	num, err := deviceCapabilities(p.pi2.PrinterName(),
		p.pi2.PortName(),
		C.DC_PAPERS,
		0,
		p.pi2.DevMode())
	if num == -1 {
		if err == syscall.Errno(0) {
			err = errors.New("function unsupported, or general error")
		}
		fyne.LogError("Error getting paper sizes", err)
		return err
	}
	if num > 0 {
		// get paper sizes
		p.papers = make([]uint16, num)
		num2, err := deviceCapabilities(p.pi2.PrinterName(),
			p.pi2.PortName(),
			C.DC_PAPERS,
			uintptr(unsafe.Pointer(&p.papers[0])),
			p.pi2.DevMode())
		if num2 != num {
			numErr := fmt.Errorf("returned paper sizes count (%d) does not match number available (%d)",
				num2, num)
			fyne.LogError("Error getting paper names", numErr)
			return err
		}
	}
	return nil
}

// close cleans up Printer-related data such as the printer handle.
func (p *Printer) close() {
	// release the printer's device context (dc)
	if p.dc != 0 {
		deleteDC(p.dc)
		p.dc = 0
	}
	// release the printer handle.
	closePrinter(p.handle)
	p.handle = 0
}
