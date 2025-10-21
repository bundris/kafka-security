#!/bin/bash

echo "создаем корневой сертификат и объединяем ключ с сертификатом в бандл ca.pem"
openssl req -new -nodes -x509 -days 365 -newkey rsa:2048 -keyout certs/ca/ca.key -out certs/ca/ca.crt -config certs/ca/ca.cnf
cat certs/ca/ca.crt certs/ca/ca.key > certs/ca/ca.pem

echo "генерируем случайный пароль"
PASSWORD="somesupersafepassword"
BROKERS=(kafka1 kafka2 kafka3)

echo "создаем приватные ключи и CSR"
for broker in "${BROKERS[@]}"; do
    openssl req -new -newkey rsa:2048 -keyout certs/brokers/${broker}/kafka.key -out certs/brokers/${broker}/kafka.csr -config certs/brokers/${broker}/kafka.cnf -nodes
done

echo "Создаем сертификаты для брокеров"
for broker in "${BROKERS[@]}"; do
    openssl x509 -req \
    -days 3650 \
    -in certs/brokers/${broker}/kafka.csr \
    -CA certs/ca/ca.crt \
    -CAkey certs/ca/ca.key \
    -CAcreateserial \
    -out certs/brokers/${broker}/kafka.crt \
    -extfile certs/brokers/${broker}/kafka.cnf \
    -extensions v3_req

done

echo "Создаем PKCS12 хранилища"
for broker in "${BROKERS[@]}"; do
  openssl pkcs12 -export \
    -in certs/brokers/${broker}/kafka.crt \
    -inkey certs/brokers/${broker}/kafka.key \
    -chain \
    -CAfile certs/ca/ca.pem \
    -name kafka1 \
    -out certs/brokers/${broker}/kafka.p12 \
    -password pass:"$PASSWORD"
done

echo "Создаем keystore"
for broker in "${BROKERS[@]}"; do
  keytool -importkeystore \
    -deststorepass "$PASSWORD" \
    -destkeystore certs/brokers/${broker}/kafka.keystore.pkcs12 \
    -srckeystore certs/brokers/${broker}/kafka.p12 \
    -deststoretype PKCS12 \
    -srcstoretype PKCS12 \
    -noprompt \
    -srcstorepass "$PASSWORD"
done

echo "Создаем truststore"
for broker in "${BROKERS[@]}"; do
  keytool -import \
    -file certs/ca/ca.crt \
    -alias ca \
    -keystore certs/brokers/${broker}/kafka.truststore.jks \
    -storepass "$PASSWORD" \
    -noprompt
done

echo "сохраняем пароль"
for broker in "${BROKERS[@]}"; do
    echo "$PASSWORD" > certs/brokers/${broker}/kafka_sslkey_creds
    echo "$PASSWORD" > certs/brokers/${broker}/kafka_keystore_creds
    echo "$PASSWORD" > certs/brokers/${broker}/kafka_truststore_creds
done