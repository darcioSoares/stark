package services

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
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

	numRequests := rand.Intn(3) + 4
	//numRequests := rand.Intn(5) + 8
	fmt.Printf("Enviando %d requisições...\n", numRequests)

	for i := 0; i < numRequests; i++ {

		body := []byte(fmt.Sprintf(`{"message": "requisição de invoice número %d"}`, i+1))

		resp, err := http.Post("http://localhost:3000/dss", "application/json", bytes.NewBuffer(body))
		if err != nil {
			fmt.Printf("Erro ao enviar requisição %d: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("Requisição %d enviada com sucesso, status: %d\n", i+1, resp.StatusCode)
	}
}
