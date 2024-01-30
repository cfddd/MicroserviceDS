# Build stage
FROM golang:alpine AS builder

LABEL stage=gobuilder

# Set environment variables
ENV CGO_ENABLED 0
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct
ENV PATH="/usr/local/ffmpeg/bin:${PATH}"

# Use Aliyun mirror for faster package installation
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# Install necessary dependencies
RUN apk update --no-cache && apk add --no-cache tzdata

# Set working directory
WORKDIR /build

# Copy Go module files and download dependencies
ADD go.mod .
ADD go.sum .
RUN go mod download

# Copy application source code
COPY . .

# Build the application
RUN go build -ldflags="-s -w" -o /app/main main.go

# Create a volume to mount configuration files
VOLUME /build/config

# Final stage
FROM alpine

# Copy SSL certificates and timezone information from builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai

# Set timezone
ENV TZ Asia/Shanghai

# Set working directory
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main /app/main

# Command to run the application
CMD ["/app/main"]
