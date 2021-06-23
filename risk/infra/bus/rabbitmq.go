package bus

import (
	"log"

	"github.com/jailtonjunior94/bank/risk/business/interfaces"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMQ(connectionString string) interfaces.IRabbitMQ {
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return &RabbitMQ{Connection: connection, Channel: channel}
}

func (r *RabbitMQ) GetChannel() *amqp.Channel {
	return r.Channel
}
