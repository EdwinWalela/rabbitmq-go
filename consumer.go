// consume msgs from queue

package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	// Consume messages in TestQueue
	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Recieved msg: %s\n", d.Body)
		}
	}()

	fmt.Println("Connected to RabbitMQ instance")
	fmt.Println(" [*] - waiting for messages\n")

	// Block program from exiting
	<-forever

}
