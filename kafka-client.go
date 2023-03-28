package main

import (
	"kafka-klient/consumer"
	"kafka-klient/model"
	"kafka-klient/producer"
	"time"
)

func main() {
	go produceMessages()
	consumer.Listen()
}

func produceMessages() {
	for {
		producer.Produce(model.Message{
			MdmCode: "123",
			Predecessors: []model.Predecessor{
				{MdmCode: "321", Id: "1"},
				{MdmCode: "321"},
			},
		})
		time.Sleep(30 * time.Second)
	}
}
