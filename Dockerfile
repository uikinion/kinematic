# Этап сборки
FROM golang:1.20 AS builder

WORKDIR /app

# Копируем только go.mod (go.sum может отсутствовать)
COPY go.mod ./

# Загружаем зависимости
RUN go mod tidy

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o app

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

CMD ["./app"]