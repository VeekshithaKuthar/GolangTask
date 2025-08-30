package messaging

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Messaging struct {
	ChMessaging chan []byte
	Topic       string
	Brokers     []string
}

func NewMessaging(topic string, brokers []string) *Messaging {
	return &Messaging{make(chan []byte), topic, brokers}
}

func (msg *Messaging) Producerecord() {
	if msg.Topic == "" {
		panic("invalid topics")
	}
	if len(msg.Brokers) < 1 {
		panic("invalid brokers")
	}
	c1, err := kgo.NewClient(
		kgo.SeedBrokers(msg.Brokers...),
		kgo.RequiredAcks(kgo.AllISRAcks()),
	)
	if err != nil {
		panic(err)
	}
	defer c1.Close()
	ctx := context.Background()
	for message := range msg.ChMessaging {
		record := &kgo.Record{Topic: msg.Topic, Value: message}

		if err := c1.ProduceSync(ctx, record).FirstErr(); err != nil {
			log.Error().Err(err).Msg("failed to produce record")
		} else {
			log.Info().Str("topic", msg.Topic).
				Bytes("value", message).
				Msg("Record Produced Successfully")
		}
	}
	c1.Flush(ctx)
}
