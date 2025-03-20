package print

import "fyne.io/fyne/v2"

// dmPaperSize represents the standard Windows paper sizes as a number.
type dmPaperSize int16

// Standard Windows paper sizes. These do not necessarily match the IPP paper sizes
// standardized by IANA.
// While there are values defined for the paper sizes containing Rotated or Transverse,
// none of these values are used in stdPaperSizes because the Rotated and Transverse
// paper sizes are the same sizes as the paper sizes without Rotated or Transverse in
// their names. The only difference is that the Rotated and Transverse ones are rotated
// 90 degrees from the non-Rotated or non-Transverse sizes. This rotation is handled
// in the code by the orientation flag.
const (
	dmPaperNone                         dmPaperSize = 0
	dmPaperFirst                        dmPaperSize = dmPaperLetter
	dmPaperLetter                       dmPaperSize = 1
	dmPaperLetterSmall                  dmPaperSize = 2
	dmPaperTabloid                      dmPaperSize = 3
	dmPaperLedger                       dmPaperSize = 4
	dmPaperLegal                        dmPaperSize = 5
	dmPaperStatement                    dmPaperSize = 6
	dmPaperExecutive                    dmPaperSize = 7
	dmPaperA3                           dmPaperSize = 8
	dmPaperA4                           dmPaperSize = 9
	dmPaperA4Small                      dmPaperSize = 10
	dmPaperA5                           dmPaperSize = 11
	dmPaperB4                           dmPaperSize = 12
	dmPaperB5                           dmPaperSize = 13
	dmPaperFolio                        dmPaperSize = 14
	dmPaperQuarto                       dmPaperSize = 15
	dmPaper10x14                        dmPaperSize = 16
	dmPaper11x17                        dmPaperSize = 17
	dmPaperNote                         dmPaperSize = 18
	dmPaperEnv9                         dmPaperSize = 19
	dmPaperEnv10                        dmPaperSize = 20
	dmPaperEnv11                        dmPaperSize = 21
	dmPaperEnv12                        dmPaperSize = 22
	dmPaperEnv14                        dmPaperSize = 23
	dmPaperCSheet                       dmPaperSize = 24
	dmPaperDSheet                       dmPaperSize = 25
	dmPaperESheet                       dmPaperSize = 26
	dmPaperEnvDL                        dmPaperSize = 27
	dmPaperEnvC5                        dmPaperSize = 28
	dmPaperEnvC3                        dmPaperSize = 29
	dmPaperEnvC4                        dmPaperSize = 30
	dmPaperEnvC6                        dmPaperSize = 31
	dmPaperEnvC65                       dmPaperSize = 32
	dmPaperEnvB4                        dmPaperSize = 33
	dmPaperEnvB5                        dmPaperSize = 34
	dmPaperEnvB6                        dmPaperSize = 35
	dmPaperEnvItaly                     dmPaperSize = 36
	dmPaperEnvMonarch                   dmPaperSize = 37
	dmPaperEnvPersonal                  dmPaperSize = 38
	dmPaperFanfoldUS                    dmPaperSize = 39
	dmPaperFanfoldGerman                dmPaperSize = 40
	dmPaperFanfoldLegal                 dmPaperSize = 41
	dmPaperISOB4                        dmPaperSize = 42
	dmPaperJapanesePostcard             dmPaperSize = 43
	dmPaper9x11                         dmPaperSize = 44
	dmPaper10x11                        dmPaperSize = 45
	dmPaper15x11                        dmPaperSize = 46
	dmPaperEnvelopeInvite               dmPaperSize = 47
	dmPaperReserved48                   dmPaperSize = 48
	dmPaperReserved49                   dmPaperSize = 49
	dmPaperLetterExtra                  dmPaperSize = 50
	dmPaperLegalExtra                   dmPaperSize = 51
	dmPaperTabloidExtra                 dmPaperSize = 52
	dmPaperA4Extra                      dmPaperSize = 53
	dmPaperLetterTransverse             dmPaperSize = 54
	dmPaperA4Transverse                 dmPaperSize = 55
	dmPaperLetterExtraTransverse        dmPaperSize = 56
	dmPaperAplus                        dmPaperSize = 57
	dmPaperBplus                        dmPaperSize = 58
	dmPaperLetterPlus                   dmPaperSize = 59
	dmPaperA4Plus                       dmPaperSize = 60
	dmPaperA5Transverse                 dmPaperSize = 61
	dmPaperB5Transverse                 dmPaperSize = 62
	dmPaperA3Extra                      dmPaperSize = 63
	dmPaperA5Extra                      dmPaperSize = 64
	dmPaperB5Extra                      dmPaperSize = 65
	dmPaperA2                           dmPaperSize = 66
	dmPaperA3Transverse                 dmPaperSize = 67
	dmPaperA3ExtraTransverse            dmPaperSize = 68
	dmPaperDblJapanesePostcard          dmPaperSize = 69
	dmPaperA6                           dmPaperSize = 70
	dmPaperJapaneseEnvelopeKaku2        dmPaperSize = 71
	dmPaperJapaneseEnvelopeKaku3        dmPaperSize = 72
	dmPaperJapaneseEnvelopeChou3        dmPaperSize = 73
	dmPaperJapaneseEnvelopeChou4        dmPaperSize = 74
	dmPaperLetterRotated                dmPaperSize = 75
	dmPaperA3Rotated                    dmPaperSize = 76
	dmPaperA4Rotated                    dmPaperSize = 77
	dmPaperA5Rotated                    dmPaperSize = 78
	dmPaperB4JISRotated                 dmPaperSize = 79
	dmPaperB5JISRotated                 dmPaperSize = 80
	dmPaperJapanesePostcardRotated      dmPaperSize = 81
	dmPaperDblJapanesePostcardRotated   dmPaperSize = 82
	dmPaperA6Rotated                    dmPaperSize = 83
	dmPaperJapaneseEnvelopeKaku2Rotated dmPaperSize = 84
	dmPaperJapaneseEnvelopeKaku3Rotated dmPaperSize = 85
	dmPaperJapaneseEnvelopeChou3Rotated dmPaperSize = 86
	dmPaperJapaneseEnvelopeChou4Rotated dmPaperSize = 87
	dmPaperB6JIS                        dmPaperSize = 88
	dmPaperB6JISRotated                 dmPaperSize = 89
	dmPaper12x11                        dmPaperSize = 90
	dmPaperJapaneseEnvelopeYou4         dmPaperSize = 91
	dmPaperJapaneseEnvelopeYou4Rotated  dmPaperSize = 92
	dmPaperP16K                         dmPaperSize = 93
	dmPaperP32K                         dmPaperSize = 94
	dmPaperP32KBIG                      dmPaperSize = 95
	dmPaperPENV1                        dmPaperSize = 96
	dmPaperPENV2                        dmPaperSize = 97
	dmPaperPENV3                        dmPaperSize = 98
	dmPaperPENV4                        dmPaperSize = 99
	dmPaperPENV5                        dmPaperSize = 100
	dmPaperPENV6                        dmPaperSize = 101
	dmPaperPENV7                        dmPaperSize = 102
	dmPaperPENV8                        dmPaperSize = 103
	dmPaperPENV9                        dmPaperSize = 104
	dmPaperPENV10                       dmPaperSize = 105
	dmPaperP16KRotated                  dmPaperSize = 106
	dmPaperP32KRotated                  dmPaperSize = 107
	dmPaperP32KBIGRotated               dmPaperSize = 108
	dmPaperPENV1Rotated                 dmPaperSize = 109
	dmPaperPENV2Rotated                 dmPaperSize = 110
	dmPaperPENV3Rotated                 dmPaperSize = 111
	dmPaperPENV4Rotated                 dmPaperSize = 112
	dmPaperPENV5Rotated                 dmPaperSize = 113
	dmPaperPENV6Rotated                 dmPaperSize = 114
	dmPaperPENV7Rotated                 dmPaperSize = 115
	dmPaperPENV8Rotated                 dmPaperSize = 116
	dmPaperPENV9Rotated                 dmPaperSize = 117
	dmPaperPENV10Rotated                dmPaperSize = 118
	dmPaperLast                         dmPaperSize = dmPaperPENV10Rotated
)

