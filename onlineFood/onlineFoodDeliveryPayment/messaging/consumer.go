package messaging

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"os/signal"
	"paymenst/models"
	"paymenst/repositories"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer struct {
	client    *kgo.Client
	paymentDb repositories.IPaymentDB
	topic     string
	groupID   string
}

func NewConsumer(brokers []string, topic, group string, paymentDb repositories.IPaymentDB) (*Consumer, error) {
	c1, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumeTopics(topic),
		kgo.ConsumerGroup(group),
	)
	if err != nil {
		return nil, err
	}
	return &Consumer{client: c1, paymentDb: paymentDb, topic: topic, groupID: group}, nil
}

func (c *Consumer) Run() {
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
		if ctx.Err() != nil {
			log.Info().Msg("context cancelled, stopping consumer loop")
			break
		}

		// if client is closed, break
		if fetches.IsClientClosed() {
			log.Info().Msg("Kafka client closed, exiting consumer loop")
			break
		}
		if errs := fetches.Errors(); len(errs) > 0 {
			for _, e := range errs {
				log.Error().Err(e.Err).Str("topic", e.Topic).Int32("partition", e.Partition).Msg("fetch error")
			}
			continue
		}

		fetches.EachRecord(func(record *kgo.Record) {
			//var payment models.Payments
			var userorder models.UserOrders
			err := json.Unmarshal(record.Value, &userorder)
			if err != nil {
				log.Error().Err(err).Msg("failed to unmarshal order")
				return
			}

			log.Info().Msgf("Received order: %+v", userorder)

			payment := &models.Payments{
				OrderID:       userorder.OrderID,
				PaymentAmount: userorder.Amount,
				PaymentStatus: "Pending",
			}

			_, err = c.paymentDb.CreatePayment(payment)
			if err != nil {
				log.Error().Err(err).Msg("failed to create payment")
				return
			}

			if rand.Int()%2 == 0 {
				payment.PaymentStatus = "SUCCESS"
			} else {
				payment.PaymentStatus = "Failure"
			}

			err = c.paymentDb.UpdatePaymentStatus(payment.OrderID, payment.PaymentStatus)
			if err != nil {
				log.Error().Err(err).Msg("failed to update payment status")
				return
			}

			evtBytes, _ := json.Marshal(map[string]interface{}{
				"order_id": payment.OrderID,
				"status":   payment.PaymentStatus,
			})
			err = c.client.ProduceSync(ctx, &kgo.Record{
				Topic: "payment_status_updated",
				Value: evtBytes,
			}).FirstErr()
			if err != nil {
				log.Error().Err(err).Msg("failed to publish payment status")
			} else {
				log.Info().Msgf("Published payment status for order %s: %s", payment.OrderID, payment.PaymentStatus)
			}
		})

	}

}
