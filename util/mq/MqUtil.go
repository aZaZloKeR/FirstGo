package mq

import (
	"context"
	"log"

	rabbit "github.com/rabbitmq/amqp091-go"
)

func SendMess(body string, qname string, ctx context.Context) {
	conn := createConnection()
	defer conn.Close()

	q, ch := createQueue(qname, conn)

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

func ReadMess(qname string, c chan string) {
	conn := createConnection()
	defer conn.Close()

	q, ch := createQueue(qname, conn)

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
		c <- string(d.Body)
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

func createConnection() *rabbit.Connection {
	conn, err := rabbit.Dial("amqp://admin:admin@amqp.rt.sm-soft.ru:5672/")
	failOnError(err, "unable to open connect to RabbitMQ server. Error: %s")

	return conn
}

func createQueue(qname string, conn *rabbit.Connection) (rabbit.Queue, *rabbit.Channel) {
	ch, err := conn.Channel()

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
	return q, ch
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
