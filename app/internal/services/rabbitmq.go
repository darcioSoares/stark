package services

import (
	"log"

	"github.com/streadway/amqp"
)

//estrutura para gerenciar o RabbitMQ
type RabbitMQService struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

// NewRabbitMQService cria e inicializa um novo RabbitMQService
func NewRabbitMQService(queueName string, rabbitURL string) (*RabbitMQService, error) {
	// Conectar ao RabbitMQ
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, err
	}

	// Criar um canal
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	// Declarar a fila
	_, err = channel.QueueDeclare(
		queueName, // Nome da fila
		true,      // Durável
		false,     // Auto-delete
		false,     // Exclusiva
		false,     // Sem espera
		nil,       // Argumentos
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, err
	}

	// Retorna a instância do RabbitMQService
	return &RabbitMQService{
		conn:    conn,
		channel: channel,
		queue:   queueName,
	}, nil
}

// Publish envia uma mensagem para a fila
func (r *RabbitMQService) Publish(body []byte) error {
	return r.channel.Publish(
		"",        // Exchange
		r.queue,   // Routing key
		false,     // Mandar de volta se não tiver fila
		false,     // Mandar de volta se não for consumido
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// ConsumeMessages consome mensagens da fila e envia para um canal Go
func (r *RabbitMQService) ConsumeMessages() (<-chan amqp.Delivery, error) {
	messages, err := r.channel.Consume(
		r.queue, // Nome da fila
		"",      // Nome do consumidor
		true,    // Auto-acknowledge
		false,   // Exclusivo
		false,   // Sem espera
		false,   // Argumentos
		nil,
	)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// Close fecha a conexão e o canal do RabbitMQ
func (r *RabbitMQService) Close() {
	if err := r.channel.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ channel: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ connection: %v", err)
	}
}
