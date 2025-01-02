package print

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/jimorc/fyne-print/printer"
)

// PrintDialog implements the print dialog
type PrintDialog struct {
	dialog  *dialog.CustomDialog
	pSelect *widget.Select
}

// NewPrintDialog creates a new PrintDialog.
func NewPrintDialog(parent fyne.Window) *PrintDialog {
	pDialog := &PrintDialog{}

	printerLabel := widget.NewLabel("Printer:")
	pDialog.pSelect = widget.NewSelect([]string{}, printerChanged)
	locLabel := widget.NewLabel("Location:")
	location := widget.NewLabel("")
	typeLabel := widget.NewLabel("Type:")
	typeLabel.Resize(fyne.NewSize(typeLabel.Size().Width, 12))
	printerType := widget.NewLabel("")
	pBox := container.New(layout.NewFormLayout(), printerLabel, pDialog.pSelect,
		locLabel, location, typeLabel, printerType)
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
func (pd PrintDialog) Show() {
	printers, err := printer.GetPrinters()
	if err != nil {
		// show error message
	}
	pd.pSelect.Options = printers
	pd.dialog.Show()
}

func optionsClicked() {}

func printClicked() {}

func (pD *PrintDialog) cancelClicked() {
	pD.dialog.Hide()
}

func printerChanged(printer string) {}
