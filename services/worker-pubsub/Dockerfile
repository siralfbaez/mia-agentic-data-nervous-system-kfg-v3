# Stage 1: Build
FROM golang:1.22-alpine AS builder

# This argument is passed in during 'docker build'
ARG SERVICE_NAME

WORKDIR /app

# 1. Copy workspace and shared packages
COPY go.work go.work.sum ./
COPY pkg/ ./pkg/

# 2. Copy the specific service folder using the ARG
COPY services/${SERVICE_NAME}/ ./services/${SERVICE_NAME}/

# 3. Build the binary (The path is now dynamic!)
RUN go build -o /main ./services/${SERVICE_NAME}/cmd/main.go

# Stage 2: Final Runtime
FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder
COPY --from=builder /main .

CMD ["./main"]
