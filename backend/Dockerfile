# Stage 1: Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Set timezone (optional, adjust as needed)
ENV TZ=Asia/Jakarta

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Copy .env file if needed (or use docker-compose environment variables)
# COPY --from=builder /app/.env .

# Expose port (adjust based on your app's port)
EXPOSE 8080

# Run the application
CMD ["./main"]