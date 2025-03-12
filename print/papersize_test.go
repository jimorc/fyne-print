package print

import (
	"testing"

	"fyne.io/fyne/v2"
)

func TestPaperSize_SizeInMM(t *testing.T) {
	tests := []struct {
		name      string
		paperSize PaperSize
		want      fyne.Size
	}{
		{
			name:      "A4",
			paperSize: newPaperSize("iso_a4_210x297mm", "A4", 21000, 29700),
			want:      fyne.NewSize(210, 297),
		},
		{
			name:      "Letter",
			paperSize: newPaperSize("na_letter_8.5x11in", "NA Letter", 8.5*2540, 11*2540),
			want:      fyne.NewSize(8.5*2540/100, 11.*2540/100),
		},
		{
			name:      "A0",
			paperSize: newPaperSize("iso_a0_841x1189mm", "A0", 84100, 118900),
			want:      fyne.NewSize(841, 1189),
		},
		{
			name:      "Custom",
			paperSize: newPaperSize("custom_100x200mm", "Custom", 10000, 20000),
			want:      fyne.NewSize(100, 200),
		},
		{
			name:      "B0",
			paperSize: newPaperSize("iso_b0_1000x1414mm", "B0", 100000, 141400),
			want:      fyne.NewSize(1000, 1414),
		},
		{
			name:      "Photo L",
			paperSize: newPaperSize("oe_photo-l_3.5x5in", "Photo L", 3.5*2540, 5*2540),
			want:      fyne.NewSize(3.5*2540/100, 5*2540/100),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.paperSize.SizeInMM(); got != tt.want {
				t.Errorf("PaperSize.SizeInMM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaperSize_SizeInInches(t *testing.T) {
	tests := []struct {
		name      string
		paperSize PaperSize
		want      fyne.Size
	}{
		{
			name:      "A4",
			paperSize: newPaperSize("iso_a4_210x297mm", "A4", 21000, 29700),
			want:      fyne.NewSize(21000.0/2540, 29700.0/2540),
		},
		{
			name:      "Letter",
			paperSize: newPaperSize("na_letter_8.5x11in", "NA Letter", 8.5*2540, 11*2540),
			want:      fyne.NewSize(8.5, 11),
		},
		{
			name:      "A0",
			paperSize: newPaperSize("iso_a0_841x1189mm", "A0", 84100, 118900),
			want:      fyne.NewSize(84100.0/2540, 118900.0/2540),
		},
		{
			name:      "Custom",
			paperSize: newPaperSize("custom_100x200mm", "Custom", 10000, 20000),
			want:      fyne.NewSize(10000.0/2540, 20000.0/2540),
		},
		{
			name:      "B0",
			paperSize: newPaperSize("iso_b0_1000x1414mm", "B0", 100000, 141400),
			want:      fyne.NewSize(100000.0/2540, 141400.0/2540),
		},
		{
			name:      "Photo L",
			paperSize: newPaperSize("oe_photo-l_3.5x5in", "Photo L", 3.5*2540, 5*2540),
			want:      fyne.NewSize(3.5, 5),
		},
		{
			name:      "Square Photo 4x4in",
			paperSize: newPaperSize("oe_square-photo_4x4in", "Square Photo 4x4in", 4*2540, 4*2540),
			want:      fyne.NewSize(4, 4),
		},
		{
			name:      "Card 54x86mm",
			paperSize: newPaperSize("om_card_54x86mm", "Card 54x86mm", 5400, 8600),
			want:      fyne.NewSize(5400.0/2540, 8600.0/2540),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.paperSize.SizeInInches(); got != tt.want {
				t.Errorf("PaperSize.SizeinInches() = %v, want %v", got, tt.want)
			}
		})
	}
}
