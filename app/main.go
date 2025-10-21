package main

import (
	"kafka-security/consumer"
	"kafka-security/producer"
	"log"
	"os"
	"os/signal"
)

func main() {
	brokers := "kafka1:9092,kafka2:9093,kafka3:9094"
	go producer.RunProducer(brokers, "topic-1", 60, 5)
	go producer.RunProducer(brokers, "topic-2", 60, 5)
	go consumer.RunSingleConsumer(brokers, "consumer-group1", "topic-1")
	go consumer.RunSingleConsumer(brokers, "consumer-group2", "topic-2")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	log.Println("Exit app")
}
