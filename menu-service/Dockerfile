# Stage 1: Build the Go binary
FROM golang:1.26 AS builder

WORKDIR /app

# Copy dependency files first for better Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build with CGO enabled (required for SQLite's C bindings)
RUN CGO_ENABLED=1 GOOS=linux go build -o menu-service .

# Stage 2: Create a slim runtime image
FROM debian:bookworm-slim

# Install the SQLite shared library needed at runtime
RUN apt-get update && apt-get install -y libsqlite3-0 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/menu-service .

EXPOSE 8080

CMD ["./menu-service"]
