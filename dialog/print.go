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

// InfoDialog implements the printer info dialog.
// It is displayed when the Info button on the Print Dialog is clicked.
type InfoDialog struct {
	dialog       *dialog.CustomDialog
	location     *widget.Label
	printerModel *widget.Label
	uri          *widget.Label
}

// PrintDialog implements the print dialog
type PrintDialog struct {
	dialog     *dialog.CustomDialog
	parent     *fyne.Window
	infoDialog *InfoDialog
	pSelect    *widget.Select

	printers            printer.Printers
	activePrinterNumber int
}

// NewPrintDialog creates a new PrintDialog.
func NewPrintDialog(parent fyne.Window) *PrintDialog {
	pDialog := &PrintDialog{parent: &parent}

	printerLabel := widget.NewLabel("Printer:")
	pDialog.pSelect = widget.NewSelect([]string{}, pDialog.printerChanged)
	infoButton := widget.NewButton("Info", pDialog.infoClicked)
	pSelInfo := container.NewBorder(nil, nil, nil, infoButton, pDialog.pSelect)
	pBox := container.New(layout.NewFormLayout(), printerLabel, pSelInfo)
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

func (pD *PrintDialog) infoClicked() {
	pD.infoDialog = pD.makeInfoDialog()
	pD.infoDialog.dialog.Show()
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
}

func (pD *PrintDialog) makeInfoDialog() *InfoDialog {
	iD := &InfoDialog{}
	printerLabel := widget.NewLabel("Printer:")
	printerName := widget.NewLabel(pD.pSelect.Selected)
	locLabel := widget.NewLabel("Location:")
	iD.location = widget.NewLabel("")
	modelLabel := widget.NewLabel("Type:")
	//	iD.modelLabel.Resize(fyne.NewSize(modelLabel.Size().Width, 12))
	iD.printerModel = widget.NewLabel("")
	uriLabel := widget.NewLabel("URI:")
	iD.uri = widget.NewLabel("")
	infoBox := container.New(layout.NewFormLayout(), printerLabel, printerName,
		locLabel, iD.location, modelLabel, iD.printerModel,
		uriLabel, iD.uri)

	iD.location.Text = pD.printers.Printers[pD.activePrinterNumber].Location
	iD.printerModel.Text = pD.printers.Printers[pD.activePrinterNumber].Model
	iD.uri.Text = pD.printers.Printers[pD.activePrinterNumber].Uri

	iD.dialog = dialog.NewCustom("Printer Information", "OK", infoBox, *pD.parent)
	return iD
}
