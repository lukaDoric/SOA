package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	conn := Conn()
	defer conn.Close()

	subject := "sub"
	queue := "queue"
	subs := 5

	// CONSUMERS
	for i := 0; i < subs; i++ {
		_, err := conn.QueueSubscribe(subject, queue, func(message *nats.Msg) {
			fmt.Printf("RECEIVED MESSAGE: %s\n", string(message.Data))
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	// PRODUCER
	err := conn.Publish(subject, []byte("hello world!"))
	if err != nil {
		log.Fatal(err)
	}

	// waiting for consumers to get the message
	time.Sleep(500 * time.Millisecond)
}

func Conn() *nats.Conn {
	conn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
