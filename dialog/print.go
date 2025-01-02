package print

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/jimorc/fyne-print/printer"
)

// PrintDialog implements the print dialog
type PrintDialog struct {
	dialog       *dialog.CustomDialog
	parent       *fyne.Window
	pSelect      *widget.Select
	location     *widget.Label
	printerModel *widget.Label
	uri          *widget.Label

	printers            printer.Printers
	activePrinterNumber int
}

// NewPrintDialog creates a new PrintDialog.
func NewPrintDialog(parent fyne.Window) *PrintDialog {
	pDialog := &PrintDialog{parent: &parent}

	printerLabel := widget.NewLabel("Printer:")
	pDialog.pSelect = widget.NewSelect([]string{}, pDialog.printerChanged)
	locLabel := widget.NewLabel("Location:")
	pDialog.location = widget.NewLabel("")
	modelLabel := widget.NewLabel("Type:")
	modelLabel.Resize(fyne.NewSize(modelLabel.Size().Width, 12))
	pDialog.printerModel = widget.NewLabel("")
	uriLabel := widget.NewLabel("URI:")
	pDialog.uri = widget.NewLabel("")
	pBox := container.New(layout.NewFormLayout(), printerLabel, pDialog.pSelect,
		locLabel, pDialog.location, modelLabel, pDialog.printerModel, uriLabel, pDialog.uri)
	box := container.NewVBox(pBox)
	printerCard := widget.NewCard("", "", box)
	bOptions := widget.NewButton("Options >>", optionsClicked)
	bOptions.Importance = widget.LowImportance
	bCancel := widget.NewButton("Cancel", pDialog.cancelClicked)
	bCancel.Importance = widget.MediumImportance
	bPrint := widget.NewButton("Print", printClicked)
	bPrint.Importance = widget.HighImportance
	printCancelBox := container.NewHBox(bCancel, bPrint)
	buttonBox := container.NewBorder(nil, nil, bOptions, printCancelBox)
	printerBox := container.NewVBox(printerCard, buttonBox)
	pDialog.dialog = dialog.NewCustomWithoutButtons("Print", printerBox, parent)
	pDialog.dialog.Resize(fyne.NewSize(600, pDialog.dialog.MinSize().Height))

	return pDialog
}

// Show displays the PrintDialog.
func (pd *PrintDialog) Show() {
	printers, err := printer.NewPrinters()
	if err != nil {
		error := errors.New("Error detected while trying to retrieve list of available printers:\n" +
			err.Error())
		dialog.ShowError(error, *pd.parent)
	}
	pd.printers = *printers
	pd.pSelect.Options = pd.printers.GetNames()

	defaultPrinter, err := printers.GetDefaultPrinter()
	pd.activePrinterNumber = defaultPrinter
	if err != nil {
		error := errors.New("Error detected while trying to retrieve the default printer:\n" +
			err.Error())
		dialog.ShowError(error, *pd.parent)
	}
	pd.pSelect.SetSelectedIndex(defaultPrinter)
	pd.dialog.Show()
}

func optionsClicked() {}

func printClicked() {}

func (pD *PrintDialog) cancelClicked() {
	pD.dialog.Hide()
}

func (pD *PrintDialog) printerChanged(printerName string) {
	index, err := pD.printers.GetPrinterIndexByName((printerName))
	if err != nil {
		error := errors.New("Error detected attempting to retrieve the selected printer:\n" +
			err.Error())
		dialog.ShowError(error, *pD.parent)
		return
	}
	pD.activePrinterNumber = index
	pD.location.Text = pD.printers.Printers[pD.activePrinterNumber].Location
	pD.printerModel.Text = pD.printers.Printers[pD.activePrinterNumber].Model
	pD.uri.Text = pD.printers.Printers[pD.activePrinterNumber].Uri
	pD.location.Refresh()
	pD.printerModel.Refresh()
}
