package config

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func GetProducerConfig(broker string, retries int) *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": broker,
		"acks":              "all",
		"retries":           retries,
		"security.protocol": "SASL_SSL",
        "sasl.mechanism":    "PLAIN",
        "sasl.username":     "admin",
        "sasl.password":     "admin-secret",
		"ssl.ca.location":   "/app/certs/ca.crt",
	}
}
