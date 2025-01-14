FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/user-service .

CMD ["/app/user-service"]