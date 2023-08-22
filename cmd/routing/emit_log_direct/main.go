package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/kaiquecaires/go-rabbitmq/cmd/helpers"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helpers.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FailOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := helpers.BodyFrom(os.Args)
	err = ch.PublishWithContext(ctx,
		"logs_direct",
		helpers.SeverityFrom(os.Args),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	helpers.FailOnError(err, "Fail on sending message")

	log.Printf(" [x] Sent %s", body)
}
