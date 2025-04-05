//go:build windows

package print

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMediaSizeString_NoData(t *testing.T) {
	ms := newMediaSize("Letter", 1)
	assert.Equal(t, "MediaSize:\n    Media Name: Letter\n    PaperSize: 1, DMPAPER_LETTER\n    No data\n", ms.String())
}

func TestMediaSizeString_SetData(t *testing.T) {
	ms := newMediaSize("Letter", 1)
	ms.setData(1000.0, 1200.0, 900.0, 1100.0, 45.0, 65.0)
	assert.Equal(t, "MediaSize:\n    Media Name: Letter\n    PaperSize: 1, DMPAPER_LETTER\n"+
		"    Width: 1000.00\n    Height: 1200.00\nMargins:    top: 65.00    bottom: 35.00"+
		"    left: 45.00    right: 55.00\n", ms.String())
}

func TestMediaSizeString_WithData(t *testing.T) {
	ms := newMediaSizeWithData("Letter", 1, 1000.0, 1200.0, 900.0, 1100.0, 35.0, 55.0)
	assert.Equal(t, "MediaSize:\n    Media Name: Letter\n    PaperSize: 1, DMPAPER_LETTER\n"+
		"    Width: 1000.00\n    Height: 1200.00\nMargins:    top: 55.00    bottom: 45.00"+
		"    left: 35.00    right: 65.00\n", ms.String())
}
