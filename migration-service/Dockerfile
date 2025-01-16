FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o migration-service main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/migration-service .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/seeds ./seeds

ENTRYPOINT ["./migration-service"]