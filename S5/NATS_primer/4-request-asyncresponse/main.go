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
	replySubject := "reply"
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
	_, err := conn.Subscribe(replySubject, func(message *nats.Msg) {
		fmt.Printf("RESPONSE: %s\n", string(message.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
	err = conn.PublishRequest(subject, replySubject, []byte("hello world"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("message sent, doing something else in the meantime ...")

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
