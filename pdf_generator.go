package documentgenerator

type DocumentInformation struct {
	Id        string `json:"id"`
	Header    string `json:"header"`
	Code      string `json:"code"`
	IssueDate string `json:"issue_date"`
}

type PersonalInformation struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	TaxId   string `json:"tax_id"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type Item struct {
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Amount    float64 `json:"amount"`
}

type Summary struct {
	PreVat     float64 `json:"pre_vat"`
	Vat        float64 `json:"vat"`
	GrandTotal float64 `json:"grand_total"`
}

type Input struct {
	Document DocumentInformation `json:"document"`
	Seller   PersonalInformation `json:"seller"`
	Customer PersonalInformation `json:"customer"`
	Items    []Item              `json:"items"`
	Summary  Summary             `json:"summary"`
}

type PdfGenerator interface {
	// set font
	SetFont(path string) error
	SetFontBold(path string) error
	SetFontPlain(path string) error
	// set font size
	SetFontSize(size float64)
	SetFontHeaderSize(size float64)
	SetFontContentSize(size float64)
	SetFontSubHeaderSize(size float64)

	Generate(input Input) error
	Write(path string) error
	Close()
}

func NewPdfGenerator() PdfGenerator {
	return newGoPdf()
}
