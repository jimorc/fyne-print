package print

import (
	"testing"

	"fyne.io/fyne/v2"
)

func TestPaperSizes_findPaperSizeFromSize(t *testing.T) {
	tests := []struct {
		name      string
		sizes     paperSizes
		inputSize fyne.Size
		want      *PaperSize
	}{
		{
			name:      "A4 found",
			sizes:     stdPaperSizes,
			inputSize: fyne.NewSize(21000, 29700),
			want: &PaperSize{
				psN: "iso_a4_297x420mm",
				n:   "A4",
				w:   21000,
				h:   29700,
			},
		},
		{
			name:      "Letter found",
			sizes:     stdPaperSizes,
			inputSize: fyne.NewSize(8.5*2540, 11*2540),
			want: &PaperSize{
				psN: "na_letter_8.5x11in",
				n:   "NA Letter",
				w:   8.5 * 2540,
				h:   11. * 2540,
			},
		},
		{
			name: "Not found",
			sizes: paperSizes{
				sizes: []PaperSize{
					newPaperSize("custom_100x200mm", "Custom", 10000, 20000),
				},
			},
			inputSize: fyne.NewSize(15000, 25000),
			want:      nil,
		},
		{
			name: "Empty paperSizes",
			sizes: paperSizes{
				sizes: []PaperSize{},
			},
			inputSize: fyne.NewSize(10000, 20000),
			want:      nil,
		},
		{
			name:      "B0 Found",
			sizes:     stdPaperSizes,
			inputSize: fyne.NewSize(100000, 141400),
			want: &PaperSize{
				psN: "iso_b0_1000x1414mm",
				n:   "B0",
				w:   100000,
				h:   141400,
			},
		},
		{
			name:      "A0 Found",
			sizes:     stdPaperSizes,
			inputSize: fyne.NewSize(84100, 118900),
			want: &PaperSize{
				psN: "iso_a0_841x1189mm",
				n:   "A0",
				w:   84100,
				h:   118900,
			},
		},
		{
			name: "custom found",
			sizes: paperSizes{
				sizes: []PaperSize{
					newPaperSize("custom_100x200mm", "Custom", 10000, 20000),
				},
			},
			inputSize: fyne.NewSize(10000, 20000),
			want: &PaperSize{
				psN: "custom_100x200mm",
				n:   "Custom",
				w:   10000,
				h:   20000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sizes.findPaperSizeFromSize(tt.inputSize)

			if tt.want == nil && got != nil {
				t.Errorf("findPaperSizeFromSize() = %v, want %v", got, tt.want)
			}
			if tt.want != nil && got == nil {
				t.Errorf("findPaperSizeFromSize() = %v, want %v", got, tt.want)
			}
			if tt.want != nil && got != nil {
				if got.psName() != tt.want.psName() || got.name() != tt.want.name() || got.width() != tt.want.width() || got.height() != tt.want.height() {
					t.Errorf("findPaperSizeFromSize() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
