package documentgenerator

import (
	"fmt"

	"github.com/signintech/gopdf"
)

// pdf constant
const Left = 8    //001000
const Top = 4     //000100
const Right = 2   //000010
const Bottom = 1  //000001
const Center = 16 //010000
const Middle = 32 //100000
const All = 15    //001111
const EndX = 590.0
const EndY = 830.0
const pageMargin = 20

var tableSection = 0.0
var summarySection = 700.0
var pageNumber = 0

type pdfGenerator struct {
	pdf *gopdf.GoPdf
}

func newGoPdf() PdfGenerator {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4, Unit: gopdf.UnitPT})
	return &pdfGenerator{pdf: &pdf}
}

func (pdfGen *pdfGenerator) textHeader() error {
	return pdfGen.pdf.SetFont(fontBold, "", fontSizeHeader)
}

func (pdfGen *pdfGenerator) textSubHeader() error {
	return pdfGen.pdf.SetFont(fontBold, "", fontSizeSubHeader)
}

func (pdfGen *pdfGenerator) textContent() error {
	return pdfGen.pdf.SetFont(fontPlain, "", fontSizeContent)
}

func (pdfGen *pdfGenerator) writeSubAndContent(sub, content string) error {
	err := pdfGen.textSubHeader()
	if err != nil {
		return err
	}
	err = pdfGen.pdf.Cell(nil, sub)
	if err != nil {
		return err
	}

	err = pdfGen.textContent()
	if err != nil {
		return err
	}
	pdfGen.pdf.SetXY(pdfGen.pdf.GetX()+5, pdfGen.pdf.GetY()+1.5)
	return pdfGen.pdf.Cell(nil, content)
}

/*
Pdf Header section
*/
func (pdfGen *pdfGenerator) headerAndFooter(doc DocumentInformation, seller, customer PersonalInformation) (err error) {
	pdfGen.pdf.AddHeader(func() {
		err = pdfGen.textHeader()
		if err != nil {
			return
		}

		// set document name
		pdfGen.pdf.SetXY(0, 0)
		err = pdfGen.pdf.CellWithOption(
			&gopdf.Rect{
				W: EndX,
				H: fontSizeHeader * 2,
			},
			doc.Header,
			gopdf.CellOption{
				Align: Center + Middle,
			},
		)

		// write issue date and document id
		pdfGen.pdf.SetXY(pageMargin, fontSizeHeader*1.75)
		pdfGen.writeSubAndContent("วันที่:", doc.IssueDate)
		pdfGen.pdf.SetXY(EndX/1.75, fontSizeHeader*1.75)
		pdfGen.writeSubAndContent("เลขที่:", doc.Id)

		// write buyer information
		pdfGen.pdf.SetXY(pageMargin, fontSizeHeader*3)
		pdfGen.writeSubAndContent("ผู้ขาย:", seller.Name)
		pdfGen.pdf.Br(fontSizeSubHeader * 1.1)
		pdfGen.writeSubAndContent("ที่อยู่:", seller.Address)
		pdfGen.pdf.Br(fontSizeSubHeader * 1.1)
		pdfGen.writeSubAndContent("เลขประจำตัวผู้เสียภาษี:", seller.TaxId)
		pdfGen.pdf.Br(fontSizeSubHeader * 1.1)
		pdfGen.writeSubAndContent("อีเมล:", seller.Email)
		pdfGen.pdf.SetXY(EndX/2, pdfGen.pdf.GetY()-1.5)
		pdfGen.writeSubAndContent("โทรศัทพ์:", seller.Phone)
		pdfGen.pdf.Br(fontSizeSubHeader * 1.25)

		// write customer information
		pdfGen.writeSubAndContent("ผู้ซื้อ:", customer.Name)
		pdfGen.pdf.Br(fontSizeSubHeader)
		pdfGen.writeSubAndContent("ที่อยู่:", customer.Address)
		pdfGen.pdf.Br(fontSizeSubHeader)
		pdfGen.writeSubAndContent("เลขประจำตัวผู้เสียภาษี:", customer.TaxId)
		pdfGen.pdf.Br(fontSizeSubHeader)
		pdfGen.writeSubAndContent("อีเมล:", customer.Email)
		pdfGen.pdf.SetXY(EndX/2, pdfGen.pdf.GetY()-1.5)
		pdfGen.writeSubAndContent("โทรศัทพ์:", customer.Phone)
		tableSection = pdfGen.pdf.GetY()
	})

	pdfGen.pdf.AddFooter(func() {
		err = pdfGen.pdf.SetFont(fontPlain, "", fontSizeContent)
		if err != nil {
			return
		}
		pdfGen.pdf.SetXY(EndX-(pageMargin*1.05), EndY-pageMargin)
		pdfGen.pdf.Cell(nil, fmt.Sprintf("%d", pageNumber))
	})
	return nil
}

