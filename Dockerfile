# Use a small base image for Go
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go apps
RUN go build -o server .

# Start a new stage from scratch
FROM alpine:latest  

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/server .

# Expose port 8080 to the outside world
EXPOSE 8090

# Command to run the executable
CMD ["./server"]
