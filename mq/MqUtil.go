package mq

import (
	"context"
	rabbit "github.com/rabbitmq/amqp091-go"
	"log"
)

var q rabbit.Queue
var ch *rabbit.Channel

func CreateConnection(ch chan int) {
	conn, err := rabbit.Dial("amqp://admin:admin@amqp.rt.sm-soft.ru:5672/") // Создаем подключение к RabbitMQ
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	createQueue(conn)
	ch <- 1
	close(ch)
}

func createQueue(conn *rabbit.Connection) {
	var err error
	ch, err = conn.Channel()
	q, err = ch.QueueDeclare(
		"TEST.GO",
		true,
		false,
		false,
		false,
		rabbit.Table{"x-dead-letter-exchange": "",
			"x-dead-letter-routing-key": "TEST.GO.DLQ"}, // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}
}

func SendMess(body string, ctx *context.Context) {
	err := ch.PublishWithContext(*ctx,
		"",
		q.Name, // routing key
		false,
		false,
		rabbit.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("failed to publish a message. Error: %s", err)
	}

	log.Printf(" [x] Sent %s\n", body)
}
