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
			paperSize: newPaperSize("iso_a4_210x297mm", "A4", dmPaperA4, 21000, 29700),
			want:      fyne.NewSize(210, 297),
		},
		{
			name:      "Letter",
			paperSize: newPaperSize("na_letter_8.5x11in", "NA Letter", dmPaperLetter, 8.5*2540, 11*2540),
			want:      fyne.NewSize(8.5*2540/100, 11.*2540/100),
		},
		{
			name:      "A0",
			paperSize: newPaperSize("iso_a0_841x1189mm", "A0", dmPaperNone, 84100, 118900),
			want:      fyne.NewSize(841, 1189),
		},
		{
			name:      "Custom",
			paperSize: newPaperSize("custom_100x200mm", "Custom", dmPaperNone, 10000, 20000),
			want:      fyne.NewSize(100, 200),
		},
		{
			name:      "B0",
			paperSize: newPaperSize("iso_b0_1000x1414mm", "B0", dmPaperNone, 100000, 141400),
			want:      fyne.NewSize(1000, 1414),
		},
		{
			name:      "Photo L",
			paperSize: newPaperSize("oe_photo-l_3.5x5in", "Photo L", dmPaperNone, 3.5*2540, 5*2540),
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
			paperSize: newPaperSize("iso_a4_210x297mm", "A4", dmPaperA4, 21000, 29700),
			want:      fyne.NewSize(21000.0/2540, 29700.0/2540),
		},
		{
			name:      "Letter",
			paperSize: newPaperSize("na_letter_8.5x11in", "NA Letter", dmPaperLetter, 8.5*2540, 11*2540),
			want:      fyne.NewSize(8.5, 11),
		},
		{
			name:      "A0",
			paperSize: newPaperSize("iso_a0_841x1189mm", "A0", dmPaperNone, 84100, 118900),
			want:      fyne.NewSize(84100.0/2540, 118900.0/2540),
		},
		{
			name:      "Custom",
			paperSize: newPaperSize("custom_100x200mm", "Custom", dmPaperNone, 10000, 20000),
			want:      fyne.NewSize(10000.0/2540, 20000.0/2540),
		},
		{
			name:      "B0",
			paperSize: newPaperSize("iso_b0_1000x1414mm", "B0", dmPaperNone, 100000, 141400),
			want:      fyne.NewSize(100000.0/2540, 141400.0/2540),
		},
		{
			name:      "Photo L",
			paperSize: newPaperSize("oe_photo-l_3.5x5in", "Photo L", dmPaperNone, 3.5*2540, 5*2540),
			want:      fyne.NewSize(3.5, 5),
		},
		{
			name:      "Square Photo 4x4in",
			paperSize: newPaperSize("oe_square-photo_4x4in", "Square Photo 4x4in", dmPaperNone, 4*2540, 4*2540),
			want:      fyne.NewSize(4, 4),
		},
		{
			name:      "Card 54x86mm",
			paperSize: newPaperSize("om_card_54x86mm", "Card 54x86mm", dmPaperNone, 5400, 8600),
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

func Test_paperNameFromDimensions(t *testing.T) {
	tests := []struct {
		name string
		w    float32
		h    float32
		want string
	}{
		{
			name: "Standard A4",
			w:    21000,
			h:    29700,
			want: "210x297mm",
		},
		{
			name: "Custom size",
			w:    15000,
			h:    20000,
			want: "150x200mm",
		},
		{
			name: "Small size",
			w:    1000,
			h:    2000,
			want: "10x20mm",
		},
		{
			name: "Large size",
			w:    100000,
			h:    141400,
			want: "1000x1414mm",
		},
		{
			name: "Zero size",
			w:    0,
			h:    0,
			want: "0x0mm",
		},
		{
			name: "Decimal size with .5",
			w:    85050,
			h:    110070,
			want: "850.5x1100.7mm",
		},
		{
			name: "Unequal dimensions",
			w:    12345,
			h:    67890,
			want: "123.45x678.9mm",
		},
		{
			name: "Decimal with .0",
			w:    10000,
			h:    20000,
			want: "100x200mm",
		},
		{
			name: "Decimal with .0 and other",
			w:    10000,
			h:    20050,
			want: "100x200.5mm",
		},
		{
			name: "Single decimal",
			w:    1050,
			h:    10000,
			want: "10.5x100mm",
		},
		{
			name: "Different Decimal places",
			w:    12345.678,
			h:    987.65,
			want: "123.46x9.88mm",
		},
		{
			name: "Rounding test",
			w:    1.0,
			h:    2.0,
			want: "0.01x0.02mm",
		},
		{
			name: "Rounding test 2",
			w:    1.5,
			h:    2.5,
			want: "0.02x0.03mm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paperNameFromDimensions(tt.w, tt.h); got != tt.want {
				t.Errorf("paperNameFromDimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createFormat(t *testing.T) {
	tests := []struct {
		name string
		dim  float32
		want string
	}{
		{
			name: "No decimal",
			dim:  100,
			want: "%.0f",
		},
		{
			name: "One decimal place",
			dim:  100.5,
			want: "%.1f",
		},
		{
			name: "Two decimal places",
			dim:  100.55,
			want: "%.2f",
		},
		{
			name: "Four decimal places with .00",
			dim:  100.0053,
			want: "%.2f",
		},
		{
			name: "Four decimal places with .05",
			dim:  100.0553,
			want: "%.2f",
		},
		{
			name: "Whole decimal",
			dim:  100.0,
			want: "%.0f",
		},
		{
			name: "Zero",
			dim:  0.0,
			want: "%.0f",
		},
		{
			name: "Large Single decimal",
			dim:  10000.5,
			want: "%.1f",
		},
		{
			name: "Two decimals with truncation",
			dim:  100.555,
			want: "%.2f",
		},
		{
			name: "One decimal with truncation",
			dim:  100.55,
			want: "%.2f",
		},
		{
			name: "No decimal with truncation",
			dim:  100.001,
			want: "%.0f",
		},
		{
			name: "Negative values",
			dim:  -100,
			want: "%.0f",
		},
		{
			name: "Negative decimal values",
			dim:  -100.5,
			want: "%.1f",
		},
		{
			name: "Negative two decimal places",
			dim:  -100.55,
			want: "%.2f",
		},
		{
			name: "Negative three decimal places",
			dim:  -100.555,
			want: "%.2f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFormat(tt.dim); got != tt.want {
				t.Errorf("createFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
