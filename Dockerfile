# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

RUN go mod tidy

# Build the application with debug information
RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-api ./cmd/main.go

# Final stage
FROM alpine:latest

# Add basic debugging tools
RUN apk add --no-cache curl tzdata

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /golang-api /golang-api

# Make sure the binary is executable
# RUN chmod +x main

# Expose port
EXPOSE 8083

# Run with more verbose output
CMD ["/golang-api"]