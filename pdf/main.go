package main

import (
	"io/ioutil"
	"log"

	generator "github.com/angelodlfrtr/go-invoice-generator"
)
// type Tax struct{
// 	Percent string
// 	Amount string
// }
// type Discount struct{
// 	Percent string
// 	Amount string
// }
func main() {
	doc, _ := generator.New(generator.Invoice, &generator.Options{
		TextTypeInvoice: "Invoice",
		AutoPrint:       true,
	})

	doc.SetHeader(&generator.HeaderFooter{
		Text:       "<center>Chantosweb developers.</center>",
		Pagination: true,
	})

	doc.SetFooter(&generator.HeaderFooter{
		Text:       "<center>Chantosweb developers.</center>",
		Pagination: true,
	})

	doc.SetRef("Chantosweb ")
	doc.SetVersion("version 2.0")

	doc.SetDescription("Items")
	doc.SetNotes("terms and conditions")

	doc.SetDate("02/03/2021")
	doc.SetPaymentTerm("02/04/2021")

	logoBytes, _ := ioutil.ReadFile("./hopitality.jpeg")

	doc.SetCompany(&generator.Contact{
		Name: "Chantosweb developers",
		Logo: &logoBytes,
		Address: &generator.Address{
			Address:    "123456 kd",
			Address2:   "Appartement 2",
			PostalCode: "12345",
			City:       "Nairobi",
			Country:    "Kenya",
		},
	})

	doc.SetCustomer(&generator.Contact{
		Name: "Tikas Electronics",
		Address: &generator.Address{
			Address:    "456 de kau",
			PostalCode: "29200",
			City:       "Kiambu",
			Country:    "Kenya",
		},
	})
  doc.Options.CurrencySymbol = "$ "
	for i := 0; i < 3; i++ {
		doc.AppendItem(&generator.Item{
			Name:        "Mac book",
			Description: "Apple mac book pro",
			UnitCost:    "2000",
			Quantity:    "2",
			// Tax: &Tax{
			// 	Percent: "20",
			// },
		})
	}

	doc.AppendItem(&generator.Item{
		Name:     "Hp laptop",
		UnitCost: "500",
		Quantity: "5",
		// Tax: &Tax{
		// 	Amount: "89",
		// },
		// Discount: &Discount{
		// 	Percent: "30",
		// },
	})

	doc.AppendItem(&generator.Item{
		Name:     "Acer d5",
		UnitCost: "700.89",
		Quantity: "2",
		// Discount: &Discount{
		// 	Percent: "50",
		// },
	})

	doc.AppendItem(&generator.Item{
		Name:     "Samsung",
		UnitCost: "889.89",
		Quantity: "7",
		// Discount: &Discount{
		// 	Amount: "234.67",
		// },
	})

	doc.SetDefaultTax(&generator.Tax{
		Percent: "10",
	})

	// doc.SetDiscount(&generator.Discount{
	// Percent: "90",
	// })
	doc.SetDiscount(&generator.Discount{
		Amount: "1340",
	})

	pdf, err := doc.Build()
	if err != nil {
		log.Fatal(err)
	}

	err = pdf.OutputFileAndClose("out.pdf")

	if err != nil {
		log.Fatal(err)
	}
}