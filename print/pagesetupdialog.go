package print

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	xlayout "fyne.io/x/fyne/layout"
)

// PageSetupDialog is a ConfirmDialog dialog with widgets that must be saved.
type PageSetupDialog struct {
	*dialog.ConfirmDialog
	parent                fyne.Window
	printerSelect         *widget.Select
	paperSizeSelect       *widget.Select
	orientationRadioGroup *widget.RadioGroup
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

	//	box := container.NewVBox(printerContainer)

	pageSetupDialog.ConfirmDialog = dialog.NewCustomConfirm("PageSetup", "OK",
		"Cancel", printerContainer, func(bool) {}, parent)
	pageSetupDialog.Resize(fyne.NewSize(500, 300))
	return pageSetupDialog.ConfirmDialog

}

// createPrinterContainer creates the container that holds the printers select and label.
func (psd *PageSetupDialog) createPrinterContainer() *fyne.Container {
	prLabel := widget.NewLabel("Format For")
	psd.printerSelect = widget.NewSelect([]string{}, nil)
	psd.printerSelect.Alignment = fyne.TextAlignTrailing
	psLabel := widget.NewLabel("Paper Size")
	psd.paperSizeSelect = widget.NewSelect([]string{}, nil)
	psd.paperSizeSelect.Alignment = fyne.TextAlignTrailing
	orLabel := widget.NewLabel("Orientation")
	psd.orientationRadioGroup = widget.NewRadioGroup([]string{"Portrait", "Landscape"}, nil)
	psd.orientationRadioGroup.Horizontal = true
	psd.populatePrinterSelect(psd.parent)
	prC := container.New(xlayout.NewHPortion([]float64{30, 70}), prLabel, psd.printerSelect)
	psC := container.New(xlayout.NewHPortion([]float64{30, 70}), psLabel, psd.paperSizeSelect)
	orC := container.New(xlayout.NewHPortion([]float64{30, 70}), orLabel, psd.orientationRadioGroup)
	box := container.NewVBox(prC, psC, orC)
	return box
}

func (psd *PageSetupDialog) populatePrinterSelect(parent fyne.Window) {
	printers, err := NewPrinters()
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
