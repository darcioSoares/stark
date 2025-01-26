package services

import (
	"fmt"
	"time"

	"github.com/darcioSoares/stark/internal/utils"

	"github.com/darcioSoares/stark/internal/config"
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/invoice"
	"github.com/starkbank/sdk-go/starkbank/transfer"
	"github.com/starkinfra/core-go/starkcore/user/project"
)

func CreateInvoice() ([]invoice.Invoice, error) {

	user := &project.Project{
		Id:          config.IDProject,
		PrivateKey:  config.PrivateKey,
		Environment: "sandbox",
	}

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

	user := &project.Project{
		Id:          config.IDProject,
		PrivateKey:  config.PrivateKey,
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

func ConsumeMessagesFila() {
	
}