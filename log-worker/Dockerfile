FROM golang:1.23 AS builder

WORKDIR /app/shared
COPY shared .
RUN go mod download

WORKDIR /app/log-worker
COPY log-worker .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/log-worker .

CMD ["./main"]