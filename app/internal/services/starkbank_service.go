package services

import (
	"fmt"
	"time"

	"github.com/darcioSoares/stark/internal/utils"

	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/invoice"
	"github.com/starkbank/sdk-go/starkbank/transfer"
	"github.com/starkinfra/core-go/starkcore/user/project"
)

var privateKeyContent = `-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIN0NFH1lGEzLXhnaXxKKBqC3J1WWuLtiRAzSEfRXBqTgoAcGBSuBBAAK
oUQDQgAEu4gONKh9t794DaLahDib/rfL5aGyR0V/0RSvZ6cd46y/j78ybFWsd04Y
kiDAFLMFGeLuP0u4n2JV1JIPyBSL6w==
-----END EC PRIVATE KEY-----`

var user = &project.Project{
	Id:          "6250122287513600",
	PrivateKey:  privateKeyContent,
	Environment: "sandbox",
}

func CreateInvoice() ([]invoice.Invoice, error) {

	starkbank.User = user

	// Obtenha dados aleatórios de contas para emitir invoices
	randomInvoice := utils.GetRandomInvoice()

	invoices, err := invoice.Create([]invoice.Invoice{
		{
			Amount: randomInvoice.Amount,
			Name:   randomInvoice.Name,
			TaxId:  randomInvoice.TaxId,
		},
	}, nil)

	if err.Errors != nil {
		return nil, fmt.Errorf("erro ao criar invoice: %v", err.Errors)
	}
	return invoices, nil
}

func CreateTransfer(amount int, name string) ([]transfer.Transfer, error) {
	// privateKeyContent := os.Getenv("PRIVATE_KEY")
	// idProject := os.Getenv("ID_PROJECT")

	// privateKey := os.Getenv("PRIVATE_KEY")
	// if privateKey == "" {
	// 	log.Fatalf("PRIVATE_KEY não configurado no arquivo .env")
	// }

	user := &project.Project{
		Id:          "6250122287513600",
		PrivateKey:  privateKeyContent,
		Environment: "sandbox",
	}

	starkbank.User = user

	scheduled := time.Date(2025, 01, 30, 0, 0, 0, 0, time.UTC)

	transfers, err := transfer.Create([]transfer.Transfer{
		{
			Amount:        amount,
			Name:          "Stark Bank S.A.",
			TaxId:         "20.018.183/0001-80",
			BankCode:      "20018183",
			BranchCode:    "0001",
			AccountNumber: "6341320293482496",
			ExternalId:    "my-external-id",
			Scheduled:     &scheduled,
			Tags:          []string{name, "invoice"},
		},
	}, nil)

	if err.Errors != nil {
		return nil, fmt.Errorf("erro ao criar transferência: %v", err.Errors)
	}

	return transfers, nil
}
