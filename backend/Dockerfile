# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies for CGO (SQLite)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy the source code
COPY . .

# Tidy the modules to match the source code and Go version
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/backend-bin ./cmd/server

# Run stage
FROM alpine:latest

# Install sqlite for database management if needed, and ca-certificates
RUN apk add --no-cache ca-certificates sqlite

WORKDIR /app

# Copy the binary and other necessary files
COPY --from=builder /app/backend-bin /app/backend-bin
COPY --from=builder /app/.env .
COPY --from=builder /app/openapi.yaml .
# Also copy dashboard.db if it exists
COPY --from=builder /app/dashboard.db .

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["/app/backend-bin"]
