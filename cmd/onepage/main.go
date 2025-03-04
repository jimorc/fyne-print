package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/jimorc/fyne-print/print"
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

}
