package print

import (
	"testing"

	"fyne.io/fyne/v2"
)

func TestPaperSizes_findPaperSizeFromWindowsPaperSize(t *testing.T) {
	tests := []struct {
		name      string
		sizes     paperSizes
		inputSize fyne.Size
		want      *PaperSize
	}{
		{
			name:      "A4 found",
			sizes:     stdPaperSizes,
			inputSize: fyne.NewSize(2100, 2970),
			want: &PaperSize{
				psN: "iso_a4_210x297mm",
				n:   "A4",
				w:   21000,
				h:   29700,
			},
		},
		{
			name:      "Letter found",
			sizes:     stdPaperSizes,
			inputSize: fyne.NewSize(8.5*254, 11*254),
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
			inputSize: fyne.NewSize(1500, 2500),
			want:      nil,
		},
		{
			name: "Empty paperSizes",
			sizes: paperSizes{
				sizes: []PaperSize{},
			},
			inputSize: fyne.NewSize(1000, 2000),
			want:      nil,
		},
		{
			name:      "B0 Found",
			sizes:     stdPaperSizes,
			inputSize: fyne.NewSize(10000, 14140),
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
			inputSize: fyne.NewSize(8410, 11890),
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
			inputSize: fyne.NewSize(1000, 2000),
			want: &PaperSize{
				psN: "custom_100x200mm",
				n:   "Custom",
				w:   10000,
				h:   20000,
			},
		},
		{
			name:      "JIS B5 found",
			sizes:     stdPaperSizes,
			inputSize: fyne.NewSize(1820, 2570),
			want: &PaperSize{
				psN: "jis_b5_182x257mm",
				n:   "JIS B5",
				w:   18200,
				h:   25700,
			},
		},
		{
			name:      "Photo L Found",
			sizes:     stdPaperSizes,
			inputSize: fyne.NewSize(3.5*254, 5*254),
			want: &PaperSize{
				psN: "oe_photo-l_3.5x5in",
				n:   "Photo L",
				w:   3.5 * 2540,
				h:   5 * 2540,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sizes.findPaperSizeFromWindowsPaperSize(tt.inputSize)

			if tt.want == nil && got != nil {
				t.Errorf("findPaperSizeFromWindowsPaperSize() = %v, want %v", got, tt.want)
			}
			if tt.want != nil && got == nil {
				t.Errorf("findPaperSizeFromWindowsPaperSize() = %v, want %v", got, tt.want)
			}
			if tt.want != nil && got != nil {
				if got.psName() != tt.want.psName() || got.name() != tt.want.name() || got.width() != tt.want.width() || got.height() != tt.want.height() {
					t.Errorf("findPaperSizeFromWindowsPaperSize() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
