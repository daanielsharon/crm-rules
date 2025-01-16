FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o log-worker .

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/log-worker .

CMD ["./log-worker"]
