package print

// paperSizes contains a slice of PaperSize objects.
type paperSizes struct {
	sizes []PaperSize
}

// stdPaperSizes contains all of the IPP media sizes registered with IANA
// as of 2024-10-09.
// newPaperSize arguments:
//	1: the IPP media size name.
//	2: the English name for the media size. The name is translated in newPaperSize
// to the system language if a translation exists.
//	3: width of the paper in 100ths of a mm for the paper in portrait mode.
//  4: height of the paper in 100ths of a mm for the paper in portrait mode. That is,
// the height is always equal to or greater than the width.
var stdPaperSizes paperSizes

func init() {
	stdPaperSizes.sizes = []PaperSize{
		newPaperSize("asme_f_28x40in", "ASME_F", 28*2540, 40*2540),
		newPaperSize("iso_2a0_1189x1682mm", "2xA0", 118900, 168200),
		newPaperSize("iso_a0_841x1189mm", "A0", 84100, 118900),
		newPaperSize("iso_a0x3_1189x2523mm", "A0x3", 118900, 252300),
		newPaperSize("iso_a0x4_1189x3003mm", "A0x4", 118900, 300300),
		newPaperSize("iso_a1_594x841mm", "A1", 59400, 84100),
		newPaperSize("iso_a1x3_841x1783mm", "A1x3", 84100, 178300),
		newPaperSize("iso_a1x4_841x2378mm", "A1x4", 84100, 237800),
		newPaperSize("iso_a2_420x594mm", "A2", 42000, 59400),
		newPaperSize("iso_a2x3_420x1783mm", "A2x3", 42000, 178300),
		newPaperSize("iso_a2x4_420x2378mm", "A2x4", 42000, 237800),
		newPaperSize("iso_a2x5_420x2970mm", "A2x5", 42000, 297000),
		newPaperSize("iso_a3-extra_322x445mm", "A3 Extra", 32200, 44500),
		newPaperSize("iso_a3_297x420mm", "A3", 29700, 42000),
		newPaperSize("iso_a3x3_420x891mm", "A3x3", 42000, 89100),
		newPaperSize("iso_a3x4_420x1189mm", "A3x4", 42000, 118900),
		newPaperSize("iso_a3x5_420x1486mm", "A3x5", 42000, 148600),
		newPaperSize("iso_a3x6_420x1783mm", "A3x6", 42000, 178300),
		newPaperSize("iso_a3x7_420x2080mm", "A3x7", 42000, 208000),
		newPaperSize("iso_a4-extra_235.5x322.3mm", "A4 Extra", 23550, 32230),
		newPaperSize("iso_a4-tab_225x297mm", "A4 Tab", 22500, 29700),
		newPaperSize("iso_a4_297x420mm", "A4", 21000, 29700),
		newPaperSize("iso_a4x3_297x630mm", "A4x3", 21000, 63000),
		newPaperSize("iso_a4x4_297x794mm", "A4x4", 21000, 79400),
		newPaperSize("iso_a4x5_297x958mm", "A4x5", 21000, 95800),
		newPaperSize("iso_a4x6_297x1122mm", "A4x6", 21000, 112200),
		newPaperSize("iso_a4x7_297x1286mm", "A4x7", 21000, 128600),
		newPaperSize("iso_a4x8_297x1450mm", "A4x8", 21000, 145000),
		newPaperSize("iso_a4x9_297x1892mm", "A4x9", 21000, 189200),
		newPaperSize("iso_a5-extra_174x235mm", "A5 Extra", 17400, 23500),
		newPaperSize("iso_a5_148x210mm", "A5", 14800, 21000),
		newPaperSize("iso_a6_105x148mm", "A6", 10500, 14800),
		newPaperSize("iso_a7_74x105mm", "A7", 7400, 10500),
		newPaperSize("iso_a8_52x74mm", "A8", 5200, 7400),
		newPaperSize("iso_a9_37x52mm", "A9", 3700, 5200),
		newPaperSize("iso_a10_26x37mm", "A10", 2600, 3700),
		newPaperSize("iso_b0_1000x1414mm", "B0", 100000, 141400),
		newPaperSize("iso_b1_707x1000mm", "B1", 70700, 100000),
		newPaperSize("iso_b2_500x707mm", "B2", 50000, 70700),
		newPaperSize("iso_b3_353x500mm", "B3", 35300, 50000),
		newPaperSize("iso_b4_250x353mm", "B4", 25000, 35300),
		newPaperSize("iso_b5-extra_201x276mm", "B5 Extra", 20100, 27600),
		newPaperSize("iso_b5_176x250mm", "B5", 17600, 25000),
		newPaperSize("iso_b6_125x176mm", "B6", 12500, 17600),
		newPaperSize("iso_b6c4_125x324mm", "B6C4", 12500, 32400),
		newPaperSize("iso_b7_88x125mm", "B7", 8800, 12500),
		newPaperSize("iso_b8_62x88mm", "B8", 6200, 8800),
		newPaperSize("iso_b9_44x62mm", "B9", 4400, 6200),
		newPaperSize("iso_b10_31x44mm", "B10", 3100, 4400),
		newPaperSize("iso_c0_917x1297mm", "C0", 91700, 129700),
		newPaperSize("iso_c1_648x917mm", "C1", 64800, 91700),
		newPaperSize("iso_c2_458x648mm", "C2", 45800, 64800),
		newPaperSize("iso_c3_324x458mm", "C3 Envelope", 32400, 45800),
		newPaperSize("iso_c4_229x324mm", "C4 Envelope", 22900, 32400),
		newPaperSize("iso_c5_162x229mm", "C5 Envelope", 16200, 22900),
		newPaperSize("iso_c6_114x162mm", "C6 Envelope", 11400, 16200),
		newPaperSize("iso_c6c5_114x229mm", "C6C5 Envelope", 11400, 22900),
		newPaperSize("iso_c7_77x114mm", "C7 Envelope", 7700, 11400),
		newPaperSize("iso_c7c6_81x162mm", "C7C6 Envelope", 8100, 16200),
		newPaperSize("iso_c8_56x77mm", "C8 Envelope", 5600, 7700),
		newPaperSize("iso_c9_38x56mm", "C9 Envelope", 3800, 5600),
		newPaperSize("iso_c10_29x38mm", "C10 Envelope", 2900, 3800),
		newPaperSize("iso_dl_110x220mm", "DL Emvelope", 11000, 22000),
		newPaperSize("iso_id-1_53.98x85.6mm", "ID-1", 5398, 8560), // same as "om_card_54x86mm"
		newPaperSize("iso_ra0_860x1220mm", "RA0", 86000, 122000),
		newPaperSize("iso_ra1_610x860mm", "RA1", 61000, 86000),
		newPaperSize("iso_ra2_420x610mm", "RA2", 42000, 61000),
		newPaperSize("iso_ra3_305x420mm", "RA3", 30500, 42000),
		newPaperSize("iso_ra4_210x305mm", "RA4", 21000, 30500),
		newPaperSize("iso_sra0_900x1280mm", "RA5", 14800, 21000),
		newPaperSize("iso_sra1_640x900mm", "SRA1", 64000, 90000),
		newPaperSize("iso_sra2_450x640mm", "SRA2", 45000, 64000),
		newPaperSize("iso_sra3_320x450mm", "SRA3", 32000, 45000),
		newPaperSize("iso_sra4_225x320mm", "SRA4", 22500, 32000),
		newPaperSize("jis_b0_1030x1456mm", "JIS B0", 103000, 145600),
		newPaperSize("jis_b1_728x1030mm", "JIS B1", 72800, 103000),
		newPaperSize("jis_b2_515x728mm", "JIS B2", 51500, 72800),
		newPaperSize("jis_b3_364x515mm", "JIS B3", 36400, 51500),
		newPaperSize("jis_b4_257x364mm", "JIS B4", 25700, 36400),
		newPaperSize("jis_b5_182x257mm", "JIS B5", 18200, 25700),
		newPaperSize("jis_b6_128x182mm", "JIS B6", 12800, 18200),
		newPaperSize("jis_b7_91x128mm", "JIS B7", 9100, 12800),
		newPaperSize("jis_b8_64x91mm", "JIS B8", 6400, 9100),
		newPaperSize("jis_b9_45x64mm", "JIS B9", 4500, 6400),
		newPaperSize("jis_b10_32x45mm", "JIS B10", 3200, 4500),
		newPaperSize("jis_exec_216x330mm", "JIS Exec", 21600, 33000),
		newPaperSize("jpn_chou2_111.1x146mm", "Chou2", 11110, 14600),
		newPaperSize("jpn_chou3_120x235mm", "Chou3", 12000, 23500),
		newPaperSize("jpn_chou4_90x205mm", "Chou4", 9000, 20500),
		newPaperSize("jpn_chou40_90x225mm", "Chou40", 9000, 22500),
		newPaperSize("jpn_hagaki_100x148mm", "Hagaki", 10000, 14800),
		newPaperSize("jpn_kahu_240x322.1mm", "Kahu", 24000, 32210),
		newPaperSize("jpn_kaku1_270x382mm", "Kaku1", 27000, 38200),
		newPaperSize("jpn_kaku2_240x332mm", "Kaku2", 24000, 33200),
		newPaperSize("jpn_kaku3_216x277mm", "Kaku3", 21600, 27700),
		newPaperSize("jpn_kaku4_197x267mm", "Kaku4", 19700, 26700),
		newPaperSize("jpn_kaku5_190x240mm", "Kaku5", 19000, 24000),
		newPaperSize("jpn_kaku7_142x205mm", "Kaku7", 14200, 20500),
		newPaperSize("jpn_kaku8_119x197mm", "Kaku8", 11900, 19700),
		newPaperSize("jpn_oufuku_148x200mm", "Oufuku", 14800, 20000),
		newPaperSize("jpn_you4_105x235mm", "You4", 10500, 23500),
		newPaperSize("na_5x7_5x7in", "5x7in", 5*2540, 7*2540),
		newPaperSize("na_6x9_6x9in", "6x9in", 6*2540, 9*2540),
		newPaperSize("na_7x9_7x9in", "7x9in", 7*2540, 9*2540),
		newPaperSize("na_9x11_9x11in", "9x11in", 9*2540, 11*2540),
		newPaperSize("na_10x11_10x11in", "10x11in", 10*2540, 11*2540),
		newPaperSize("na_10x13_10x13in", "10x13in", 10*2540, 13*2540),
		newPaperSize("na_10x14_10x14in", "10x14in", 10*2540, 14*2540),
		newPaperSize("na_10x15_10x15in", "10x15in", 10*2540, 15*2540),
		newPaperSize("na_11x12_11x12in", "11x12in", 11*2540, 12*2540),
		newPaperSize("na_11x15_11x15in", "11x1in", 11*2540, 15*2540),
		newPaperSize("na_12x19_12x19in", "12x19in", 12*2540, 19*2540),
		newPaperSize("na_a2_4.375x5.75in", "NA A2", 4.375*2540, 5.75*2540),
		newPaperSize("na_arch-a_9x12in", "NA Arch-A", 9*2540, 12*2540),
		newPaperSize("na_arch-b_12x18in", "NA Arch-B", 12*2540, 18*2540),
		newPaperSize("na_arch-c_18x24in", "NA Arch-C", 18*2540, 24*2540),
		newPaperSize("na_arch-d_24x36in", "NA Arch-D", 24*2540, 36*2540),
		newPaperSize("na_arch-e2_26x38in", "NA Arch-E2", 26*2540, 38*2540),
		newPaperSize("na_arch-e3_27x39in", "NA Arch-E3", 27*2540, 39*2540),
		newPaperSize("na_arch-e_36x48in", "NA Arch-E", 36*2540, 48*2540),
		newPaperSize("na_b-plus_12x19.17in", "NA B+", 12*2540, 19.17*2540),
		newPaperSize("na_c5_6.5x9.5in", "6.5x9.5in", 6.5*2540, 9.5*2540),
		newPaperSize("na_c_17x22in", "17x22in", 17*2540, 22*2540),
		newPaperSize("na_d_22x34in", "22x34in", 22*2540, 34*2540),
		newPaperSize("na_e_34x44in", "34x44in", 34*2540, 44*2540),
		newPaperSize("na_edp_11x14in", "11x14in", 11*2540, 14*2540),
		newPaperSize("na_eur-edp_12x14in", "12x14in", 12*2540, 14*2540),
		newPaperSize("na_executive_7.25x10.5in", "NA Executive", 7.25*2540, 10.5*2540),
		newPaperSize("na_f_44x68in", "44x68in", 44*2540, 68*25),
		newPaperSize("na_fanfold-eur_8.5x12in", "German Std. Fanfold", 8.5*2540, 12*2540),
		newPaperSize("na_fanfold-us_11x14.875in", "US Fanfold", 11*2540, 14.875*2540),
		newPaperSize("na_foolscap_8.5x13in", "German Legal Fanfold", 8.5*2540, 11*2540),
		newPaperSize("na_govt-legal_8x13in", "NA Govt. Legal", 8*2540, 13*2540),
		newPaperSize("na_govt-letter_8x10in", "NA Govt. Letter", 8*2540, 10*2540),
		newPaperSize("na_index-3x5_3x5in", "3.5x5in", 3.5*2540, 5*2540),
		newPaperSize("na_index-4x6-ext_6x8in", "6x8in", 6*2540, 8*2540),
		newPaperSize("na_index-4x6_4x6in", "4x6in", 4*2540, 6*2540),
		newPaperSize("na_index-5x8_5x8in", "5x8in", 5*2540, 8*2540),
		newPaperSize("na_invoice_5.5x8.5in", "NA Invoice", 5.5*2540, 8.5*2540),
		newPaperSize("na_ledger_11x17in", "NA Ledger", 11*2540, 17*2540),
		newPaperSize("na_legal-extra_9.5x15in", "9.5x15in", 9.5*2540, 15*2540),
		newPaperSize("na_legal_8.5x14in", "NA Legal", 8.5*2540, 14*2540),
		newPaperSize("na_letter-extra_9.5x12in", "NA Letter Extra", 9.5*2540, 12*2540),
		newPaperSize("na_letter-plus_8.5x12.69in", "NA Letter Plus", 8.5*2540, 12.69*2540),
		newPaperSize("na_letter_8.5x11in", "NA Letter", 8.5*2540, 11*2540),
		newPaperSize("na_monarch_3.875x7.5in", "Monarch Envelope", 3.875*2540, 7.5*2540),
		newPaperSize("na_number-9_3.875x8.875in", "No. 9 Envelope", 3.875*2540, 8.875*2540),
		newPaperSize("na_number-10_4.125x9.5in", "No. 10 Envelope", 4.125*2540, 9.5*2540),
		newPaperSize("na_number-11_4.5x10.375in", "No. 11 Envelope", 4.5*2540, 10.375*2540),
		newPaperSize("na_number-12_4.75x11in", "No. 12 Envelope", 4.75*2540, 11*2540),
		newPaperSize("na_number-14_5x11.5in", "No. 14 Envelope", 5*2540, 11.5*2540),
		newPaperSize("na_oficio_8.5x13.4in", "Officio", 8.5*2540, 13.4*2540),
		newPaperSize("na_personal_3.625x6.5in", "Personal", 3.625*2540, 6.5*2540),
		newPaperSize("na_quarto_8.5x10.83in", "Quarto", 8.5*2540, 10.83*2540),
		newPaperSize("na_super-a_8.94x14in", "NA Super-A", 8.94*2540, 14*2540),
		newPaperSize("na_super-b_13x19in", "A3+", 13*2540, 19*2540),
		newPaperSize("na_wide-format_30x42in", "Wide Format", 30*2540, 42*2540),
		newPaperSize("oe_12x16_12x16in", "12x16in", 12*2540, 16*2540),
		newPaperSize("oe_14x17_14x17in", "14x17in", 14*2540, 17*2540),
		newPaperSize("oe_18x22_18x22in", "18x22in", 18*2540, 22*2540),
		newPaperSize("oe_a2plus_17x24in", "17x24in", 17*2540, 24*2540),
		newPaperSize("oe_business-card_2x3.5in", "Business Card", 2*2540, 3.5*2540),
		newPaperSize("oe_photo-10r_10x12in", "Photo 10R", 10*2540, 12*2540),
		newPaperSize("oe_photo-12r_12x15in", "Photo 12R", 12*2540, 15*2540),
		newPaperSize("oe_photo-14x18_14x18in", "Photo 14x18in", 14*2540, 18*2540),
		newPaperSize("oe_photo-16r_16x20in", "Photo 16R", 16*2540, 20*2540),
		newPaperSize("oe_photo-20r_20x24in", "Photo 20R", 20*2540, 24*2540),
		newPaperSize("oe_photo-22r_22x29.5in", "Photo 22R", 22*2540, 29.5*2540),
		newPaperSize("oe_photo-22x28_22x28in", "Photo 22x28in", 22*2540, 28*2540),
		newPaperSize("oe_photo-24r_24x31.5in", "Photo 24R", 24*2540, 31.5*2540),
		newPaperSize("oe_photo-24x30_24x30in", "Photo 24x30in", 24*2540, 30*2540),
		newPaperSize("oe_photo-30r_30x40in", "Photo 30R", 30*2540, 40*2540),
		newPaperSize("oe_photo-l_3.5x5in", "Photo L", 3.5*2540, 5*2540),
		newPaperSize("oe_photo-s8r_8x12in", "Photo S8R", 8*2540, 12*2540),
		newPaperSize("oe_square-photo_4x4in", "Square Photo 4x4in", 4*2540, 4*2540),
		newPaperSize("oe_square-photo_5x5in", "Square Photo 5x5in", 5*2540, 5*2540),
		newPaperSize("om_16k_184x260mm", "16K 184x260mm", 18400, 26000),
		newPaperSize("om_16k_195x270mm", "16K 195x270mm", 19500, 27000),
		newPaperSize("om_business-card_55x85mm", "Business Card 55x85mm", 5500, 8500),
		newPaperSize("om_business-card_55x91mm", "Business Card 55x91mm", 5500, 9100),
		newPaperSize("om_card_54x86mm", "Card 54x86mm", 5400, 8600), // same as "iso_id-1_53.98x85.6mm"
		newPaperSize("om_card_54x92mm", "Card 54x92mm", 5400, 9200),
		newPaperSize("om_dai-pa-kai_275x395mm", "Dai-Pa-Kai", 27500, 39500),
		newPaperSize("om_dsc-photo_89x119mm", "DSC Photo 89x119mm", 8900, 11900),
		newPaperSize("om_folio-sp_215x315mm", "Folio SP", 21500, 31500),
		newPaperSize("om_folio_210x330mm", "Folio", 21000, 33000),
		newPaperSize("om_invite_220x220mm", "Invite", 22000, 22000),
		newPaperSize("om_italian_110x230mm", "Italian", 11000, 23000),
		newPaperSize("om_juuro-ku-kai_198x275mm", "Juuro-Ku-Kai", 19800, 27500),
		newPaperSize("om_large-photo_200x300mm", "Large Photo", 20000, 30000),
		newPaperSize("om_medium-photo_130x180mm", "Medium Photo", 13000, 18000),
		newPaperSize("om_pa-kai_267x389mm", "Pa-Kai", 26700, 38900),
		newPaperSize("om_photo-30x40_300x400mm", "Photo 30x40cm", 30000, 40000),
		newPaperSize("om_photo-30x45_300x450mm", "Photo 30x45cm", 30000, 45000),
		newPaperSize("om_photo-35x46_350x460mm", "Photo 35x46cm", 35000, 46000),
		newPaperSize("om_photo-40x60_400x600mm", "Photo 40x60cm", 40000, 60000),
		newPaperSize("om_photo-50x75_500x750mm", "Photo 50x75cm", 50000, 75000),
		newPaperSize("om_photo-50x76_500x760mm", "Photo 50x76cm", 50000, 76000),
		newPaperSize("om_photo-60x90_600x900mm", "Photo 60x90cm", 60000, 90000),
		newPaperSize("om_small-photo_100x150mm", "Small Photo", 10000, 15000),
		newPaperSize("om_square-photo_89x89mm", "Square Photo 89x89mm", 8900, 8900),
		newPaperSize("om_wide-photo_100x200mm", "Wide Photo", 10000, 20000),
		newPaperSize("prc_1_102x165mm", "PRC 1 Envelope", 10200, 16500),
		newPaperSize("prc_2_102x176mm", "PRC 2 Envelope", 10200, 17600),
		newPaperSize("prc_4_110x208mm", "PRC 4 Envelope", 11000, 20800),
		newPaperSize("prc_6_120x320mm", "PRC 6 Envelope", 12000, 32000),
		newPaperSize("prc_7_160x230mm", "PRC 7 Envelope", 16000, 23000),
		newPaperSize("prc_8_120x309mm", "PRC 8 Envelope", 12000, 30900),
		newPaperSize("prc_16k_146x215mm", "PRC 16K", 14600, 21500),
		newPaperSize("prc_32k_97x151mm", "PRC 32K", 9700, 15100),
		newPaperSize("roc_8k_10.75x15.5in", "ROC 8K", 10.75*2540, 15.5*2540),
		newPaperSize("roc_16k_7.75x10.75in", "ROC 16K", 7.75*2540, 10.75*2540),
	}
}
