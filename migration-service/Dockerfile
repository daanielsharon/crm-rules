FROM golang:1.23-alpine

RUN apk add --no-cache bash build-base git openssh postgresql-client

# Install golang-migrate for database migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY . .

RUN go build -o migration-service main.go

COPY ./migrations ./migrations

ENTRYPOINT ["/app/migration-service"]
