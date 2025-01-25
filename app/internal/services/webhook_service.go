package services

import (
	"encoding/json"
	"fmt"

	"github.com/darcioSoares/stark/internal/config"
	"github.com/starkbank/sdk-go/starkbank/event"
	"github.com/starkinfra/core-go/starkcore/user/project"
)


// ProcessWebhookEvent processa o evento recebido no webhook
func ProcessWebhookEvent(body []byte, signature string) error {

	user := &project.Project{
		Id:          config.IDProject,
		PrivateKey:  config.PrivateKey,
		Environment: "sandbox",
	}

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

	// Extrai o payload 
	payload, ok := eventMap["payload"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("erro: payload não está no formato esperado")
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("erro ao converter payload para JSON: %v", err)
	}

	var transferData map[string]interface{}
	err = json.Unmarshal(payloadJSON, &transferData)
	if err != nil {
		return fmt.Errorf("erro ao decodificar o payload: %v", err)
	}

	fmt.Printf("Transferência recebida: %+v\n", transferData)

	return nil
}
