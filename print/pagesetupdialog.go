package print

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// PageSetupDialog is a ConfirmDialog dialog with widgets that must be saved.
type PageSetupDialog struct {
	*dialog.ConfirmDialog
	parent        fyne.Window
	printerSelect *widget.Select
}

// create a PageSetupDialog.
var pageSetupDialog *PageSetupDialog = &PageSetupDialog{}

// NewPageSetupDialog creates a PageSetupDialog.
//
// Params:
//
//	parent is the parent window for the dialog.
func NewPageSetupDialog(parent fyne.Window) *dialog.ConfirmDialog {
	pageSetupDialog.parent = parent
	printerContainer := pageSetupDialog.createPrinterContainer()

	box := container.NewVBox(printerContainer)

	pageSetupDialog.ConfirmDialog = dialog.NewCustomConfirm("PageSetup", "OK",
		"Cancel", box, func(bool) {}, parent)

	return pageSetupDialog.ConfirmDialog

}

// createPrinterContainer creates the container that holds the printers select and label.
func (psd *PageSetupDialog) createPrinterContainer() *fyne.Container {
	label := widget.NewLabel("Format For")
	psd.printerSelect = widget.NewSelect([]string{}, nil)
	psd.printerSelect.Resize(fyne.NewSize(250, psd.printerSelect.Size().Height))
	psd.populatePrinterSelect(psd.parent)
	c := container.NewBorder(nil, nil, label, psd.printerSelect)
	return c
}

func (psd *PageSetupDialog) populatePrinterSelect(parent fyne.Window) {
	printers, err := newPrinters()
	if err != nil {
		fyne.LogError("Error retrieving printers", err)
		err1 := errors.New("Error retrieving list of printers:\n" +
			err.Error() +
			"\nCannot continue page setup.")
		dialog.ShowError(err1, parent)
		return
	}
	prNames := printers.getNames()
	psd.printerSelect.Options = prNames
}
