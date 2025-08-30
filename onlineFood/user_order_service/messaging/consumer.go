package messaging

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"userorders/repository"

	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
)

type PaymentConsumer struct {
	client  *kgo.Client
	orderDB repository.IOrderRepository
	topic   string
	groupID string
}

func NewPaymentConsumer(brokers []string, topic, group string, orderDB repository.IOrderRepository) (*PaymentConsumer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumeTopics(topic),
		kgo.ConsumerGroup(group),
	)
	if err != nil {
		return nil, err
	}

	return &PaymentConsumer{client: client, orderDB: orderDB, topic: topic, groupID: group}, nil
}

func (c *PaymentConsumer) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Info().Msg("shutting down consumer...")
		c.client.Close()
		cancel()
	}()

	for {
		fetches := c.client.PollFetches(ctx)
		if ctx.Err() != nil || fetches.IsClientClosed() {
			break
		}

		fetches.EachRecord(func(record *kgo.Record) {
			var evt struct {
				OrderID string `json:"order_id"`
				Status  string `json:"status"`
			}

			if err := json.Unmarshal(record.Value, &evt); err != nil {
				log.Error().Err(err).Msg("failed to unmarshal payment event")
				return
			}

			order, err := c.orderDB.UpdateOrderStatus(evt.OrderID, evt.Status)
			if err != nil {
				log.Error().Err(err).Msg("failed to update order status")
				return
			}

			log.Info().Msgf("Updated order %s status to %s", order.OrderID, order.Status)
		})
	}
}
