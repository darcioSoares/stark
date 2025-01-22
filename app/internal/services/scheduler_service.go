package services

import (
	"fmt"
	"math/rand"
	"time"
)

func SendRequestsEveryHour() {
	for {

		sendRequests()

		fmt.Println("Aguardando 1 hora para envio...")
		//time.Sleep(1 * time.Hour)
		time.Sleep(20 * time.Minute)
	}
}

func sendRequests() {
	// Define número aleatório de requisições entre 4 e 6
	numRequests := rand.Intn(2) + 3

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


