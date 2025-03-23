//go:build !windows

package print

/*
// Printer contains a number of printer properties. This struct makes interacting
// with the printer easier.
type Printer struct {
	name         string
	location     string
	comment      string
	uri          string
	defPaperSize string
	pSizes       paperSizes
}

// newPrinter creates a Printer object based on CUPS ipp group info.
func newPrinter(ippGroup goipp.Group) *Printer {
	pr := &Printer{}
	for _, attr := range ippGroup.Attrs {
		if attr.Name == "printer-name" {
			pr.name = attr.Values.String()
		}
		if attr.Name == "printer-location" {
			pr.location = attr.Values.String()
		}
		if attr.Name == "printer-info" {
			pr.comment = attr.Values.String()
		}
		if attr.Name == "printer-more-info" {
			pr.uri = attr.Values.String()
			err := pr.retrievePaperSizes()
			if err != nil {
				fyne.LogError("Error retrieving paper sizes", err)
			}
		}
		if attr.Name == "media-default" {
			pr.defPaperSize = attr.Values.String()
		}
	}
	return pr
}

func (pr *Printer) Comment() string {
	return pr.comment
}

// Location retrieves the printer's location as returned by CUPS.
func (pr *Printer) Location() string {
	return pr.location
}

// Name retrieves the printer's name as returned by CUPS.
func (pr *Printer) Name() string {
	return pr.name
}

// defaultPrinterSize returns the PaperSize corresponding to the default paper size for the printer.
func (p *Printer) defaultPaperSize() *PaperSize {
	return stdPaperSizes.findPaperSizeFromName(p.defPaperSize)
}

// paperSizes returns the paper sizes for the printer. The first time
// this method is called, the paper sizes are retrieved.
func (p *Printer) paperSizes() (paperSizes, error) {
	return p.pSizes, nil
}

func (p *Printer) retrievePaperSizes() error {
	mcds, err := getResponseGroups(goipp.OpCupsGetPrinters,
		p.uri, "media-col-database")
	if err != nil {
		fyne.LogError("Error getting CUPS media-col-database", err)
		return err
	}
	papers, err := p.getMediaSupported()
	if err != nil {
		return err
	}
	for _, mcd := range *mcds {
		if mcd.Tag.String() == "printer-attributes-tag" {
			var w, h, mt, mb, ml, mr float32
			for _, attr := range mcd.Attrs {
				for i, mediaSize := range attr.Values {
					//					fmt.Printf("mediaSize: %v\n", mediaSize)
					if col, ok := mediaSize.V.(goipp.Collection); ok {
						for _, param := range col {
							switch param.Name {
							case "media-size":
								if sz, ok := param.Values[0].V.(goipp.Collection); ok {
									for _, xy := range sz {
										switch xy.Name {
										case "x-dimension":
											switch x := xy.Values[0].V.(type) {
											case goipp.Integer:
												w = float32(x)
												fmt.Printf("x: %d\n", x)
											case goipp.Range:
												fmt.Printf("x: %d-%d\n", x.Lower, x.Upper)
											}
										case "y-dimension":
											switch y := xy.Values[0].V.(type) {
											case goipp.Integer:
												h = float32(y)
											case goipp.Range:
												fmt.Printf("y: %d-%d\n", y.Lower, y.Upper)
											}
										}
									}
								}

							case "media-top-margin":
								if mrg, ok := param.Values[0].V.(goipp.Integer); ok {
									mt = float32(mrg)
								}

							case "media-bottom-margin":
								if mrg, ok := param.Values[0].V.(goipp.Integer); ok {
									mb = float32(mrg)
								}
							case "media-left-margin":
								if mrg, ok := param.Values[0].V.(goipp.Integer); ok {
									ml = float32(mrg)
								}

							case "media-right-margin":
								if mrg, ok := param.Values[0].V.(goipp.Integer); ok {
									mr = float32(mrg)
								}
							}
						}
						ps := stdPaperSizes.findPaperSize(papers[i], papers[i], w, h,
							newMargin(mt, mb, ml, mr))
						if ps == nil {
							ps = newPaperSizeWithMargin(papers[i], papers[i], 0, w, h, newMargin(mt, mb, ml, mr))
						}
						p.pSizes.add(*ps)
						fmt.Printf("Added ps: %v\n", ps)
					}
				}
			}
		}
	}
	return nil
}

func (p *Printer) getMediaSupported() ([]string, error) {
	ms, err := getResponseGroups(goipp.OpCupsGetPrinters,
		p.uri, "media-supported")
	if err != nil {
		fyne.LogError("Error getting CUPS media-supported", err)
		return []string{}, err
	}
	var mSupported []string
	for _, m := range *ms {
		if m.Tag.String() == "printer-attributes-tag" {
			for _, attr := range m.Attrs {
				//				fmt.Printf("mediaSupported: %v\n", attr.Values.String())
				for _, val := range attr.Values {
					mSupported = append(mSupported, val.V.String())
					fmt.Printf("%v\n", val.V.String())
				}
			}
		}
	}
	return mSupported, nil
}
*/
