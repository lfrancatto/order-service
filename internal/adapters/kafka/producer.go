package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

// Producer is a Kafka adapter that implements EventPublisher
type Producer struct {
	producer *kafka.Producer
}

// NewProducer creates a Kafka producer with idempotence enabled.
// This guarantees no duplicate messages in case of retries.
func NewProducer(brokers string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"acks":               "all",
		"enable.idempotence": true,
	})

	if err != nil {
		return nil, err
	}

	return &Producer{producer: p}, nil

}

// Publish sends a message to Kafka
func (p *Producer) Publish(topic string, key string, payload []byte) error {
	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(key),
		Value: payload,
	}, nil)
}
