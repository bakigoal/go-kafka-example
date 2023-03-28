package main

import (
	"kafka-klient/consumer"
	"kafka-klient/producer"
	"time"
)

func main() {
	go consumer.Listen()
	for {
		producer.Produce("hello world!!!")
		time.Sleep(5 * time.Second)
	}
}
