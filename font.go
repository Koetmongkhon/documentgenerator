package documentgenerator

// font size
var fontSizeHeader = 24.0
var fontSizeSubHeader = 16.0
var fontSizeContent = 14.0

const fontPlain = "FONT_CONTENT"
const fontBold = "FONT_HEADER"

func (pdfGen *pdfGenerator) SetFontBold(path string) error {
	return pdfGen.pdf.AddTTFFont(fontBold, path)
}

func (pdfGen *pdfGenerator) SetFontPlain(path string) error {
	return pdfGen.pdf.AddTTFFont(fontPlain, path)
}

func (pdfGen *pdfGenerator) SetFont(path string) error {
	err := pdfGen.SetFontBold(path)
	if err != nil {
		return err
	}
	err = pdfGen.SetFontPlain(path)
	if err != nil {
		return err
	}
	return nil
}

func (pdfGen *pdfGenerator) SetFontHeaderSize(size float64) {
	fontSizeHeader = size
}

func (pdfGen *pdfGenerator) SetFontSubHeaderSize(size float64) {
	fontSizeSubHeader = size
}

func (pdfGen *pdfGenerator) SetFontContentSize(size float64) {
	fontSizeContent = size
}

func (pdfGen *pdfGenerator) SetFontSize(size float64) {
	pdfGen.SetFontHeaderSize(size)
	pdfGen.SetFontSubHeaderSize(size)
	pdfGen.SetFontContentSize(size)
}