func (pdfGen *pdfGenerator) genHeaderTable() {
	pdfGen.pdf.CellWithOption(&gopdf.Rect{W: 30, H: 20}, "ลำดับ", gopdf.CellOption{Align: Middle + Center, Border: All})
	pdfGen.pdf.CellWithOption(&gopdf.Rect{W: 370, H: 20}, "รายละเอียด", gopdf.CellOption{Align: Middle + Center, Border: All})
	pdfGen.pdf.CellWithOption(&gopdf.Rect{W: 50, H: 20}, "จำนวน", gopdf.CellOption{Align: Middle + Center, Border: All})
	pdfGen.pdf.CellWithOption(&gopdf.Rect{W: 55, H: 20}, "ราคาต่อหน่วย", gopdf.CellOption{Align: Middle + Center, Border: All})
	pdfGen.pdf.CellWithOption(&gopdf.Rect{W: 50, H: 20}, "รวม", gopdf.CellOption{Align: Middle + Center, Border: All})
}

func (pdfGen *pdfGenerator) genTableItem(rowNo int, item Item) {
	no := ""
	if rowNo != 0 {
		no = fmt.Sprintf("%d", rowNo)
	}

	quantity := ""
	if item.Quantity != 0 {
		quantity = fmt.Sprintf("%d", item.Quantity)
	}

	unitPrice := ""
	if item.UnitPrice != 0 {
		unitPrice = fmt.Sprintf("%v", item.UnitPrice)
	}

	amount := ""
	if item.Quantity != 0 && item.UnitPrice != 0 {
		item.Amount = float64(item.Quantity) * item.UnitPrice
	}
	if item.Amount != 0 {
		amount = fmt.Sprintf("%v", item.Amount)
	}

	// write row number
	pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: 30, H: fontSizeContent * 1.3},
		no,
		gopdf.CellOption{
			Align:  Middle + Center,
			Border: Left + Right,
		},
	)
	// write item description
	pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: 370, H: fontSizeContent * 1.3},
		fmt.Sprintf(" %s", item.Name),
		gopdf.CellOption{
			Align:  Left + Middle,
			Border: Right,
		},
	)
	// write quantity
	pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: 50, H: fontSizeContent * 1.3},
		quantity,
		gopdf.CellOption{
			Align:  Middle + Center,
			Border: Right,
		},
	)
	// write unit price
	pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: 55, H: fontSizeContent * 1.3},
		unitPrice,
		gopdf.CellOption{
			Align:  Middle + Center,
			Border: Right,
		},
	)
	// write price
	pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: 50, H: fontSizeContent * 1.3},
		amount,
		gopdf.CellOption{
			Align:  Middle + Center,
			Border: Right,
		},
	)
}

func (pdfGen *pdfGenerator) genLastRow() {
	pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: 30, H: fontSizeContent * 1.3},
		"",
		gopdf.CellOption{
			Align:  Middle + Center,
			Border: Left + Right + Bottom,
		},
	)
	pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: 370, H: fontSizeContent * 1.3},
		"",
		gopdf.CellOption{
			Align:  Left,
			Border: Right + Bottom,
		},
	)
	pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: 50, H: fontSizeContent * 1.3},
		"",
		gopdf.CellOption{
			Align:  Middle + Center,
			Border: Right + Bottom,
		},
	)
	pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: 55, H: fontSizeContent * 1.3},
		"",
		gopdf.CellOption{
			Align:  Middle + Center,
			Border: Right + Bottom,
		},
	)
	pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: 50, H: fontSizeContent * 1.3},
		"",
		gopdf.CellOption{
			Align: Middle + Center, Border: Right + Bottom,
		},
	)
}

