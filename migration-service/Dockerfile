FROM golang:1.23 AS builder

WORKDIR /app/shared
COPY shared .
RUN go mod download

WORKDIR /app/migration-service
COPY migration-service .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/migration-service .
# COPY --from=builder /app/migrations ./migrations
# COPY --from=builder /app/seeds ./seeds

CMD ["./main"]