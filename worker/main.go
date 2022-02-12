package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load .env file")
	}

	conn, err := amqp.Dial(os.Getenv("BROKER_CONNECTION"))
	if err != nil {
		log.Println("unexpected error ocurred while trying to open connection: ", err.Error())
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println("unexpected error ocurred while trying to open connection: ", err.Error())
		return
	}
	defer ch.Close()

	messages, err := ch.Consume("changes_customers", "changes_customers", true, false, false, false, nil)

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf("Received a message: %s", message.Body)
		}
	}()

	log.Printf(" [*] waiting for messages. To exit press CTRL+C")

	<-forever
}
