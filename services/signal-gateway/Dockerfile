# Stage 1: Build
FROM golang:1.25-alpine AS builder

# This is passed in by GitHub Actions (e.g., ai-agent)
ARG SERVICE_NAME
WORKDIR /app

# 1. Copy the shared internal packages first
COPY pkg/ ./pkg/

# 2. Copy the specific service code
COPY services/${SERVICE_NAME}/ ./services/${SERVICE_NAME}/

# 3. Build the binary from WITHIN the service directory
WORKDIR /app/services/${SERVICE_NAME}
# Check if cmd/main.go exists; if not, build from the root of the service
RUN if [ -f "./cmd/main.go" ]; then \
        GOWORK=off go build -o /main ./cmd/main.go; \
    else \
        GOWORK=off go build -o /main ./main.go; \
    fi

# Stage: Final Runtime (keeps the image small)
FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /main .

# Start the service
CMD ["./main"]