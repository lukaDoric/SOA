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
			reply := []byte(fmt.Sprintf("reply to %s", string(message.Data)))
			err := conn.Publish(message.Reply, reply)
			if err != nil {
				log.Fatal(err)
			}
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	// PRODUCER
	// waiting on the first response
	response, err := conn.Request(subject, []byte("Hello world"), 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("RESPONSE: %s\n", string(response.Data))

	// waiting for consumers to get the message
	//time.Sleep(500 * time.Millisecond)
}

func Conn() *nats.Conn {
	conn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
