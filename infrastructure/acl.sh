#!/bin/bash

# Сперва сбрасываем на всякий случай старые права
docker exec kafka1 kafka-acls --authorizer-properties zookeeper.connect=zookeeper:2181 --remove User:user1 --topic topic-1 --force
docker exec kafka1 kafka-acls --authorizer-properties zookeeper.connect=zookeeper:2181 --remove User:user2 --topic topic-2 --force

docker exec kafka1 kafka-acls --authorizer-properties zookeeper.connect=zookeeper:2181  --add \
    --allow-principal User:user1 \
    --operation READ \
    --group consumer-group1 
docker exec kafka1 kafka-acls --authorizer-properties zookeeper.connect=zookeeper:2181  --add \
    --allow-principal User:user1 \
    --operation DESCRIBE \
    --group consumer-group1
docker exec kafka1 kafka-acls --authorizer-properties zookeeper.connect=zookeeper:2181 --add \
    --allow-principal User:user1 \
    --operation READ \
    --topic topic-1