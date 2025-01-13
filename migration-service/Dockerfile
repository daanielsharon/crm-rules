FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

COPY ./migrations /app/migrations

RUN go build -o migration-service main.go

ENTRYPOINT ["/app/migration-service"]
