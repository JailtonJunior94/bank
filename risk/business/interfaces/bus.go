package interfaces

import "github.com/streadway/amqp"

type IRabbitMQ interface {
	GetChannel() *amqp.Channel
}