// paperSizes contains a slice of PaperSize objects.
type paperSizes struct {
	sizes []PaperSize
}

// stdPaperSizes contains all of the IPP media sizes registered with IANA
// as of 2024-10-09.
// newPaperSize arguments:
//
//	1: the IPP media size name.
//	2: the English name for the media size. The name is translated in newPaperSize
//
// to the system language if a translation exists.
//
//		3: width of the paper in 100ths of a mm for the paper in portrait mode.
//	 4: height of the paper in 100ths of a mm for the paper in portrait mode. That is,
//
// the height is always equal to or greater than the width.
var stdPaperSizes paperSizes

func init() {
	stdPaperSizes.sizes = []PaperSize{
		*newPaperSize("asme_f_28x40in", "ASME_F", dmPaperNone, 28*2540, 40*2540),
		*newPaperSize("iso_2a0_1189x1682mm", "2xA0", dmPaperNone, 118900, 168200),
		*newPaperSize("iso_a0_841x1189mm", "A0", dmPaperNone, 84100, 118900),
		*newPaperSize("iso_a0x3_1189x2523mm", "A0x3", dmPaperNone, 118900, 252300),
		*newPaperSize("iso_a0x4_1189x3003mm", "A0x4", dmPaperNone, 118900, 300300),
		*newPaperSize("iso_a1_594x841mm", "A1", dmPaperNone, 59400, 84100),
		*newPaperSize("iso_a1x3_841x1783mm", "A1x3", dmPaperNone, 84100, 178300),
		*newPaperSize("iso_a1x4_841x2378mm", "A1x4", dmPaperNone, 84100, 237800),
		*newPaperSize("iso_a2_420x594mm", "A2", dmPaperA2, 42000, 59400),
		*newPaperSize("iso_a2x3_420x1783mm", "A2x3", dmPaperNone, 42000, 178300),
		*newPaperSize("iso_a2x4_420x2378mm", "A2x4", dmPaperNone, 42000, 237800),
		*newPaperSize("iso_a2x5_420x2970mm", "A2x5", dmPaperNone, 42000, 297000),
		*newPaperSize("iso_a3-extra_322x445mm", "A3 Extra", dmPaperA3Extra, 32200, 44500),
		*newPaperSize("iso_a3_297x420mm", "A3", dmPaperA3, 29700, 42000),
		*newPaperSize("iso_a3x3_420x891mm", "A3x3", dmPaperNone, 42000, 89100),
		*newPaperSize("iso_a3x4_420x1189mm", "A3x4", dmPaperNone, 42000, 118900),
		*newPaperSize("iso_a3x5_420x1486mm", "A3x5", dmPaperNone, 42000, 148600),
		*newPaperSize("iso_a3x6_420x1783mm", "A3x6", dmPaperNone, 42000, 178300),
		*newPaperSize("iso_a3x7_420x2080mm", "A3x7", dmPaperNone, 42000, 208000),
		*newPaperSize("iso_a4-extra_235.5x322.3mm", "A4 Extra", dmPaperA4Extra, 23550, 32230),
		*newPaperSize("iso_a4-tab_225x297mm", "A4 Tab", dmPaperNone, 22500, 29700),
		*newPaperSize("iso_a4_210x297mm", "A4", dmPaperA4, 21000, 29700),
		*newPaperSize("iso_a4x3_297x630mm", "A4x3", dmPaperNone, 29700, 63000),
		*newPaperSize("iso_a4x4_297x794mm", "A4x4", dmPaperNone, 29700, 79400),
		*newPaperSize("iso_a4x5_297x958mm", "A4x5", dmPaperNone, 29700, 95800),
		*newPaperSize("iso_a4x6_297x1122mm", "A4x6", dmPaperNone, 29700, 112200),
		*newPaperSize("iso_a4x7_297x1286mm", "A4x7", dmPaperNone, 29700, 128600),
		*newPaperSize("iso_a4x8_297x1450mm", "A4x8", dmPaperNone, 29700, 145000),
		*newPaperSize("iso_a4x9_297x1892mm", "A4x9", dmPaperNone, 29700, 189200),
		*newPaperSize("iso_a5-extra_174x235mm", "A5 Extra", dmPaperA5Extra, 17400, 23500),
		*newPaperSize("iso_a5_148x210mm", "A5", dmPaperA5, 14800, 21000),
		*newPaperSize("iso_a6_105x148mm", "A6", dmPaperA6, 10500, 14800),
		*newPaperSize("iso_a7_74x105mm", "A7", dmPaperNone, 7400, 10500),
		*newPaperSize("iso_a8_52x74mm", "A8", dmPaperNone, 5200, 7400),
		*newPaperSize("iso_a9_37x52mm", "A9", dmPaperNone, 3700, 5200),
		*newPaperSize("iso_a10_26x37mm", "A10", dmPaperNone, 2600, 3700),
		*newPaperSize("iso_b0_1000x1414mm", "B0", dmPaperNone, 100000, 141400),
		*newPaperSize("iso_b1_707x1000mm", "B1", dmPaperNone, 70700, 100000),
		*newPaperSize("iso_b2_500x707mm", "B2", dmPaperNone, 50000, 70700),
		*newPaperSize("iso_b3_353x500mm", "B3", dmPaperNone, 35300, 50000),
		*newPaperSize("iso_b4_250x353mm", "B4 Envelope", dmPaperB4, 25000, 35300),
		*newPaperSize("iso_b5-extra_201x276mm", "B5 Extra", dmPaperB5Extra, 20100, 27600),
		*newPaperSize("iso_b5_176x250mm", "B5 Envelope", dmPaperB5, 17600, 25000),
		*newPaperSize("iso_b6_125x176mm", "B6 Envelope", dmPaperNone, 12500, 17600),
		*newPaperSize("iso_b6c4_125x324mm", "B6C4", dmPaperNone, 12500, 32400),
		*newPaperSize("iso_b7_88x125mm", "B7", dmPaperNone, 8800, 12500),
		*newPaperSize("iso_b8_62x88mm", "B8", dmPaperNone, 6200, 8800),
		*newPaperSize("iso_b9_44x62mm", "B9", dmPaperNone, 4400, 6200),
		*newPaperSize("iso_b10_31x44mm", "B10", dmPaperNone, 3100, 4400),
		*newPaperSize("iso_c0_917x1297mm", "C0", dmPaperNone, 91700, 129700),
		*newPaperSize("iso_c1_648x917mm", "C1", dmPaperNone, 64800, 91700),
		*newPaperSize("iso_c2_458x648mm", "C2", dmPaperNone, 45800, 64800),
		*newPaperSize("iso_c3_324x458mm", "C3 Envelope", dmPaperEnvC3, 32400, 45800),
		*newPaperSize("iso_c4_229x324mm", "C4 Envelope", dmPaperEnvC4, 22900, 32400),
		*newPaperSize("iso_c5_162x229mm", "C5 Envelope", dmPaperEnvC5, 16200, 22900),
		*newPaperSize("iso_c6_114x162mm", "C6 Envelope", dmPaperEnvC6, 11400, 16200),
		*newPaperSize("iso_c6c5_114x229mm", "C6C5 Envelope", dmPaperEnvC65, 11400, 22900),
		*newPaperSize("iso_c7_77x114mm", "C7 Envelope", dmPaperNone, 7700, 11400),
		*newPaperSize("iso_c7c6_81x162mm", "C7C6 Envelope", dmPaperNone, 8100, 16200),
		*newPaperSize("iso_c8_56x77mm", "C8 Envelope", dmPaperNone, 5600, 7700),
		*newPaperSize("iso_c9_38x56mm", "C9 Envelope", dmPaperNone, 3800, 5600),
		*newPaperSize("iso_c10_29x38mm", "C10 Envelope", dmPaperNone, 2900, 3800),
		*newPaperSize("iso_dl_110x220mm", "DL Envelope", dmPaperEnvDL, 11000, 22000),
		*newPaperSize("iso_id-1_53.98x85.6mm", "ID-1", dmPaperNone, 5398, 8560),
		*newPaperSize("iso_ra0_860x1220mm", "RA0", dmPaperNone, 86000, 122000),
		*newPaperSize("iso_ra1_610x860mm", "RA1", dmPaperNone, 61000, 86000),
		*newPaperSize("iso_ra2_420x610mm", "RA2", dmPaperNone, 42000, 61000),
		*newPaperSize("iso_ra3_305x420mm", "RA3", dmPaperNone, 30500, 42000),
		*newPaperSize("iso_ra4_210x305mm", "RA4", dmPaperNone, 21000, 30500),
		*newPaperSize("iso_sra0_900x1280mm", "RA5", dmPaperNone, 14800, 21000),
		*newPaperSize("iso_sra1_640x900mm", "SRA1", dmPaperNone, 64000, 90000),
		*newPaperSize("iso_sra2_450x640mm", "SRA2", dmPaperNone, 45000, 64000),
		*newPaperSize("iso_sra3_320x450mm", "SRA3", dmPaperNone, 32000, 45000),
		*newPaperSize("iso_sra4_225x320mm", "SRA4", dmPaperNone, 22500, 32000),
		*newPaperSize("jis_b0_1030x1456mm", "JIS B0", dmPaperNone, 103000, 145600),
		*newPaperSize("jis_b1_728x1030mm", "JIS B1", dmPaperNone, 72800, 103000),
		*newPaperSize("jis_b2_515x728mm", "JIS B2", dmPaperNone, 51500, 72800),
		*newPaperSize("jis_b3_364x515mm", "JIS B3", dmPaperNone, 36400, 51500),
		*newPaperSize("jis_b4_257x364mm", "JIS B4", dmPaperNone, 25700, 36400),
		*newPaperSize("jis_b5_182x257mm", "JIS B5", dmPaperNone, 18200, 25700),
		*newPaperSize("jis_b6_128x182mm", "JIS B6", dmPaperB6JIS, 12800, 18200),
		*newPaperSize("jis_b7_91x128mm", "JIS B7", dmPaperNone, 9100, 12800),
		*newPaperSize("jis_b8_64x91mm", "JIS B8", dmPaperNone, 6400, 9100),
		*newPaperSize("jis_b9_45x64mm", "JIS B9", dmPaperNone, 4500, 6400),
		*newPaperSize("jis_b10_32x45mm", "JIS B10", dmPaperNone, 3200, 4500),
		*newPaperSize("jis_exec_216x330mm", "JIS Exec", dmPaperNone, 21600, 33000),
		*newPaperSize("jpn_chou2_111.1x146mm", "Chou2 Envelope", dmPaperNone, 11110, 14600),
		*newPaperSize("jpn_chou3_120x235mm", "Chou3 Envelope", dmPaperJapaneseEnvelopeChou3, 12000, 23500),
		*newPaperSize("jpn_chou4_90x205mm", "Chou4 Envelope", dmPaperJapaneseEnvelopeChou4, 9000, 20500),
		*newPaperSize("jpn_chou40_90x225mm", "Chou40 Envelope", dmPaperNone, 9000, 22500),
		*newPaperSize("jpn_hagaki_100x148mm", "Hagaki", dmPaperNone, 10000, 14800),
		*newPaperSize("jpn_kahu_240x322.1mm", "Kahu", dmPaperNone, 24000, 32210),
		*newPaperSize("jpn_kaku1_270x382mm", "Kaku1 Envelope", dmPaperNone, 27000, 38200),
		*newPaperSize("jpn_kaku2_240x332mm", "Kaku2 Envelope", dmPaperJapaneseEnvelopeKaku2, 24000, 33200),
		*newPaperSize("jpn_kaku3_216x277mm", "Kaku3 Envelope", dmPaperJapaneseEnvelopeKaku3, 21600, 27700),
		*newPaperSize("jpn_kaku4_197x267mm", "Kaku4 Envelope", dmPaperNone, 19700, 26700),
		*newPaperSize("jpn_kaku5_190x240mm", "Kaku5 Envelope", dmPaperNone, 19000, 24000),
		*newPaperSize("jpn_kaku7_142x205mm", "Kaku7 Envelope", dmPaperNone, 14200, 20500),
		*newPaperSize("jpn_kaku8_119x197mm", "Kaku8 Envelope", dmPaperNone, 11900, 19700),
		*newPaperSize("jpn_oufuku_148x200mm", "Oufuku", dmPaperNone, 14800, 20000),
		*newPaperSize("jpn_you4_105x235mm", "You4 Envelope", dmPaperJapaneseEnvelopeYou4, 10500, 23500),
		*newPaperSize("na_5x7_5x7in", "5x7in", dmPaperNone, 5*2540, 7*2540),
		*newPaperSize("na_6x9_6x9in", "6x9in", dmPaperNone, 6*2540, 9*2540),
		*newPaperSize("na_7x9_7x9in", "7x9in", dmPaperNone, 7*2540, 9*2540),
		*newPaperSize("na_9x11_9x11in", "9x11in", dmPaper9x11, 9*2540, 11*2540),
		*newPaperSize("na_10x11_10x11in", "10x11in", dmPaper10x11, 10*2540, 11*2540),
		*newPaperSize("na_10x13_10x13in", "10x13in", dmPaperNone, 10*2540, 13*2540),
		*newPaperSize("na_10x14_10x14in", "10x14in", dmPaper10x14, 10*2540, 14*2540),
		*newPaperSize("na_10x15_10x15in", "10x15in", dmPaperNone, 10*2540, 15*2540),
		*newPaperSize("na_11x12_11x12in", "11x12in", dmPaper12x11, 11*2540, 12*2540),
		*newPaperSize("na_11x15_11x15in", "11x15in", dmPaperNone, 11*2540, 15*2540),
		*newPaperSize("na_11x17_11x17in", "11x17in", dmPaper11x17, 11*2540, 15*2540),
		*newPaperSize("na_12x19_12x19in", "12x19in", dmPaperNone, 12*2540, 19*2540),
		*newPaperSize("na_a2_4.375x5.75in", "NA A2", dmPaperNone, 4.375*2540, 5.75*2540),
		*newPaperSize("na_arch-a_9x12in", "NA Arch-A", dmPaperNone, 9*2540, 12*2540),
		*newPaperSize("na_arch-b_12x18in", "NA Arch-B", dmPaperNone, 12*2540, 18*2540),
		*newPaperSize("na_arch-c_18x24in", "NA Arch-C", dmPaperNone, 18*2540, 24*2540),
		*newPaperSize("na_arch-d_24x36in", "NA Arch-D", dmPaperNone, 24*2540, 36*2540),
		*newPaperSize("na_arch-e2_26x38in", "NA Arch-E2", dmPaperNone, 26*2540, 38*2540),
		*newPaperSize("na_arch-e3_27x39in", "NA Arch-E3", dmPaperNone, 27*2540, 39*2540),
		*newPaperSize("na_arch-e_36x48in", "NA Arch-E", dmPaperNone, 36*2540, 48*2540),
		*newPaperSize("na_b-plus_12x19.17in", "NA B+", dmPaperBplus, 12*2540, 19.17*2540),
		*newPaperSize("na_c5_6.5x9.5in", "6.5x9.5in", dmPaperNone, 6.5*2540, 9.5*2540),
		*newPaperSize("na_c_17x22in", "17x22in", dmPaperCSheet, 17*2540, 22*2540),
		*newPaperSize("na_d_22x34in", "22x34in", dmPaperDSheet, 22*2540, 34*2540),
		*newPaperSize("na_e_34x44in", "34x44in", dmPaperESheet, 34*2540, 44*2540),
		*newPaperSize("na_edp_11x14in", "11x14in", dmPaperNone, 11*2540, 14*2540),
		*newPaperSize("na_eur-edp_12x14in", "12x14in", dmPaperNone, 12*2540, 14*2540),
		*newPaperSize("na_executive_7.25x10.5in", "NA Executive", dmPaperExecutive, 7.25*2540, 10.5*2540),
		*newPaperSize("na_f_44x68in", "44x68in", dmPaperNone, 44*2540, 68*25),
		*newPaperSize("na_fanfold-eur_8.5x12in", "German Std. Fanfold", dmPaperFanfoldGerman, 8.5*2540, 12*2540),
		*newPaperSize("na_fanfold-us_11x14.875in", "US Fanfold", dmPaperFanfoldUS, 11*2540, 14.875*2540),
		*newPaperSize("na_foolscap_8.5x13in", "German Legal Fanfold", dmPaperFanfoldLegal, 8.5*2540, 13*2540),
		*newPaperSize("na_govt-legal_8x13in", "NA Govt. Legal", dmPaperNone, 8*2540, 13*2540),
		*newPaperSize("na_govt-letter_8x10in", "NA Govt. Letter", dmPaperNone, 8*2540, 10*2540),
		*newPaperSize("na_index-3x5_3x5in", "3x5in", dmPaperNone, 3*2540, 5*2540),
		*newPaperSize("na_index-4x6-ext_6x8in", "6x8in", dmPaperNone, 6*2540, 8*2540),
		*newPaperSize("na_index-4x6_4x6in", "4x6in", dmPaperNone, 4*2540, 6*2540),
		*newPaperSize("na_index-5x8_5x8in", "5x8in", dmPaperNone, 5*2540, 8*2540),
		*newPaperSize("na_invoice_5.5x8.5in", "NA Invoice", dmPaperStatement, 5.5*2540, 8.5*2540),
		*newPaperSize("na_ledger_11x17in", "NA Ledger", dmPaperTabloid, 11*2540, 17*2540),
		*newPaperSize("na_legal-extra_9.5x15in", "9.5x15in", dmPaperNone, 9.5*2540, 15*2540),
		*newPaperSize("na_legal_8.5x14in", "NA Legal", dmPaperLegal, 8.5*2540, 14*2540),
		*newPaperSize("na_letter-extra_9.5x12in", "NA Letter Extra", dmPaperLetterExtra, 9.5*2540, 12*2540),
		*newPaperSize("na_letter-plus_8.5x12.69in", "NA Letter Plus", dmPaperNone, 8.5*2540, 12.69*2540),
		*newPaperSize("na_letter_8.5x11in", "NA Letter", dmPaperLetter, 8.5*2540, 11*2540),
		*newPaperSize("na_monarch_3.875x7.5in", "Monarch Envelope", dmPaperEnvMonarch, 3.875*2540, 7.5*2540),
		*newPaperSize("na_number-9_3.875x8.875in", "No. 9 Envelope", dmPaperEnv9, 3.875*2540, 8.875*2540),
		*newPaperSize("na_number-10_4.125x9.5in", "No. 10 Envelope", dmPaperEnv10, 4.125*2540, 9.5*2540),
		*newPaperSize("na_number-11_4.5x10.375in", "No. 11 Envelope", dmPaperEnv11, 4.5*2540, 10.375*2540),
		*newPaperSize("na_number-12_4.75x11in", "No. 12 Envelope", dmPaperEnv12, 4.75*2540, 11*2540),
		*newPaperSize("na_number-14_5x11.5in", "No. 14 Envelope", dmPaperEnv14, 5*2540, 11.5*2540),
		*newPaperSize("na_oficio_8.5x13.4in", "Officio", dmPaperNone, 8.5*2540, 13.4*2540),
		*newPaperSize("na_personal_3.625x6.5in", "Personal Envelope", dmPaperEnvPersonal, 3.625*2540, 6.5*2540),
		*newPaperSize("na_quarto_8.5x10.83in", "Quarto", dmPaperQuarto, 8.5*2540, 10.83*2540),
		*newPaperSize("na_super-a_8.94x14in", "Super-A", dmPaperNone, 8.94*2540, 14*2540),
		*newPaperSize("na_super-b_13x19in", "A3+", dmPaperNone, 13*2540, 19*2540),
		*newPaperSize("na_wide-format_30x42in", "Wide Format", dmPaperNone, 30*2540, 42*2540),
		*newPaperSize("oe_12x16_12x16in", "12x16in", dmPaperNone, 12*2540, 16*2540),
		*newPaperSize("oe_14x17_14x17in", "14x17in", dmPaperNone, 14*2540, 17*2540),
		*newPaperSize("oe_18x22_18x22in", "18x22in", dmPaperNone, 18*2540, 22*2540),
		*newPaperSize("oe_a2plus_17x24in", "17x24in", dmPaperNone, 17*2540, 24*2540),
		*newPaperSize("oe_business-card_2x3.5in", "Business Card", dmPaperNone, 2*2540, 3.5*2540),
		*newPaperSize("oe_photo-10r_10x12in", "Photo 10R", dmPaperNone, 10*2540, 12*2540),
		*newPaperSize("oe_photo-12r_12x15in", "Photo 12R", dmPaperNone, 12*2540, 15*2540),
		*newPaperSize("oe_photo-14x18_14x18in", "Photo 14x18in", dmPaperNone, 14*2540, 18*2540),
		*newPaperSize("oe_photo-16r_16x20in", "Photo 16R", dmPaperNone, 16*2540, 20*2540),
		*newPaperSize("oe_photo-20r_20x24in", "Photo 20R", dmPaperNone, 20*2540, 24*2540),
		*newPaperSize("oe_photo-22r_22x29.5in", "Photo 22R", dmPaperNone, 22*2540, 29.5*2540),
		*newPaperSize("oe_photo-22x28_22x28in", "Photo 22x28in", dmPaperNone, 22*2540, 28*2540),
		*newPaperSize("oe_photo-24r_24x31.5in", "Photo 24R", dmPaperNone, 24*2540, 31.5*2540),
		*newPaperSize("oe_photo-24x30_24x30in", "Photo 24x30in", dmPaperNone, 24*2540, 30*2540),
		*newPaperSize("oe_photo-30r_30x40in", "Photo 30R", dmPaperNone, 30*2540, 40*2540),
		*newPaperSize("oe_photo-l_3.5x5in", "Photo L", dmPaperNone, 3.5*2540, 5*2540),
		*newPaperSize("oe_photo-s8r_8x12in", "Photo S8R", dmPaperNone, 8*2540, 12*2540),
		*newPaperSize("oe_square-photo_4x4in", "Square Photo 4x4in", dmPaperNone, 4*2540, 4*2540),
		*newPaperSize("oe_square-photo_5x5in", "Square Photo 5x5in", dmPaperNone, 5*2540, 5*2540),
		*newPaperSize("om_16k_184x260mm", "16K 184x260mm", dmPaperNone, 18400, 26000),
		*newPaperSize("om_16k_195x270mm", "16K 195x270mm", dmPaperNone, 19500, 27000),
		*newPaperSize("om_business-card_55x85mm", "Business Card 55x85mm", dmPaperNone, 5500, 8500),
		*newPaperSize("om_business-card_55x91mm", "Business Card 55x91mm", dmPaperNone, 5500, 9100),
		*newPaperSize("om_card_54x86mm", "Card 54x86mm", dmPaperNone, 5400, 8600),
		*newPaperSize("om_card_54x92mm", "Card 54x92mm", dmPaperNone, 5400, 9200),
		*newPaperSize("om_dai-pa-kai_275x395mm", "Dai-Pa-Kai", dmPaperNone, 27500, 39500),
		*newPaperSize("om_dsc-photo_89x119mm", "DSC Photo 89x119mm", dmPaperNone, 8900, 11900),
		*newPaperSize("om_folio-sp_215x315mm", "Folio SP", dmPaperNone, 21500, 31500),
		*newPaperSize("om_folio_210x330mm", "Folio", dmPaperNone, 21000, 33000),
		*newPaperSize("om_invite_220x220mm", "Invite Envelope", dmPaperEnvelopeInvite, 22000, 22000),
		*newPaperSize("om_italian_110x230mm", "Italy Envelope", dmPaperEnvItaly, 11000, 23000),
		*newPaperSize("om_juuro-ku-kai_198x275mm", "Juuro-Ku-Kai", dmPaperNone, 19800, 27500),
		*newPaperSize("om_large-photo_200x300mm", "Large Photo", dmPaperNone, 20000, 30000),
		*newPaperSize("om_medium-photo_130x180mm", "Medium Photo", dmPaperNone, 13000, 18000),
		*newPaperSize("om_pa-kai_267x389mm", "Pa-Kai", dmPaperNone, 26700, 38900),
		*newPaperSize("om_photo-30x40_300x400mm", "Photo 30x40cm", dmPaperNone, 30000, 40000),
		*newPaperSize("om_photo-30x45_300x450mm", "Photo 30x45cm", dmPaperNone, 30000, 45000),
		*newPaperSize("om_photo-35x46_350x460mm", "Photo 35x46cm", dmPaperNone, 35000, 46000),
		*newPaperSize("om_photo-40x60_400x600mm", "Photo 40x60cm", dmPaperNone, 40000, 60000),
		*newPaperSize("om_photo-50x75_500x750mm", "Photo 50x75cm", dmPaperNone, 50000, 75000),
		*newPaperSize("om_photo-50x76_500x760mm", "Photo 50x76cm", dmPaperNone, 50000, 76000),
		*newPaperSize("om_photo-60x90_600x900mm", "Photo 60x90cm", dmPaperNone, 60000, 90000),
		*newPaperSize("om_small-photo_100x150mm", "Small Photo", dmPaperNone, 10000, 15000),
		*newPaperSize("om_square-photo_89x89mm", "Square Photo 89x89mm", dmPaperNone, 8900, 8900),
		*newPaperSize("om_wide-photo_100x200mm", "Wide Photo", dmPaperNone, 10000, 20000),
		*newPaperSize("prc_1_102x165mm", "PRC 1 Envelope", dmPaperPENV1, 10200, 16500),
		*newPaperSize("prc_2_102x176mm", "PRC 2 Envelope", dmPaperPENV2, 10200, 17600),
		*newPaperSize("prc_4_110x208mm", "PRC 4 Envelope", dmPaperPENV4, 11000, 20800),
		*newPaperSize("prc_6_120x320mm", "PRC 6 Envelope", dmPaperPENV6, 12000, 32000),
		*newPaperSize("prc_7_160x230mm", "PRC 7 Envelope", dmPaperPENV7, 16000, 23000),
		*newPaperSize("prc_8_120x309mm", "PRC 8 Envelope", dmPaperPENV8, 12000, 30900),
		*newPaperSize("prc_16k_146x215mm", "PRC 16K", dmPaperP16K, 14600, 21500),
		*newPaperSize("prc_32k_97x151mm", "PRC 32K", dmPaperP32K, 9700, 15100),
		*newPaperSize("roc_8k_10.75x15.5in", "ROC 8K", dmPaperNone, 10.75*2540, 15.5*2540),
		*newPaperSize("roc_16k_7.75x10.75in", "ROC 16K", dmPaperNone, 7.75*2540, 10.75*2540),
	}
}

