package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Producer struct {
	producer *kafka.Producer
}

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

func (p *Producer) Publish(topic string, key string, payload []byte) error {
	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Key:            []byte(key),
		Value:          payload,
	}, nil)
}
