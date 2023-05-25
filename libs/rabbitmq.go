package libs

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var mqlock = &sync.Mutex{}

type MQ struct {
	Conn *amqp.Connection
}

var singleInstance *MQ

func GetMQ() *amqp.Connection {
	if singleInstance == nil {
		mqlock.Lock()
		defer mqlock.Unlock()
		conn, err := connect()
		if err != nil {
			log.Fatal(err)
		}

		singleInstance = &MQ{
			Conn: conn,
		}
	}

	return singleInstance.Conn
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	//don't continue util rabbitmq is ready
	for {
		c, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
		if err != nil {
			fmt.Println("rabbitmq not yet ready")
			counts++
		} else {
			fmt.Println("connected to rabbitmq")
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
