# ���� ������
FROM golang:1.20 AS builder

WORKDIR /app

# �������� ������ go.mod (go.sum ����� �������������)
COPY go.mod ./

# ��������� �����������
RUN go mod tidy

# �������� �������� ���
COPY . .

# �������� ����������
RUN go build -o app

# ��������� �����
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

CMD ["./app"]