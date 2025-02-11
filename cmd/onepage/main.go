package main

import (
	"jimorc/github.com/fyne-print/print"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("One Page Print")
	pSetup := print.NewPageSetupDialog(w)
	pSetup.Show()
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
