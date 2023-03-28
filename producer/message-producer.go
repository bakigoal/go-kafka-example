package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"kafka-klient/model"
)

const (
	bootstrapServers = "localhost"
	topic            = "golang-topic"
)

func Produce(message model.Message) {
	p := createProducer(bootstrapServers)
	defer p.Close()
	t := topic
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &t, Partition: kafka.PartitionAny},
		Value:          message.ToByteArray(),
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

func createProducer(servers string) *kafka.Producer {
	config := kafka.ConfigMap{
		"bootstrap.servers": servers,
	}
	p, err := kafka.NewProducer(&config)
	if err != nil {
		panic(any(err))
	}
	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	return p
}
