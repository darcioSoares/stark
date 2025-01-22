package services

import (
	"fmt"
	"time"

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

func CreateInvoice() ([]invoice.Invoice, error) {
	
	user := &project.Project{
		Id:          "6250122287513600",
		PrivateKey:  privateKeyContent,
		Environment: "sandbox", 
	}

	
	starkbank.User = user

	// Criação de invoices
	invoices, err := invoice.Create([]invoice.Invoice{
		{
			Amount: 400050, 
			Name:   "Jow Snow B Stark",
			TaxId:  "983.418.590-17", // CPF ou CNPJ
		},
	}, nil)

	if err.Errors != nil {
		return nil, fmt.Errorf("erro ao criar invoice: %v", err.Errors)
	}
	return invoices, nil
}

func CreateTransfer() ([]transfer.Transfer, error) {
	// Inicialize o usuário do projeto
	user := &project.Project{
		Id:          "6250122287513600", 
		PrivateKey:  privateKeyContent,
		Environment: "sandbox", 
	}

	
	starkbank.User = user
	
	scheduled := time.Date(2025, 01, 30, 0, 0, 0, 0, time.UTC)

	
	transfers, err := transfer.Create([]transfer.Transfer{
		{
			Amount:        1000000,
			Name:          "Stark Bank S.A.",
			TaxId:         "20.018.183/0001-80",                 
			BankCode:      "20018183",                           
			BranchCode:    "0001",                               
			AccountNumber: "6341320293482496",                   
			ExternalId:    "my-external-id",                    
			Scheduled:     &scheduled,                           
			Tags:          []string{"daenerys", "invoice/1234"}, 
			
		},
	}, nil)

	if err.Errors != nil {
		return nil, fmt.Errorf("erro ao criar transferência: %v", err.Errors)
	}

	return transfers, nil
}
