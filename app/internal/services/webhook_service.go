package services

import (
	"encoding/json"
	"fmt"

	"github.com/starkbank/sdk-go/starkbank/event"
	"github.com/starkinfra/core-go/starkcore/user/project"
)

// ProcessWebhookEvent processa o evento recebido no webhook
func ProcessWebhookEvent(body []byte, signature string) error {
	// Configuração do usuário do StarkBank
	var privateKeyContent = `-----BEGIN EC PRIVATE KEY-----
	MHQCAQEEIN0NFH1lGEzLXhnaXxKKBqC3J1WWuLtiRAzSEfRXBqTgoAcGBSuBBAAK
	oUQDQgAEu4gONKh9t794DaLahDib/rfL5aGyR0V/0RSvZ6cd46y/j78ybFWsd04Y
	kiDAFLMFGeLuP0u4n2JV1JIPyBSL6w==
	-----END EC PRIVATE KEY-----`

	var user = &project.Project{
		Id:          "6250122287513600",
		PrivateKey:  privateKeyContent,
		Environment: "sandbox", // ou "production"
	}
	// Faz o parse do evento
	parsedEvent := event.Parse(string(body), signature, user)

	// Faça uma verificação de tipo no evento
	eventMap, ok := parsedEvent.(map[string]interface{})
	if !ok {
		return fmt.Errorf("erro: evento não está no formato esperado")
	}

	// Verifica se o evento é de transferência
	subscription, ok := eventMap["subscription"].(string)
	if !ok || subscription != "transfer" {
		return fmt.Errorf("evento não é do tipo 'transfer'")
	}

	// Extrai o payload do evento
	payload, ok := eventMap["payload"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("erro: payload não está no formato esperado")
	}

	// Decodifica os dados da transferência
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("erro ao converter payload para JSON: %v", err)
	}

	var transferData map[string]interface{}
	err = json.Unmarshal(payloadJSON, &transferData)
	if err != nil {
		return fmt.Errorf("erro ao decodificar o payload: %v", err)
	}

	// Processa a transferência (adicione sua lógica aqui)
	fmt.Printf("Transferência recebida: %+v\n", transferData)

	return nil
}
