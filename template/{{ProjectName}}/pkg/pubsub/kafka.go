package pubsub

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"math/rand"
	"time"
	"{{ProjectName}}/pkg"
)

type KafkaPubSub struct {
	log      *pkg.Logger
	consumer *kafka.Consumer
}

type KafkaOption struct {
	Servers     string
	GroupId     string
	OffsetReset string
}

func NewKafka(log *pkg.Logger, option KafkaOption) *KafkaPubSub {
	rand.Seed(time.Now().Unix())
	var gId = fmt.Sprintf("groupid_%d", rand.Int31())
	if option.GroupId != "random" {
		gId = option.GroupId
	}
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": option.Servers,
		"group.id":          gId,
		"auto.offset.reset": option.OffsetReset,
	})

	if err != nil {
		log.Fatal(err)
	}
	return &KafkaPubSub{
		log:      log,
		consumer: consumer,
	}
}

func (pb *KafkaPubSub) Subscribe(topics []string) error {
	pb.log.Debug("Subscribe to kafka topic:", topics)
	return pb.consumer.SubscribeTopics(topics, nil)
}

func (pb *KafkaPubSub) ReadMessage() (*PubSubMessage, error) {
	message, err := pb.consumer.ReadMessage(-1)
	if err != nil {
		return nil, err
	}

	return &PubSubMessage{Value: message.Value, Timestamp: message.Timestamp}, nil
}

func (pb *KafkaPubSub) Ack() error {
	panic("implement me")
}

func (pb *KafkaPubSub) Close() error {
	panic("implement me")
}
