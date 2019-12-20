package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn,err := amqp.Dial("amqp://admin:admin@127.0.0.1:5672")
	failOnError(err,"Failed to connect to RabbitMQ")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
