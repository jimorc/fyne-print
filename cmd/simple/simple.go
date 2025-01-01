package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	pdialog "github.com/jimorc/fyne-print/dialog"
)

func main() {
	a := app.New()
	w := a.NewWindow("Simple Print")
	l1 := widget.NewLabel("Label 1")
	l1.Alignment = fyne.TextAlignCenter
	l2 := widget.NewLabel("Label 2")
	box := container.NewVBox(l1, l2)

	c1 := widget.NewCard("", "", box)
	v3 := container.NewVBox(c1)
	w.SetContent(v3)
	w.Resize(fyne.NewSize(800, 600))
	pDialog := pdialog.NewPrintDialog(w)
	pDialog.Show()
	w.ShowAndRun()
}
