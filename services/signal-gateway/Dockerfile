# Stage 1: Build
FROM golang:1.22-alpine AS builder

# This is passed in by GitHub Actions (e.g., ai-agent)
ARG SERVICE_NAME

WORKDIR /app

# 1. Copy the shared internal packages first
COPY pkg/ ./pkg/

# 2. Copy the specific service code
COPY services/${SERVICE_NAME}/ ./services/${SERVICE_NAME}/

# 3. Build the binary from WITHIN the service directory
# This ensures go.mod is found correctly
WORKDIR /app/services/${SERVICE_NAME}
RUN GOWORK=off go build -o /main ./cmd/main.go

# Stage 2: Final Runtime (keeps the image small)
FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /main .

# Start the service
CMD ["./main"]