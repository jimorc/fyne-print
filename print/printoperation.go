package print

import (
	"fyne.io/fyne/v2"
)

var PrintOP *PrintOperation

// PrintOperation is the object that controls fyne print operations.
type PrintOperation struct {
	psd *PageSetupDialog
}

// NewPrintOperation creates a new PrintOperation object.
//
// Params:
//
//	window is the window that will contain the menu items for page setup and print.
func NewPrintOperation(window fyne.Window) *PrintOperation {
	printOp := &PrintOperation{}
	printOp.psd = NewPageSetupDialog(window, nil)

	return printOp
}

func (po *PrintOperation) PageSetupDialog() (*PageSetupDialog, error) {
	err := po.psd.populatePrinterSelect()
	return po.psd, err
}
