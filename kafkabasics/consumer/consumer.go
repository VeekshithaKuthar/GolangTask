package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "group-1",
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Failed to read message:", err)
		}
		log.Printf("Consumed message: key=%s value=%s\n", string(msg.Key), string(msg.Value))
	}
}
