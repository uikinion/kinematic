
FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN go build -o app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

CMD ["./app"]
