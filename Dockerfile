# Use the official Go image with Debian as the base for building the binary
FROM golang:1.20-bullseye AS builder

# Install necessary build tools and SQLite dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc musl-dev sqlite3 libsqlite3-dev && \
    rm -rf /var/lib/apt/lists/*

# Set environment variables for static linking
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary with static linking
RUN go build -ldflags="-linkmode external -extldflags -static" -o /importarr cmd/main.go

# Use a minimal Alpine image for the final container
FROM alpine:latest

# Install SQLite runtime dependencies
RUN apk add --no-cache sqlite-libs

# Set the working directory
WORKDIR /root/

# Copy the statically built binary from the builder stage
COPY --from=builder /importarr .

# Ensure the binary is executable
RUN chmod +x /root/importarr

# Command to run the binary
CMD ["./importarr"]