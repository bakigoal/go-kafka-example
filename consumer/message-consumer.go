package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"kafka-klient/model"
	"time"
)

const (
	bootstrapServers = "localhost"
	topic            = "golang-topic"
	groupId          = "my-group"
	offsetReset      = "earliest"
)

func Listen() {
	c := subscribeToTopic(topic)
	defer c.Close()

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			go handleMessage(model.FromByteArray(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func handleMessage(message model.Message) {
	fmt.Printf("Handling Message: %q \n", message)
}

func subscribeToTopic(topic string) *kafka.Consumer {
	config := kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          groupId,
		"auto.offset.reset": offsetReset,
	}
	c, err := kafka.NewConsumer(&config)

	if err != nil {
		panic(any(err))
	}

	c.SubscribeTopics([]string{topic}, nil)
	return c
}
