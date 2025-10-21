#!/bin/bash

docker exec kafka1 kafka-topics \
  --create \
  --topic topic-1 \
  --partitions 3 \
  --replication-factor 2 \
  --bootstrap-server kafka1:9092 \
  --command-config /etc/kafka/adminclient-configs.conf

docker exec kafka1 kafka-topics \
  --create \
  --topic topic-2 \
  --partitions 3 \
  --replication-factor 2 \
  --bootstrap-server kafka1:9092 \
  --command-config /etc/kafka/adminclient-configs.conf
