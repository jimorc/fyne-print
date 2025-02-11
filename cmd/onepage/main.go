package main

import (
	"jimorc/github.com/fyne-print/print"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("One Page Print")
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Page Setup", func() {
			pSetup := print.NewPageSetupDialog(w)
			pSetup.Show()
		}),
		fyne.NewMenuItem("Quit", func() {
			w.Close()
		}))

	mainMenu := fyne.NewMainMenu(fileMenu)
	w.SetMainMenu(mainMenu)
	// we need some content in the main window or we get some
	// weird results (menu displays or doesn't)
	w.SetContent(widget.NewLabel("Put something in the window"))
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
