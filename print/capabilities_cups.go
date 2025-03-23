//go:build !windows

package print

//#include "cups/cups.h"
import "C"

// capabilities contains bits that define a printer's capabilities.
type capabilities uint32

// AsString outputs a string slice with one string for each set capability.
func (c capabilities) AsStrings() []string {
	caps := make([]string, 0)
	if c.CanBind() {
		caps = append(caps, "CanBind")
	}
	if c.CanCollate() {
		caps = append(caps, "CanCollate")
	}
	if c.CanCover() {
		caps = append(caps, "CanCover")
	}
	if c.CanDoCopies() {
		caps = append(caps, "CanDoCopies")
	}
	if c.CanDuplex() {
		caps = append(caps, "CanDuplex")
	}
	if c.CanPrintBnW() {
		caps = append(caps, "CanPrintBnW")
	}
	if c.CanPrintColor() {
		caps = append(caps, "CanPrintColor")
	}
	if c.CanPrintVariable() {
		caps = append(caps, "CanPrintVariable")
	}
	if c.CanPrintLarge() {
		caps = append(caps, "CanPrintLarge")
	}
	if c.CanPrintMedium() {
		caps = append(caps, "CanPrintMedium")
	}
	if c.CanPrintSmall() {
		caps = append(caps, "CanPrintSmall")
	}
	if c.CanPunch() {
		caps = append(caps, "CanPunch")
	}
	if c.CanSort() {
		caps = append(caps, "CanSort")
	}
	if c.IsDefault() {
		caps = append(caps, "IsDefault")
	}
	if c.IsFax() {
		caps = append(caps, "IsFax")
	}
	if c.IsLocal() {
		caps = append(caps, "IsLocal")
	}
	if c.IsRejectingJobs() {
		caps = append(caps, "IsRejectingJobs")
	}
	if c.IsRemote() {
		caps = append(caps, "IsRemote")
	}
	if c.IsShared() {
		caps = append(caps, "IsShared")
	}
	return caps
}

// CanBind returns true if the printer supports binding.
func (c capabilities) CanBind() bool {
	return c&C.CUPS_PRINTER_BIND == C.CUPS_PRINTER_BIND
}

// CanCollate returns true if the printer can collate copies
func (c capabilities) CanCollate() bool {
	return c&C.CUPS_PRINTER_COLLATE == C.CUPS_PRINTER_COLLATE
}

// CanCover returns true if the printer can print a cover page.
func (c capabilities) CanCover() bool {
	return c&C.CUPS_PRINTER_COVER == C.CUPS_PRINTER_COVER
}

// CanDoCopies returns true if the printer supports printing multiple copies.
func (c capabilities) CanDoCopies() bool {
	return c&C.CUPS_PRINTER_COPIES == C.CUPS_PRINTER_COPIES
}

// CanDuplex returns true if the printer supports two-sided printing.
func (c capabilities) CanDuplex() bool {
	return c&C.CUPS_PRINTER_DUPLEX == C.CUPS_PRINTER_DUPLEX
}

// CanPrintBnW returns true if the printer supports B&W printing.
func (c capabilities) CanPrintBnW() bool {
	return c&C.CUPS_PRINTER_BW == C.CUPS_PRINTER_BW
}

// CanPrintColor returns true if the printer supports printing in color.
func (c capabilities) CanPrintColor() bool {
	return c&C.CUPS_PRINTER_COLOR == C.CUPS_PRINTER_COLOR
}

// CanPrintVariable returns true if the printer supports printing on
// rolls or custom-sized media.
func (c capabilities) CanPrintVariable() bool {
	return c&C.CUPS_PRINTER_VARIABLE == C.CUPS_PRINTER_VARIABLE
}

// CanPrintLarge returns true if the printer cna print on large pages
// such as D/E/A1/A0.
func (c capabilities) CanPrintLarge() bool {
	return c&C.CUPS_PRINTER_LARGE == C.CUPS_PRINTER_LARGE
}

// CanPrintMedium returns true if the printer can print on medium pages
// such as Tabloid/B/C/A3/A2.
func (c capabilities) CanPrintMedium() bool {
	return c&C.CUPS_PRINTER_MEDIUM == C.CUPS_PRINTER_MEDIUM
}

// CanPrintSmall returns true if the printer can print on small pages
// such as Letter/Legal/A4.
func (c capabilities) CanPrintSmall() bool {
	return c&C.CUPS_PRINTER_SMALL == C.CUPS_PRINTER_SMALL
}

// CanPunch returns true if the printer supports punching the output.
func (c capabilities) CanPunch() bool {
	return c&C.CUPS_PRINTER_PUNCH == C.CUPS_PRINTER_PUNCH
}

// CanSort returns true if the printer supports sorting the output.
func (c capabilities) CanSort() bool {
	return c&C.CUPS_PRINTER_SORT == C.CUPS_PRINTER_SORT
}

// IsDefault returns true is the printer is the network default
// printer. This is not necessarily the same printer as indicated
// by the Printer.IsDefault method.
func (c capabilities) IsDefault() bool {
	return c&C.CUPS_PRINTER_DEFAULT == C.CUPS_PRINTER_DEFAULT
}

// IsFax returns true if the printer is a fax.
func (c capabilities) IsFax() bool {
	return c&C.CUPS_PRINTER_FAX == C.CUPS_PRINTER_FAX
}

// IsLocal returns true if the printer is a local printer.
func (c capabilities) IsLocal() bool {
	return c&C.CUPS_PRINTER_LOCAL == C.CUPS_PRINTER_LOCAL
}

// IsRejectingJobs returns true if the printer is currently
// rejecting jobs.
func (c capabilities) IsRejectingJobs() bool {
	return c&C.CUPS_PRINTER_REJECTING == C.CUPS_PRINTER_REJECTING
}

// IsRemote returns true if the printer is a remote printer.
func (c capabilities) IsRemote() bool {
	return c&C.CUPS_PRINTER_REMOTE == C.CUPS_PRINTER_REMOTE
}

// IsShared returns true if the printer is shared.
func (c capabilities) IsShared() bool {
	return c&C.CUPS_PRINTER_NOT_SHARED == 0
}
