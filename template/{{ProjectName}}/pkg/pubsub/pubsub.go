package pubsub

import "time"

type PubSubMessage struct {
	Value     []byte
	Timestamp time.Time
}

type PubSub interface {
	Subscribe(topics []string) error
	ReadMessage() (*PubSubMessage, error)
	Ack() error
	Close() error
}
