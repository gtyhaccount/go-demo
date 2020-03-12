package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func SendMessage(message string) {
	conn, err := amqp.Dial("amqp://admin:C4ATYHYS@mdm-allinone-rmq.ebz-chn-dev.mkaws.com:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Fail to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		fmt.Printf("%s: %s", msg, err)
	}
}
