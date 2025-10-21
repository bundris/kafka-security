package config

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func GetSingleConsumerConfig(broker, group string) *kafka.ConfigMap {
	var user, password string
	if group == "consumer-group1" {
		user = "user1"
		password = "user1-secret"
	} else {
		user = "user2"
		password = "user2-secret"
	}
	return &kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"group.id":           group,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": true,
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.mechanism":    "PLAIN",
        "sasl.username":     user,
        "sasl.password":     password,
	}
}
