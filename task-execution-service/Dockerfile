# Use a minimal base image with Go installed
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o task-execution-service .

# Final stage: Use a minimal alpine image
FROM alpine:latest

# Set working directory for the final container
WORKDIR /root/

# Copy only the built binary from the builder stage
COPY --from=builder /app/task-execution-service .

# Command to run when container starts
CMD ["./task-execution-service"]
