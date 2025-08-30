package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "test-topic",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	for i := 0; i < 10; i++ {
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("key-" + string(rune(i))),
				Value: []byte("Hello Kafka Message " + time.Now().Format(time.RFC3339)),
			})
		if err != nil {
			log.Fatal("Failed to write message:", err)
		}
		log.Println("Produced message", i)
		time.Sleep(1 * time.Second)
	}
}
