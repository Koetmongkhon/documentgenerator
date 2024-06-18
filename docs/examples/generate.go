package generate_examples

import (
	"fmt"

	"github.com/wasawat.ko/documentgenerator"
)

func generate(input documentgenerator.Input, outputPath string) {
	// create pdf generator instant
	pdfGen := documentgenerator.NewPdfGenerator()
	defer pdfGen.Close()

	// set font
	err := pdfGen.SetFontBold("./fonts/THSarabunNew Bold.ttf")
	if err != nil {
		fmt.Printf("Set header font error: %s", err)
	}
	err = pdfGen.SetFontPlain("./fonts/THSarabunNew.ttf")
	if err != nil {
		fmt.Printf("Set header font error: %s", err)
	}

	err = pdfGen.Generate(input)
	if err != nil {
		fmt.Printf("Generate pdf error: %s", err)
	}

	err = pdfGen.Write(outputPath)
	if err != nil {
		fmt.Printf("Write to pdf error: %s", err)
	}

	fmt.Println("Generate done")
}

func Generate() {
	// input
	input := documentgenerator.Input{
		Document: documentgenerator.DocumentInformation{
			Id:        "INV1234567890",
			Header:    "ใบกำกับภาษี",
			Code:      "INV",
			IssueDate: "7/6/2024",
		},
		Seller: documentgenerator.PersonalInformation{
			Name:    "Company Name",
			Address: "1234 ถ.XXXX แขวงXXXX เขตXXXX จ.กรุงเทพมหานคร 00000",
			TaxId:   "1234567890",
			Phone:   "02-000-0000",
			Email:   "seller.contact@co.th",
		},
		Customer: documentgenerator.PersonalInformation{
			Name:    "ชื่อของ ผู้ซื้อ",
			Address: "9876 ถ.YYYY ต.YYYY อ.YYYY จ.นครราชสีมา 99999",
			TaxId:   "1234567890456",
			Phone:   "098-123-4567",
			Email:   "buyer.contact@hotmail.com",
		},
		Items: []documentgenerator.Item{
			{
				Name:      "A",
				Quantity:  1,
				UnitPrice: 100,
				Amount:    100,
			},
			{
				Name:      "B",
				Quantity:  2,
				UnitPrice: 1000,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
		},
		Summary: documentgenerator.Summary{
			PreVat:     2000,
			Vat:        140,
			GrandTotal: 2140,
		},
	}
	generate(input, "./generate.pdf")
}

func GenerateMultiplePages() {
	// input
	input := documentgenerator.Input{
		Document: documentgenerator.DocumentInformation{
			Id:        "INV1234567890",
			Header:    "ใบกำกับภาษี",
			Code:      "INV",
			IssueDate: "7/6/2024",
		},
		Seller: documentgenerator.PersonalInformation{
			Name:    "Company Name",
			Address: "1234 ถ.XXXX แขวงXXXX เขตXXXX จ.กรุงเทพมหานคร 00000",
			TaxId:   "1234567890",
			Phone:   "02-000-0000",
			Email:   "seller.contact@co.th",
		},
		Customer: documentgenerator.PersonalInformation{
			Name:    "ชื่อของ ผู้ซื้อ",
			Address: "9876 ถ.YYYY ต.YYYY อ.YYYY จ.นครราชสีมา 99999",
			TaxId:   "1234567890456",
			Phone:   "098-123-4567",
			Email:   "buyer.contact@hotmail.com",
		},
		Items: []documentgenerator.Item{
			{
				Name:      "A",
				Quantity:  1,
				UnitPrice: 100,
				Amount:    100,
			},
			{
				Name:      "A",
				Quantity:  1,
				UnitPrice: 100,
				Amount:    100,
			},
			{
				Name:      "A",
				Quantity:  1,
				UnitPrice: 100,
				Amount:    100,
			},
			{
				Name:      "A",
				Quantity:  1,
				UnitPrice: 100,
				Amount:    100,
			},
			{
				Name:      "A",
				Quantity:  1,
				UnitPrice: 100,
				Amount:    100,
			},
			{
				Name:      "A",
				Quantity:  1,
				UnitPrice: 100,
				Amount:    100,
			},
			{
				Name:      "B",
				Quantity:  2,
				UnitPrice: 1000,
			},
			{
				Name:      "B",
				Quantity:  2,
				UnitPrice: 1000,
			},
			{
				Name:      "B",
				Quantity:  2,
				UnitPrice: 1000,
			},
			{
				Name:      "B",
				Quantity:  2,
				UnitPrice: 1000,
			},
			{
				Name:      "B",
				Quantity:  2,
				UnitPrice: 1000,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
			{
				Name:      "Baksldjfklajskldncklasjdfklasjdfljaslkdfj askdjfalks;jdfklasdkvn",
				Quantity:  3,
				UnitPrice: 300,
			},
		},
		Summary: documentgenerator.Summary{
			PreVat:     2000,
			Vat:        140,
			GrandTotal: 2140,
		},
	}
	generate(input, "./generate_multiple_pages.pdf")
}
