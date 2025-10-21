# Этап сборки
FROM golang:1.24-bookworm AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y \
librdkafka-dev \
&& rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o main .

# Финальный этап
FROM debian:bookworm-slim

WORKDIR /app
RUN apt-get update && apt-get install -y \
ca-certificates \
librdkafka1 \
&& rm -rf /var/lib/apt/lists/*
# Копируем скомпилированное приложение из builder
COPY --from=builder /app/main .
COPY --from=builder /app/certs ./certs
# Запуск приложения
CMD ["./main"]