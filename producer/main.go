package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {

	servers := []string{"localhost:9092"}

	producers, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer producers.Close()

	msg := sarama.ProducerMessage{
		Topic: "bondhello",
		Value: sarama.StringEncoder("Hello World"),
	}

	p, o, err := producers.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Partition=%v, Offset=%v", p, o)
}
