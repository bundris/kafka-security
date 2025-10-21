package consumer

import (
	"encoding/json"
	"kafka-security/config"
	"kafka-security/message"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/fatih/color"
)

func RunSingleConsumer(broker, group, topic string) {
	cfg := config.GetSingleConsumerConfig(broker, group)

	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		color.Green("Failed to create consumer %s: %s", group, err)
	}
	defer c.Close()

	color.Green("Consumer %s started", group)
	err = c.Subscribe(topic, nil)
	if err != nil {
		color.Green("Consumer %s failed to subscribe: %s", group, err)
	}

	color.Green("Consumer %s subscribed to %s",group, topic)
	for {
		ev := c.Poll(100)
		if ev == nil {
			continue
		}
		switch e := ev.(type) {
		case *kafka.Message:
			var m message.Message
			if err := json.Unmarshal(e.Value, &m); err != nil {
				color.Green("Consumer %s deserialization error: %s", group, string(e.Value))
				continue
			}
			upper := strings.ToUpper(m.Text)
			color.Green("Consumer %s processed: %s", group, upper)
		case kafka.Error:
			color.Green("Consumer %s error: %s", group, e)
		default:
			color.Green("Consumer %s default behavior: %s", group, e)
		}
	}
}
