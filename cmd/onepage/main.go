package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/software"
	"fyne.io/fyne/v2/widget"

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

	circle := canvas.NewCircle(color.Gray16{Y: 10})
	circle.FillColor = color.Gray{Y: 0xC0}
	circle.StrokeColor = color.Black
	circle.StrokeWidth = 10
	circle.Resize(fyne.NewSize(50, 50))
	circle.Move(fyne.NewPos(100, 100))

	line := canvas.NewLine(color.Black)
	line.StrokeWidth = 10
	line.Position1 = fyne.NewPos(0, 0)
	line.Position2 = fyne.NewPos(120, 60)

	text := canvas.NewText("Test string", color.RGBA{255, 0, 0, 255})
	text.Move(fyne.NewPos(200, 200))
	button := widget.NewButton("Press Here", func() { fmt.Println("Button") })
	button.Move(fyne.NewPos(200, 250))

	c3 := container.New(print.NewPrintPageLayout(), circle, line, text, button)

	i := software.Render(c3, button.Theme())
	img := canvas.NewImageFromImage(i)
	w.SetContent(img)
	w.Resize(fyne.NewSize(600, 850))
	w.ShowAndRun()

}