func (s *paperSizes) add(size PaperSize) {
	s.sizes = append(s.sizes, size)
}

func (s *paperSizes) empty() {
	s.sizes = []PaperSize{}
}

func (s *paperSizes) findPaperSizeFromName(name string) *PaperSize {
	for _, ps := range s.sizes {
		if ps.name() == name {
			return &ps
		}
	}
	return nil
}

func (s *paperSizes) findPaperSize(psName, name string, width, height float32, margin Margin) *PaperSize {
	for _, ps := range s.sizes {
		if ps.name() == psName {
			return &ps
		} else if ps.name() == name {
			return &ps
		} else if ps.width() == width && ps.height() == height && ps.margin() == margin {
			return &ps
		}
	}
	return nil
}

// findPaperSizeFromPsName returns the PaperSize object that matches the
// PS name.
func (s *paperSizes) findPaperSizeFromPsName(psName string) *PaperSize {
	for _, ps := range s.sizes {
		if ps.psName() == psName {
			return &ps
		}
	}
	return nil
}

// findPaperSizeFromDmPaperSize returns the PaperSize object that matches the
// Windows-specified paper size.
func (s *paperSizes) findPaperSizeFromDmPaperSize(dm dmPaperSize) *PaperSize {
	for _, sz := range s.sizes {
		if sz.winSize == dm {
			return &sz
		}
	}
	return nil
}

func (s *paperSizes) findPaperSizeFromWindowsPaperSize(size fyne.Size) *PaperSize {
	for _, sz := range s.sizes {
		if int32(size.Width) == int32(sz.width())/10 && int32(size.Height) == int32(sz.height())/10 {
			return &sz
		}
	}
	return nil
}

func (s *paperSizes) isEmpty() bool {
	return len(s.sizes) == 0
}

// names returns the names of all paper sizes.
func (s *paperSizes) names() []string {
	var names []string
	for _, ps := range s.sizes {
		names = append(names, ps.name())
	}
	return names
}
