package utils

import (
	"math/rand"
	"time"

	"github.com/starkbank/sdk-go/starkbank/invoice"
)

func GetRandomInvoice() invoice.Invoice {

	invoiceData := []invoice.Invoice{
		{Amount: 100050, Name: "Jon Snow", TaxId: "165.044.680-28"},
		{Amount: 200000, Name: "Arya Stark", TaxId: "052.338.810-12"},
		{Amount: 300000, Name: "Sansa Stark", TaxId: "673.044.700-11"},
		{Amount: 450000, Name: "Bran Stark", TaxId: "520.293.600-15"},
		{Amount: 700000, Name: "Daenerys Targaryen", TaxId: "722.905.050-21"},
		{Amount: 710000, Name: "Tyrion Lannister", TaxId: "265.692.570-30"},
		{Amount: 550000, Name: "Cersei Lannister", TaxId: "670.835.800-06"},
		{Amount: 800000, Name: "Jaime Lannister", TaxId: "676.056.690-46"},
	}

	// Gera um índice aleatório
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(invoiceData))

	return invoiceData[randomIndex]
}
