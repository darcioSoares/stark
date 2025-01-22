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
		time.Sleep(8 * time.Minute)
	}
}

func sendRequests() {
	// Define número aleatório de requisições entre 4 e 6
	numRequests := rand.Intn(2) + 4

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

// func sendRequests() {

// 	numRequests := rand.Intn(3) + 4
// 	//numRequests := rand.Intn(5) + 8
// 	fmt.Printf("Enviando %d requisições...\n", numRequests)

// 	for i := 0; i < numRequests; i++ {

// 		body := []byte(fmt.Sprintf(`{"message": "requisição de invoice número %d"}`, i+1))

// 		resp, err := http.Post("http://localhost:3000/dss", "application/json", bytes.NewBuffer(body))
// 		if err != nil {
// 			fmt.Printf("Erro ao enviar requisição %d: %v\n", i+1, err)
// 			continue
// 		}
// 		defer resp.Body.Close()

// 		fmt.Printf("Requisição %d enviada com sucesso, status: %d\n", i+1, resp.StatusCode)
// 	}
// }
