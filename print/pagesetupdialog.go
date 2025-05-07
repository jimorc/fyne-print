package print

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// PageSetupInfo contains information used to initialize the widgets in the PageSetupDialog and
// to return data from it.
type PageSetupInfo struct {
	//	printer     *Printer
	//	mediaSize   *MediaSize
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
/*func NewPageSetupInfo(printer *Printer, //paperSize *PaperSize,
	orientation string) *PageSetupInfo {
	return &PageSetupInfo{
		printer: printer,
		//		paperSize:   paperSize,
		orientation: orientation,
	}
}*/

// PageSetupDialog is a ConfirmDialog dialog with widgets that must be saved.
type PageSetupDialog struct {
	*dialog.CustomDialog
	pageSetupInfo         *PageSetupInfo
	parent                fyne.Window
	printers              *Printers
	printerSelect         *widget.Select
	location              *widget.Label
	comment               *widget.Label
	paperSizeSelect       *widget.Select
	orientationRadioGroup *widget.RadioGroup
}

// NewPageSetupDialog creates a PageSetupDialog which is a ConfirmDialog.
//
// Params:
//
//		parent is the parent window for the dialog.
//	 psInfo is the PageSetupInfo struct containing the data to initialize the dialog's
//
// widgets to. If nil is passed, then an initialized PageSetupInfo struct is created.
// The contents of the dialog is saved to the struct when the dialog's submit button is clicked.
func NewPageSetupDialog(parent fyne.Window, psInfo *PageSetupInfo) *PageSetupDialog {
	psd := &PageSetupDialog{}
	if psInfo == nil {
		psInfo = &PageSetupInfo{}
	}
	psd.pageSetupInfo = psInfo
	psd.parent = parent
	psd.CustomDialog = psd.createDialog()
	psd.Resize(fyne.NewSize(500, 300))
	return psd

}

// createFormDialog creates the dialog containing PageSetup properties.
func (psd *PageSetupDialog) createDialog() *dialog.CustomDialog {
	psd.printerSelect = widget.NewSelect([]string{}, psd.printerSelected)
	psd.printerSelect.Alignment = fyne.TextAlignTrailing
	printerFI := widget.NewFormItem("Format For", psd.printerSelect)
	psd.location = widget.NewLabel("")
	locFI := widget.NewFormItem("Location", psd.location)
	psd.comment = widget.NewLabel("")
	commentFI := widget.NewFormItem("Comment", psd.comment)
	psd.paperSizeSelect = widget.NewSelect([]string{}, nil)
	paperSizeFI := widget.NewFormItem("Paper Size", psd.paperSizeSelect)
	psd.paperSizeSelect.Alignment = fyne.TextAlignTrailing
	psd.orientationRadioGroup = widget.NewRadioGroup([]string{"Portrait", "Landscape"}, nil)
	psd.orientationRadioGroup.Horizontal = true
	orientationFI := widget.NewFormItem("Orientation", psd.orientationRadioGroup)

	form := widget.NewForm(printerFI, locFI, commentFI, paperSizeFI, orientationFI)
	form.OnCancel = psd.onCancel
	form.OnSubmit = psd.onSubmit
	form.SubmitText = "Accept"
	d := dialog.NewCustomWithoutButtons("Page Setup", form, psd.parent)
	return d
}

// populatePrinterSelect populates the printerSelect with the names of all available printers
// and sets the selected item to the previously set printer, the only printer, or the user's
// default printer if set.
func (psd *PageSetupDialog) populatePrinterSelect() error {
	ps := NewPrinters()
	psd.printers = ps

	prNames := psd.printers.Names()
	if len(prNames) == 0 {
		err := errors.New("no printers were found")
		fyne.LogError("Check printer configuration", err)
		err1 := errors.New(err.Error() +
			"\nCannot continue page setup.")
		return err1
	}
	psd.printerSelect.Options = prNames

	// set selected
	/*
		if len(prNames) > 0 {
			if psd.pageSetupInfo.printer != nil {
				psd.printerSelect.SetSelected(psd.pageSetupInfo.printer.Name())
			} else if len(prNames) == 1 {
				psd.printerSelect.SetSelected(prNames[0])
			} else {
				defPr := psd.printers.DefaultPrinter()
				if defPr != nil {
					psd.printerSelect.SetSelected(defPr.Name())
				}
			}
		}*/
	return nil
}

func (psd *PageSetupDialog) printerSelected(name string) {
	/*			if psd.printers == nil {
					return
				}
				pr := psd.printers.getPrinterByName(name)
				loc := ""

				if pr != nil && pr.Name() == name {
					loc = pr.Location()
				}

				psd.location.Text = loc
				psd.location.Refresh()
				psd.paperSizeSelect.Options = pr.MediaNames()
				// set selected
				if psd.pageSetupInfo.mediaSize != nil {
					psd.paperSizeSelect.SetSelected(psd.pageSetupInfo.mediaSize.LocalName())
				} else if len(psd.paperSizeSelect.Options) == 1 {
					psd.paperSizeSelect.SetSelectedIndex(0)
						} else {
						defPS := pr.defaultPaperSize()
						if defPS != nil {
							n := defPS.name()
							psd.paperSizeSelect.SetSelected(n)
						}
					}
				}*/
}

// onCancel handles cancel button clicks. It just hides the dialog.
func (psd *PageSetupDialog) onCancel() {
	psd.Hide()
}

// onSubmit handles the submit button clicks. It stores values in the PageSetup
// struct and hides the dialog.
func (psd *PageSetupDialog) onSubmit() {
	psd.Hide()
}
