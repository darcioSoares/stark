package services

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQService struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func (r *RabbitMQService) Initialize(amqpURL, queueName, exchangeName, exchangeType string) error {
	var err error

	// Tentativas de conexão com retry
	for i := 1; i <= 20; i++ {
		r.conn, err = amqp.Dial(amqpURL)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to RabbitMQ (attempt %d): %v", i, err)
		time.Sleep(4 * time.Second)
	}

	if err != nil {
		return err
	}

	// Criar o canal
	r.channel, err = r.conn.Channel()
	if err != nil {
		return err
	}

	// Declarar a exchange
	err = r.channel.ExchangeDeclare(
		exchangeName, // Nome da exchange
		exchangeType, // Tipo da exchange (direct, fanout, topic, headers)
		true,         // Durable
		false,        // Auto-delete
		false,        // Interna
		false,        // No-wait
		nil,          // Argumentos
	)
	if err != nil {
		return err
	}

	r.queue, err = r.channel.QueueDeclare(
		queueName, // Nome da fila
		true,      // Durable
		false,     // Auto-delete
		false,     // Exclusiva
		false,     // No-wait
		nil,       // Argumentos
	)
	if err != nil {
		return err
	}

	// fila com exchange
	err = r.channel.QueueBind(
		queueName,    // Nome da fila
		"",           // Routing key (use "" para fanout)
		exchangeName, // Nome da exchange
		false,        // No-wait
		nil,          // Argumentos
	)
	if err != nil {
		return err
	}

	log.Printf("Connected to RabbitMQ: queue='%s', exchange='%s'", queueName, exchangeName)
	return nil
}

func (r *RabbitMQService) SendMessage(exchangeName, routingKey, message string) error {
	err := r.channel.Publish(
		exchangeName, // Nome da exchange
		routingKey,   // Routing key (use "" para fanout)
		false,        // Mandatory
		false,        // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	log.Printf("Message published to exchange '%s': %s", exchangeName, message)
	return nil
}

func (r *RabbitMQService) ConsumeMessages() (<-chan amqp.Delivery, error) {
	messages, err := r.channel.Consume(
		r.queue.Name, // Nome da fila
		"",           // Consumer tag
		true,         // Auto-acknowledge
		false,        // Exclusiva
		false,        // No-local
		false,        // No-wait
		nil,          // Argumentos
	)
	if err != nil {
		log.Printf("Failed to consume messages: %v", err)
		return nil, err
	}

	log.Println("Started consuming messages")
	return messages, nil
}

func (r *RabbitMQService) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
	log.Println("RabbitMQ connection and channel closed")
}
