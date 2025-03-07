package print

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

var PrintOP *PrintOperation

// PrintOperation is the object that controls fyne print operations.
type PrintOperation struct {
	pageSetupDialog *dialog.ConfirmDialog
}

// NewPrintOperation creates a new PrintOperation object.
//
// Params:
//
//	window is the window that will contain the menu items for page setup and print.
func NewPrintOperation(window fyne.Window) *PrintOperation {
	printOp := &PrintOperation{}
	printOp.pageSetupDialog = NewPageSetupDialog(window, nil)

	return printOp
}

func (po *PrintOperation) PageSetupDialog() *dialog.ConfirmDialog {
	return po.pageSetupDialog
}
