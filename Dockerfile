# ---- Build Stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git (required for go mod) and build tools
RUN apk add --no-cache git

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o app .

# ---- Run Stage ----
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Expose Encore's default port (change if needed)
EXPOSE 4000

# Run the application
CMD ["./app"]
