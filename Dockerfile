# Stage 1: Build the Go application
FROM golang:1.24.11-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
# The -s -w flags are used to strip the debug information and symbol tables, reducing the binary size
RUN go build -o bookstore ./cmd/main/main.go

# Stage 2: Create a minimal image to run the application
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /app/bookstore /usr/local/lib/bookstore/bookstore

# Expose port 8080
EXPOSE 8080

# Command to run the application when the container starts
CMD ["/usr/local/lib/bookstore/bookstore"]
