package print

import (
	"embed"

	"fyne.io/fyne/v2/lang"
)

//go:embed printtranslation
var translations embed.FS

// init loads translations related to fyne printing.
func init() {
	lang.AddTranslationsFS(translations, "printtranslation")
}
