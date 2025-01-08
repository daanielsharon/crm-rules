# Use a minimal base image with Go installed
FROM golang:1.20-alpine

# Install dependencies
RUN apk add --no-cache bash build-base git openssh postgresql-client

# Install golang-migrate for database migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Set working directory inside the container
WORKDIR /app

# Copy migration service code
COPY . .

# Build the migration service binary
RUN go build -o migration-service main.go

# Copy the migrations folder into the container
COPY ./migrations ./migrations

# Make the migration service binary the entrypoint
ENTRYPOINT ["/app/migration-service"]
