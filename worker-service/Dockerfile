FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the worker service code
COPY . .

# Build the worker binary
RUN go build -o worker-service main.go

# Run the worker service
CMD ["./worker-service"]
