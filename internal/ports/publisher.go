package ports

type EventPublisher interface {
	Publish(topic string, key string, payload []byte) error
}
