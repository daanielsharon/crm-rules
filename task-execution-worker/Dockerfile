FROM golang:1.23 AS builder

WORKDIR /app/shared
COPY shared .
RUN go mod download

WORKDIR /app/task-execution-worker
COPY task-execution-worker .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/task-execution-worker .

CMD ["./main"]