func (pdfGen *pdfGenerator) addNewPage() {
	pageNumber += 1
	pdfGen.pdf.AddPage()
	pdfGen.pdf.SetXY(pageMargin, tableSection+(fontSizeSubHeader*1.5))
	pdfGen.genHeaderTable()
}

/*
Create table with item section
*/
func (pdfGen *pdfGenerator) table(items []Item) error {
	spacing := fontSizeContent * 1.3

	// table header
	pdfGen.addNewPage()
	pdfGen.pdf.Br(spacing)

	// row items
	for i, v := range items {
		if pdfGen.pdf.GetY()+(fontSizeContent*2.6) >= summarySection {
			pdfGen.genLastRow()
			pdfGen.addNewPage()
			pdfGen.pdf.Br(spacing)
		}
		pdfGen.genTableItem(i+1, v)
		pdfGen.pdf.Br(spacing)
	}

	for pdfGen.pdf.GetY()+(fontSizeContent*2.6) < summarySection {
		pdfGen.genTableItem(0, Item{})
		pdfGen.pdf.Br(spacing)
	}

	pdfGen.genLastRow()

	return nil
}

type summaryContent struct {
	header string
	value  string
}

func (pdfGen *pdfGenerator) writeSummary(sumCon summaryContent) error {
	pdfGen.pdf.SetX((4.0 / 6.0) * EndX)
	err := pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: (1.0 / 6.0) * EndX, H: fontSizeContent * 1.3},
		sumCon.header,
		gopdf.CellOption{
			Align: Right + Middle,
		},
	)
	if err != nil {
		return err
	}

	pdfGen.pdf.SetX((4.5 / 6.0) * EndX)
	return pdfGen.pdf.CellWithOption(
		&gopdf.Rect{W: (1.0 / 6.0) * EndX, H: fontSizeContent * 1.3},
		sumCon.value,
		gopdf.CellOption{
			Align: Right + Middle,
		},
	)
}

/*
Pdf summary header section
*/
func (pdfGen *pdfGenerator) summary(sum Summary) error {
	pdfGen.pdf.SetY(summarySection)

	// write summary header
	content := []summaryContent{
		{
			header: "มูลค่าก่อนรวมภาษี",
			value:  fmt.Sprintf("%v", sum.PreVat),
		},
		{
			header: "ภาษี",
			value:  fmt.Sprintf("%v", sum.Vat),
		},
	}

	for _, v := range content {
		pdfGen.writeSummary(v)
		pdfGen.pdf.Br(fontSizeContent * 1.3)
	}

	pdfGen.pdf.SetFont(fontBold, "", fontSizeContent)
	grandTotal := summaryContent{
		header: "ยอดรวม",
		value:  fmt.Sprintf("%v", sum.GrandTotal),
	}
	pdfGen.writeSummary(grandTotal)

	return nil
}

/*
Generate whole pdf
*/
func (pdfGen *pdfGenerator) Generate(input Input) error {
	pdfGen.pdf.SetMarginLeft(pageMargin)
	if pdfGen.pdf == nil {
		return fmt.Errorf("pdf generator is empty")
	}

	if err := pdfGen.headerAndFooter(input.Document, input.Seller, input.Customer); err != nil {
		return err
	}
	if err := pdfGen.table(input.Items); err != nil {
		return err
	}
	if err := pdfGen.summary(input.Summary); err != nil {
		return err
	}

	return nil
}

/*
Write pdf to file
*/
func (pdfGen *pdfGenerator) Write(path string) error {
	return pdfGen.pdf.WritePdf(path)
}

/*
Clear pdf buffer
*/
func (pdfGen *pdfGenerator) Close() {
	pdfGen.pdf.Close()
	pdfGen.pdf = nil
}
