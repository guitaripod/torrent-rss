# Build stage
FROM golang:1.22.5-alpine AS builder

# Install git for private repos if needed
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code maintaining the directory structure
COPY cmd/ ./cmd/
COPY internal/ ./internal/

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/torrent-rss

# Final stage
FROM alpine:latest

# Add necessary runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Create directory for downloads
RUN mkdir -p /downloads

# Set default environment variables
ENV TD_BASE_URL=https://www.torrentday.com \
    TD_DOWNLOAD_PATH=/downloads \
    TD_CHECK_INTERVAL="0 */12 * * *"

# Run the binary
CMD ["./main"]