package event

import (
	"log/slog"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewConsumer(connStr string) (*Consumer, error) {
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{conn: conn, ch: ch}, nil
}

// chamar apenas com Goroutine.
func (c *Consumer) Start(queueName string) error {
	_, err := c.ch.QueueDeclare(
		queueName,
		true,
		false, false, false, nil,
	)
	if err != nil {
		return err
	}

	msgs, err := c.ch.Consume(
		queueName,
		"go-wallet-consumer",
		true,
		false, false, false, nil,
	)
	if err != nil {
		return err
	}

	slog.Info("Waiting for messages on queue", "queue", queueName)
	for msg := range msgs {
		slog.Info("Notification sent", "payload", string(msg.Body))
	}

	return nil
}

func (c *Consumer) Close() {
	c.ch.Close()
	c.conn.Close()
}
