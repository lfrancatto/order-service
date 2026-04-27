package consumers

import (
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"order-service/internal/ports"
)

// Consumer handles Kafka message consumption with retry and DLQ
type Consumer struct {
	consumer   *kafka.Consumer
	publisher  ports.EventPublisher
	processor  func([]byte) error
	maxRetries int
	retryDelay time.Duration
	dlqTopic   string
}

// NewConsumer creates a new Kafka consumer
func NewConsumer(
	brokers string,
	groupID string,
	processor func([]byte) error,
	publisher ports.EventPublisher,
) (*Consumer, error) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
		"enable.auto.commit": false,
	})

	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer:   c,
		publisher:  publisher,
		processor:  processor,
		maxRetries: 3,
		retryDelay: 2 * time.Second,
		dlqTopic:   "orders-dlq",
	}, nil
}

// Start begins consuming messages from Kafka
func (c *Consumer) Start(topic string) {
	c.consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			log.Println("consumer error:", err)
			continue
		}

		c.handleMessage(msg)
	}
}

// handleMessage processes a message with retry and DLQ logic
func (c *Consumer) handleMessage(msg *kafka.Message) {
	var err error

	for i := 0; i < c.maxRetries; i++ {
		err = c.processor(msg.Value)

		if err == nil {
			_, commitErr := c.consumer.CommitMessage(msg)
			if commitErr != nil {
				log.Println("commit error:", commitErr)
			}
			return
		}

		log.Printf("retry %d failed: %v\n", i+1, err)
		time.Sleep(c.retryDelay)
	}

	// Send to DLQ after retries exhausted
	log.Println("sending message to DLQ")

	err = c.publisher.Publish(
		c.dlqTopic,
		string(msg.Key),
		msg.Value,
	)

	if err != nil {
		log.Println("failed to publish to DLQ:", err)
	}

	// Commit even on failure (avoid poison message loop)
	_, commitErr := c.consumer.CommitMessage(msg)
	if commitErr != nil {
		log.Println("commit error:", commitErr)
	}
}