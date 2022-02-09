package infrastructures

import (
	"os"

	"github.com/streadway/amqp"
)

//PublishMessage -> send message to broker
func PublishMessage(exchangeName string, message []byte) error {
	conn, err := amqp.Dial(os.Getenv("BROKER_CONNECTION"))
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.Publish(exchangeName, "", false, false, amqp.Publishing{
		AppId:       "go.customers",
		ContentType: "application/json",
		Body:        message,
	})

	if err != nil {
		return err
	}

	return nil
}
