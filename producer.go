// produce msgs

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {

	// Connect to local RabbitMQ instance
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Connected to RabbitMQ instance")

	ch, err := conn.Channel()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer ch.Close()

	// Declare new queue to hold msgs and deliver to consumers
	_, err = ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Publish msg to queue (via exchange) every 2 seconds
	for {
		time.Sleep(time.Second * 2)

		err = ch.Publish(
			"",
			"TestQueue",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("hello world"),
			},
		)

		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		log.Println("Published msg to queue")
	}
}
