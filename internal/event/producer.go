package event

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ProducerInterface interface {
	Publish(msg []byte) error
}

type Producer struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	queueName string
}

func NewProducer(connStr string, queueName string) (*Producer, error) {
	// conexão com o RabbitMQ.
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}

	// cria a fila de produção -- canal da conexão.
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queueName,
		true, // mantém mesmo ao reiniciar
		false, false, false, nil,
	)
	if err != nil {
		return nil, err
	}

	return &Producer{conn: conn, ch: ch, queueName: queueName}, nil
}

func (p *Producer) Publish(msg []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.ch.PublishWithContext(ctx,
		"",
		p.queueName,
		false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)
}

func (p *Producer) Close() {
	p.ch.Close()
	p.conn.Close()
}
