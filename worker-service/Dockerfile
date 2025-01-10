# Use a minimal base image with Go installed
FROM golang:1.20-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy dependency files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the worker service code
COPY . .

# Build the binary with static linking and Linux target
RUN CGO_ENABLED=0 GOOS=linux go build -o worker-service .

# Use minimal alpine image for the final stage
FROM alpine:latest

WORKDIR /root/

# Copy only the compiled binary from builder stage
COPY --from=builder /app/worker-service .

# Run the worker service
CMD ["./worker-service"]
