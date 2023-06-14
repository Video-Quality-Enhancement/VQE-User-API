package producers

import (
	"context"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type WelcomeProducer interface {
	Publish(userId string) error
}

type welcomeProducer struct {
	queue string
	conn  config.AMQPconnection
}

func NewWelcomeProducer(conn config.AMQPconnection) WelcomeProducer {
	queue := "welcome_queue"
	return &welcomeProducer{
		conn:  conn,
		queue: queue,
	}
}

func (producer *welcomeProducer) Publish(userId string) error {

	ch, err := producer.conn.NewChannel()
	if err != nil {
		slog.Error("Failed to open a channel", "error", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		producer.queue, // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		slog.Error("Failed to declare a queue", "error", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := []byte(userId)

	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		slog.Error("Failed to publish a message", "error", err)
		return err
	}

	slog.Debug("Message Published", "body", body)
	return nil

}
