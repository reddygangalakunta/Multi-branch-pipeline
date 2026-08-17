# ==========================================
# Stage 1: Build the Go application
# ==========================================
FROM golang:1.22-alpine AS builder

# Install SSL certificates, timezone data, and git
RUN apk add --no-cache ca-certificates tzdata git

WORKDIR /app

# Copy dependency definition files first to leverage Docker layer caching
COPY go.mod go.sum* ./
RUN go mod download

# Copy application source code
COPY . .

# Build statically linked binary with stripped debug symbols (-s -w)
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /app/server .

# Create non-root user for security
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid 10001 \
    appuser

# ==========================================
# Stage 2: Minimal Runtime Environment
# ==========================================
FROM alpine:3.20 AS runner

# Install essential certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy non-root user and group configuration
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy compiled binary from builder stage
COPY --from=builder /app/server /app/server

# Run as non-root user
USER appuser:appuser

# Expose application port
EXPOSE 8080

# Set default environment variables
ENV PORT=8080 \
    APP_ENV=production

# Health check configuration
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT}/healthz || exit 1

# Execute binary
ENTRYPOINT ["/app/server"]
