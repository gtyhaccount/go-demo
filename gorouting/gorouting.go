package main

import (
	"context"
	amq "github.com/rabbitmq/amqp091-go"
	"github.com/streadway/amqp"
	"log"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex

const EventExchange = "event_exchange"
const ConsumerTopic = "event_article"
const url = "amqp://admin:admin@rabbitmq:5672/"

type amqpClient struct {
	addr         string
	conn         *amq.Connection
	pubChannel   *amq.Channel
	subChannel   *amq.Channel
	messageTTL   int64 //单位毫秒
	counter      int
	exchangeName string
}

func main() {
	var conn *amq.Connection
	var err error

	conn, err = amq.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %s", err)
	} else {
		log.Println("Connection Connected")
	}
	channel, err := conn.Channel()
	if err != nil {
		return
	}

	channel2, err := conn.Channel()
	if err != nil {
		return
	}
	client := &amqpClient{conn: conn, pubChannel: channel, subChannel: channel2}
	// Notify on connection close
	//quit := make(chan struct{})

	go func() {
		client.ReceiveEventMessage(ConsumerTopic, ConsumerTopic, OnArticleActivityHandler)
	}()

	// Keep the application running
	for {
		select {
		case <-time.After(10 * time.Second):
			if err != nil {
				log.Println("连接channel失败")
				continue
			}

			client.SendData("{}", "event_article")

		}
	}

}

func (mq *amqpClient) SendData(Data string, routingKey string) {
	err := mq.pubChannel.Publish(
		EventExchange, // exchange
		routingKey,    // routing key
		false,         // mandatory
		false,
		amq.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(Data),
		})

	if err != nil {
		log.Printf("Sent Error %s", err)
		return
	}
	log.Printf("Sent %s", Data)
}

func (mq *amqpClient) ReceiveEventMessage(queue string, routingKey string, f func(ctx context.Context, delivery amq.Delivery) error) {

	ch := mq.subChannel
	message, err := ConsumeMessage(ch, queue, routingKey)
	if err != nil {
		return
	}

	go func() {
		<-ch.NotifyClose(make(chan *amq.Error))
		log.Println("connection was closed, attempting to reconnect...")
		for i := 0; i < 10; i++ {
			conns, err := amq.Dial(url)
			log.Printf("Reconnected to RabbitMg %d", i+1)
			if err == nil && conns != nil {
				mq.conn = conns
				mq.subChannel, _ = conns.Channel()
				mq.pubChannel, _ = conns.Channel()
				break
			}
			time.Sleep(10 * time.Second)
		}
		message, err = ConsumeMessage(mq.subChannel, queue, routingKey)
		if err != nil {
			log.Printf("close cancel %s", err)
		} else {
			log.Printf("open channel %s", err)
		}
		//quit <- struct{}{}

	}()

	forever := make(chan bool)
	go func() {
		for {
			for d := range message {
				if !strings.Contains(d.RoutingKey, routingKey) {
					continue
				}
				err = f(context.Background(), d)
				if err != nil {
					log.Printf(" [*] consume errror %s", err)
				}
			}
			time.Sleep(time.Second * 5)
			log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
			log.Printf("Closing consumer...")
		}
	}()

	<-forever

}

func ConsumeMessage(channel *amq.Channel, queue, routingKey string) (<-chan amq.Delivery, error) {
	q, err := channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	err = channel.QueueBind(
		q.Name,        // queue name
		routingKey,    // routing key
		EventExchange, // exchange
		false,
		nil,
	)

	msgs, err := channel.Consume(
		q.Name,        // queue
		ConsumerTopic, // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)

	return msgs, err
}

func OnArticleActivityHandler(ctx context.Context, mq amq.Delivery) error {

	log.Printf("consumer ok routingkey %s body %s", mq.RoutingKey, string(mq.Body))
	return nil
}
