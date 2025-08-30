package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"orders/models"

	"github.com/twmb/franz-go/pkg/kgo"
)

type OrderProducer struct {
	client *kgo.Client
	topic  string
}

func NewOrderProducer(brokers []string, topic string) *OrderProducer {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.RequiredAcks(kgo.AllISRAcks()),
	)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	return &OrderProducer{
		client: cl,
		topic:  topic,
	}
}

func (p *OrderProducer) SendOrder(order *models.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}
	record := &kgo.Record{
		Topic: p.topic,
		Value: data,
	}
	fmt.Print("the topic is", record.Topic)
	// p.client.Produce(context.Background(), record, nil)
	// if err != nil {
	// 	log.Printf("Failed to send order to kafka: %v,err")
	// } else {
	// 	log.Printf("Successfully sent order ID %d to Kafka topic '%s'", order.ID, p.topic)
	// }
	res := p.client.ProduceSync(context.Background(), record)
	err = res.FirstErr()
	if err != nil {
		log.Printf("Failed to send order to Kafka: %v", err)
		return err
	}
	log.Printf("Successfully sent order ID %d to Kafka topic '%s'", order.ID, p.topic)
	return nil

}
