package services

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func SendRequestsEveryHour() {
	for {

		sendRequests()

		fmt.Println("Aguardando 1 hora para envio...")
		time.Sleep(3 * time.Hour)
		//time.Sleep(10 * time.Minute)
	}
}

func sendRequests() {

	numRequests := rand.Intn(14-8+1) + 8

	fmt.Printf("Enviando %d requisições...\n", numRequests)

	for i := 0; i < numRequests; i++ {

		invoices, err := CreateInvoice()
		if err != nil {
			fmt.Printf("Erro ao criar invoice %d: %v\n", i+1, err)
			continue
		}

		for _, inv := range invoices {
			fmt.Printf("Invoice criada na requisição %d: %+v\n", i+1, inv)
		}
	}
}
