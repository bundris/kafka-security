package producer

import (
	"encoding/json"
	"math/rand"
	"time"

	"kafka-security/config"
	"kafka-security/message"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/fatih/color"
)

func RunProducer(broker string, topic string, scrapTime int, retries int) {
	cfg := config.GetProducerConfig(broker, retries)
	interval := time.Minute / time.Duration(scrapTime)

	p, err := kafka.NewProducer(cfg)
	if err != nil {
		color.Red("Failed to create producer: %s", err)
	}
	defer p.Close()

	color.Red("Producer started")
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					color.Red("Delivery failed: %s", ev.TopicPartition.Error)
				} else {
					color.Red("Delivered message to %s", ev.TopicPartition)
				}
			}
		}
	}()

	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
	for {
		m := message.Message{
			Id:   randomID(),
			Text: randomText(),
		}
		b, err := json.Marshal(m)
		if err != nil {
			color.Red("Serialization error: %s", err)
			time.Sleep(interval)
			continue
		}
		color.Red("Producing message for topic %s: %s", topic, m)
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          b,
		}, nil)
		if err != nil {
			color.Red("Produce error on topic %s: %s", topic, err)
		}
		time.Sleep(interval)
	}
}

const letters = "abcdefghijklmnopqrstuvwxyz"

func randomID() string {
	return time.Now().Format("20060102T150405") + "-" + randSeq(6)
}

func randomText() string {
	// length from 5 to 30
	length := 5 + rand.Intn(26)
	return randSeq(length)
}

func randSeq(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
