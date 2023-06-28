package mq

import (
	"context"
	rabbit "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func SendMess(body string, qname string) {
	conn := createConnection()
	defer func() {
		_ = conn.Close()
	}()

	q, ch, ctx, cancel := createQueue(qname, conn)
	defer cancel()

	err := ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		rabbit.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "failed to publish a message. Error: %s")

	log.Printf("Sent %s\n", body)
}

func ReadMess(qname string) string {
	conn := createConnection()
	defer func() {
		_ = conn.Close()
	}()

	q, ch, _, cancel := createQueue(qname, conn)
	defer cancel()

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		rabbit.Table{"x-dead-letter-exchange": "",
			"x-dead-letter-routing-key": "TEST.GO.DLQ"}, // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		return string(d.Body)
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	return ""
}

func createConnection() *rabbit.Connection {
	conn, err := rabbit.Dial("amqp://admin:admin@amqp.rt.sm-soft.ru:5672/")
	failOnError(err, "unable to open connect to RabbitMQ server. Error: %s")

	return conn
}

func createQueue(qname string, conn *rabbit.Connection) (rabbit.Queue, *rabbit.Channel, context.Context, context.CancelFunc) {
	ch, err := conn.Channel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	q, err := ch.QueueDeclare(
		qname,
		true,
		false,
		false,
		false,
		rabbit.Table{"x-dead-letter-exchange": "",
			"x-dead-letter-routing-key": "TEST.GO.DLQ"},
	)
	failOnError(err, "failed to declare a queue. Error: %s")
	return q, ch, ctx, cancel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
