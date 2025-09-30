# Build stage
FROM golang:1.23-alpine AS builder

# Install git for go modules
RUN apk add --no-cache git curl

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o chezmoi-tui .

# Final stage
FROM alpine:latest

# Install chezmoi and other dependencies
RUN apk --no-cache add git openssh-client curl && \
    curl -fsSL https://get.chezmoi.io | sh

# Create a non-root user
RUN addgroup -g 65532 nobody && \
    adduser -D -u 65532 -G nobody nobody

# Create home directory for user
RUN mkdir -p /home/nobody

# Copy the binary from builder stage
COPY --from=builder /app/chezmoi-tui /usr/local/bin/chezmoi-tui

# Change ownership
RUN chown -R nobody:nobody /home/nobody

# Set non-root user
USER nobody

# Set working directory
WORKDIR /home/nobody

# Set entrypoint
ENTRYPOINT ["chezmoi-tui"]