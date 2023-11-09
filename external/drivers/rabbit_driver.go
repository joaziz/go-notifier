package drivers

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMQDrive struct {
	conn   *amqp.Connection
	ch     *amqp.Channel
	queues map[string]amqp.Queue
}

func (c *RabbitMQDrive) Listen(queue string, f func(err error, msg string, ack func())) {

	ch, err := getConsumer(c, queue)
	if err != nil {
		f(err, "", func() {})
		return
	}
	fmt.Println("dddd")
	for {
		select {
		case message := <-ch:
			log.Printf("Received a message: %s", message.Exchange)
			f(nil, string(message.Body), func() { message.Ack(false) })
		}
	}
}

func (c *RabbitMQDrive) Send(ctx context.Context, queue string, payload string) error {
	if _, ok := c.queues[queue]; !ok {
		err := c.ch.ExchangeDeclare(
			queue,    // name
			"fanout", // type
			true,     // durable
			false,    // auto-deleted
			false,    // internal
			false,    // no-wait
			nil,      // arguments
		)
		if err != nil {
			return err
		}
	}

	return c.ch.PublishWithContext(ctx,
		queue, // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})
}

func (c *RabbitMQDrive) Close() {
	defer c.conn.Close()
	defer c.ch.Close()
}

func NewRabbitMQDrive(url string) *RabbitMQDrive {
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return &RabbitMQDrive{
		conn: conn,
		ch:   ch,
	}
}

func getConsumer(c *RabbitMQDrive, queue string) (<-chan amqp.Delivery, error) {
	if _, ok := c.queues[queue]; !ok {

		err := c.ch.ExchangeDeclare(
			queue,    // name
			"fanout", // type
			true,     // durable
			false,    // auto-deleted
			false,    // internal
			false,    // no-wait
			nil,      // arguments
		)
		if err != nil {
			return nil, err
		}
	}

	q, err := c.ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return nil, err
	}
	err = c.ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		queue,  // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return c.ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}
