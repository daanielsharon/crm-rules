# Use a minimal base image with Go installed
FROM golang:1.20-alpine

RUN apk add --no-cache bash build-base git openssh postgresql-client

# Install golang-migrate for database migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2

WORKDIR /app

COPY . .

RUN go build -o migration-service main.go

COPY ./migrations ./migrations

ENTRYPOINT ["/app/migration-service"]
