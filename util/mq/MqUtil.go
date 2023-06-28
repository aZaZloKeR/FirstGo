package mq

import (
	"awesomeProject/util"
	"awesomeProject/util/syncer"
	"context"
	"log"

	rabbit "github.com/rabbitmq/amqp091-go"
)

func SendMess(ctx context.Context, body string, qname string) {
	syncer.SetAlive()
	defer syncer.SetCompleted()
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

func ReadMess(ctx context.Context, qname string) chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		syncer.SetAlive()
		defer syncer.SetCompleted()
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
		for true {
			select {
			case d := <-msgs:
				log.Printf("Received a message: %s", d.Body)
				util.SendWithoutBlock(string(d.Body), c)
			case <-ctx.Done():
				return
			}
		}
	}()
	return c
}

func createConnection() *rabbit.Connection {
	//DON'T DO IT ((
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
