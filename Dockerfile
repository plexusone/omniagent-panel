# Build stage
FROM golang:1.24-bookworm AS builder

# Install opus dependencies for audio encoding
RUN apt-get update && apt-get install -y --no-install-recommends \
    libopus-dev \
    libopusfile-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with opus support
RUN CGO_ENABLED=1 GOOS=linux go build -tags opus -o /panel ./cmd/panel

# Runtime stage
FROM debian:bookworm-slim

# Install runtime dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    libopus0 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /panel /app/panel

# Create non-root user
RUN useradd -r -u 1000 panel
USER panel

# Health check
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD ["/app/panel", "-health-check"]

# Expose health port
EXPOSE 8080

ENTRYPOINT ["/app/panel"]
