package broker

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"time"
)

var (
	Ctx        context.Context
	BrokerConn amqp.Connection
)

// TODO:
func Conenct() {
	Ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	BrokerConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", Ctx)
	fmt.Printf("%+v", BrokerConn)
}
