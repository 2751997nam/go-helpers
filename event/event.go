package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Event struct {
	Name string
	Data any
}

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"chillitee", //name of the exchange
		"topic",     //type
		true,        //durable
		false,       //auto-deleted
		false,       //internal
		false,       //no-wait
		nil,         //args
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    //name
		false, //durable
		false, //delete when unuse
		true,  //exclusive
		false, //nowait
		nil,   //args
	)
}
