# syntax=docker/dockerfile:1
FROM golang:1.25.1-alpine3.21 AS builder

ARG DB_CONNECTION

# Install git and ca-certificates (needed for go get if any)
RUN apk update && apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Install tools
RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Download wait tool
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

# Copy source code
COPY . .

# Proper CompileDaemon command
CMD ["sh", "-c", "\
    /wait && \
    goose -dir './internal/infrastructure/database/migrations' mysql \"${DB_CONNECTION}\" up && \
    CompileDaemon \
      -log-prefix=false \
      -color=true \
      -build=\"go build -o main ./cmd/service\" \
      -command=\"./main\" \
"]