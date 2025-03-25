//go:build windows

package print

/*
import (
	"testing"
)

func TestPrinter_defaultPaperSize(t *testing.T) {
	// Create some custom paper sizes for testing
	customPaperSizes := paperSizes{
		sizes: []PaperSize{
			*newPaperSize("iso_a4_210x297mm", "A4", dmPaperA4, 21000, 29700),
			*newPaperSize("na_letter_8.5x11in", "NA Letter", dmPaperLetter, 8.5*2540, 11*2540),
			*newPaperSize("Custom_size", "Custom", dmPaperNone, 1000, 2000),
			*newPaperSize("custom_150x250mm", "Custom2", dmPaperNone, 15000, 25000),
			*newPaperSize("custom_200x300mm", "Custom3", dmPaperNone, 20000, 30000),
			*newPaperSize("Custom_size_with_.5", "Custom_4", dmPaperNone, 85050, 110070),
		},
	}
	tests := []struct {
		name        string
		printer     Printer
		want        *PaperSize
		shouldError bool
	}{
		{
			name: "Standard A4",
			printer: Printer{
				printerInfo2: PrinterInfo2{
					DevMode: &PrinterDevMode{
						dmPaperSize: dmPaperA4,
					},
				},
			},
			want: &PaperSize{
				psN:     "iso_a4_210x297mm",
				n:       "A4",
				winSize: dmPaperA4,
				w:       21000,
				h:       29700,
			},
		},
		{
			name: "Standard Letter",
			printer: Printer{
				printerInfo2: PrinterInfo2{
					DevMode: &PrinterDevMode{
						dmPaperSize: dmPaperLetter,
					},
				},
			},
			want: &PaperSize{
				psN:     "na_letter_8.5x11in",
				n:       "NA Letter",
				winSize: dmPaperLetter,
				w:       8.5 * 2540,
				h:       11. * 2540,
			},
		},
		{
			name: "Standard None",
			printer: Printer{
				printerInfo2: PrinterInfo2{
					DevMode: &PrinterDevMode{
						dmPaperSize: dmPaperNone,
					},
				},
			},
			want: nil,
		},
		{
			name: "Custom size",
			printer: Printer{
				printerInfo2: PrinterInfo2{
					DevMode: &PrinterDevMode{
						dmPaperSize:   dmPaperNone,
						dmPaperWidth:  100,
						dmPaperLength: 200,
					},
				},
			},
			want: &PaperSize{
				psN:     "Custom_size",
				n:       "Custom",
				w:       1000,
				h:       2000,
				winSize: dmPaperNone,
			},
		},
		{
			name: "Custom size with .5",
			printer: Printer{
				printerInfo2: PrinterInfo2{
					DevMode: &PrinterDevMode{
						dmPaperSize:   dmPaperNone,
						dmPaperWidth:  8505,
						dmPaperLength: 11007,
					},
				},
			},
			want: &PaperSize{
				psN:     "Custom_size_with_.5",
				n:       "Custom_4",
				w:       85050,
				h:       110070,
				winSize: dmPaperNone,
			},
		},
		{
			name: "custom paper size not exist",
			printer: Printer{
				printerInfo2: PrinterInfo2{
					DevMode: &PrinterDevMode{
						dmPaperSize:   dmPaperNone,
						dmPaperWidth:  101,
						dmPaperLength: 202,
					},
				},
			},
			want: nil,
		},
	}
	//Add custom paper sizes to stdPaperSizes
	stdPaperSizes = customPaperSizes

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.printer.defaultPaperSize()
			if tt.want != nil && got == nil {
				t.Errorf("Printer.defaultPaperSize() = %v, want %v", got, tt.want)
			}

			if tt.want == nil && got != nil {
				t.Errorf("Printer.defaultPaperSize() = %v, want %v", got, tt.want)
			}
			if tt.want != nil && got != nil {
				if got.psName() != tt.want.psName() ||
					got.name() != tt.want.name() ||
					got.winSize != tt.want.winSize ||
					got.w != tt.want.w ||
					got.h != tt.want.h {
					t.Errorf("Printer.defaultPaperSize() = %v, want %v", got, tt.want)
				}

			}
		})
	}
}
*/
