# Build stage
FROM golang:1.23-alpine

RUN apk update && apk add --no-cache git make

WORKDIR /app

# Install air for live reloading
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application with debug information
RUN go build -o ./tmp/main ./cmd/main.go

# Expose port
EXPOSE 8083

# Set environment variable for Gin mode
ENV GIN_MODE=release

# Run executable
CMD ["air", "-c", ".air.toml"]
