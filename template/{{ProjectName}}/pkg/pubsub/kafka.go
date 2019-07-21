package pubsub

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"{{ProjectName}}/pkg"
	"math/rand"
	"sync"
	"time"
)

type KafkaPubSub struct {
	consumer *kafka.Consumer
}

var (
	once        sync.Once
	kafkaPubSub *KafkaPubSub
)

func GetKafka() *KafkaPubSub {
	once.Do(func() {
		kafkaPubSub = NewKafka()
	})
	return kafkaPubSub
}

func NewKafka() *KafkaPubSub {
	return &KafkaPubSub{}
}

func (pb *KafkaPubSub) Initialize(servers string, groupId string, offsetReset string) {
	rand.Seed(time.Now().Unix())
	var gId = fmt.Sprintf("groupid_%d", rand.Int31())
	if groupId != "random" {
		gId = groupId
	}
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          gId,
		"auto.offset.reset": offsetReset,
	})

	if err != nil {
		pkg.GetLog().Fatal(err)
	}
	pb.consumer = consumer
}

func (pb *KafkaPubSub) Subscribe(topics []string) error {
	pkg.GetLog().Debug("Subscribe to kafka topic:", topics)
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
