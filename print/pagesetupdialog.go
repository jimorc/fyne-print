package print

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	xlayout "fyne.io/x/fyne/layout"
)

// PageSetupInfo contains information used to initialize the widgets in the PageSetupDialog and
// to return data from it.
type PageSetupInfo struct {
	printer     *Printer
	paperSize   string
	orientation string
}

// NewPageSetupInfo creates a PageSetupInfo object.
//
// Params:
//
//	printer is a pointer to a Printer object. This is the printer that prints will be sent to.
//	paperSize is the size of the paper that will be printed. The default paper size is the printer's
//
// paper size.
//
//	orientation is the paper orientation (portrait or landscape). The default is portrait.
func NewPageSetupInfo(printer *Printer, paperSize string, orientation string) *PageSetupInfo {
	return &PageSetupInfo{
		printer:     printer,
		paperSize:   paperSize,
		orientation: orientation,
	}
}

// PageSetupDialog is a ConfirmDialog dialog with widgets that must be saved.
type PageSetupDialog struct {
	*dialog.ConfirmDialog
	pageSetupInfo         *PageSetupInfo
	parent                fyne.Window
	printerSelect         *widget.Select
	paperSizeSelect       *widget.Select
	orientationRadioGroup *widget.RadioGroup
}

// NewPageSetupDialog creates a PageSetupDialog which is a ConfirmDialog.
//
// Params:
//
//	parent is the parent window for the dialog.
func NewPageSetupDialog(parent fyne.Window, psInfo *PageSetupInfo) *dialog.ConfirmDialog {
	psd := &PageSetupDialog{}
	if psInfo == nil {
		psInfo = &PageSetupInfo{}
	}
	psd.pageSetupInfo = psInfo
	psd.parent = parent
	printerContainer := psd.createPrinterContainer()
	psd.ConfirmDialog = dialog.NewCustomConfirm("PageSetup", "OK",
		"Cancel", printerContainer, func(bool) {}, parent)
	psd.Resize(fyne.NewSize(500, 300))
	return psd.ConfirmDialog

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

	// set selected
	if len(prNames) > 0 {
		if psd.pageSetupInfo.printer != nil {
			psd.printerSelect.SetSelected(psd.pageSetupInfo.printer.Name())
		} else if len(prNames) == 1 {
			psd.printerSelect.SetSelected(prNames[0])
		} else {
			defPr := printers.DefaultPrinter()
			if defPr != nil {
				psd.printerSelect.SetSelected(defPr.Name())
			}
		}
	}
}
