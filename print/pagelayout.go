package print

import "fyne.io/fyne/v2"

// Declare conformity with Layout interface
var _ fyne.Layout = (*PrintPageLayout)(nil)

type PrintPageLayout struct{}

func NewPrintPageLayout() fyne.Layout {
	return PrintPageLayout{}
}

func (p PrintPageLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(0, size.Height-p.MinSize(objects).Height)
	for _, child := range objects {
		s := child.Size()
		child.Resize(s)
		child.Move(child.Position())

		pos = pos.Add(fyne.NewPos(s.Width, s.Height))
	}

}

func (p PrintPageLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(1000, 1000)
}
