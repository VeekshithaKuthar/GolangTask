// package workerbo

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math"
// 	"orders/models"
// 	"orders/redis"

// 	"github.com/twmb/franz-go/pkg/kgo"
// )

// func StartBOConsumer(brokers []string, topic string) {
// 	fmt.Printf("DEBUG: brokers=%v, topic=%q\n", brokers, topic)

// 	cl, err := kgo.NewClient(
// 		kgo.SeedBrokers(brokers...),
// 		kgo.ConsumeTopics(topic),
// 		kgo.ConsumerGroup("bo-worker-group"),
// 		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
// 		//kgo.WithLogger(kgo.BasicLogger(os.Stdout, kgo.LogLevelDebug, nil)),
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to create Kafka consumer: %v", err)
// 	}
// 	defer cl.Close()

// 	fmt.Println("BO Worker started and listening...")

// 	ctx := context.Background()

// 	for {
// 		fetches := cl.PollFetches(ctx)
// 		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
// 			fmt.Printf("📦 Partition fetched from topic=%q, partition=%d, recordCount=%d\n",
// 				p.Topic, p.Partition, len(p.Records))
// 			for _, record := range p.Records {
// 				fmt.Printf("📝 Record Key=%s, Value=%s\n", string(record.Key), string(record.Value))
// 				fmt.Println("the record is 2", record)
// 				var order models.Order
// 				if err := json.Unmarshal(record.Value, &order); err != nil {
// 					fmt.Println("Invalid Kafka message:", err)
// 					continue
// 				}

// 				fmt.Printf("Received order: %+v\n", order)

// 				currentNAV, err := redis.GetNAV(order.SchemeCode)
// 				if err != nil {
// 					fmt.Printf("Failed to get NAV from Redis: %v\n", err)
// 					continue
// 				}
// 				updatedNAV := calculateNewNAV(currentNAV, order.Side, order.Amount)

// 				err = redis.SetNAV(order.SchemeCode, updatedNAV)
// 				if err != nil {
// 					fmt.Printf("Redis NAV update failed: %v\n", err)
// 					continue
// 				}

// 				fmt.Printf("NAV for %s updated: %.4f → %.4f\n", order.SchemeCode, currentNAV, updatedNAV)
// 			}
// 		})
// 	}
// }

// func calculateNewNAV(currentNAV float64, side string, amount float64) float64 {
// 	change := amount / 100000.0

// 	if side == "BUY" {
// 		return math.Round((currentNAV+change)*10000) / 10000
// 	}
// 	return math.Round((currentNAV-change)*10000) / 10000
// }

package workerbo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"orders/models"
	"orders/redis"

	"github.com/twmb/franz-go/pkg/kgo"
)

func StartBOConsumer(brokers []string, topic string) {
	fmt.Printf("DEBUG: brokers=%v, topic=%q\n", brokers, topic)

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumeTopics(topic),
		kgo.ConsumerGroup("bo-worker-group"),
		// 👇 This ensures the consumer reads from the beginning
		// if the group has no committed offsets yet.
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	)
	fmt.Print("hii")
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer cl.Close()

	fmt.Println("BO Worker started and listening...")

	ctx := context.Background()

	for {

		fetches := cl.PollFetches(ctx)
		if ctx.Err() != nil || fetches.IsClientClosed() {
			break
		}
		if errs := fetches.Errors(); len(errs) > 0 {
			for _, fe := range errs {
				fmt.Printf("⚠️ Fetch error: %+v\n", fe)
			}
		}

		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
			fmt.Printf("📦 Partition fetched from topic=%q, partition=%d, recordCount=%d\n",
				p.Topic, p.Partition, len(p.Records))

			for _, record := range p.Records {
				fmt.Printf("📝 Record Key=%s, Value=%s\n",
					string(record.Key), string(record.Value))

				var order models.Order
				if err := json.Unmarshal(record.Value, &order); err != nil {
					fmt.Println("Invalid Kafka message:", err)
					continue
				}

				fmt.Printf("✅ Received order: %+v\n", order)

				currentNAV, err := redis.GetNAV(order.SchemeCode)
				if err != nil {
					fmt.Printf("Failed to get NAV from Redis: %v\n", err)
					continue
				}
				updatedNAV := calculateNewNAV(currentNAV, order.Side, order.Amount)

				err = redis.SetNAV(order.SchemeCode, updatedNAV)
				if err != nil {
					fmt.Printf("Redis NAV update failed: %v\n", err)
					continue
				}

				fmt.Printf("📊 NAV for %s updated: %.4f → %.4f\n",
					order.SchemeCode, currentNAV, updatedNAV)
			}
		})
	}
}

func calculateNewNAV(currentNAV float64, side string, amount float64) float64 {
	change := amount / 100000.0

	if side == "BUY" {
		return math.Round((currentNAV+change)*10000) / 10000
	}
	return math.Round((currentNAV-change)*10000) / 10000
}
