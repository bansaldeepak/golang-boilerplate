# Build stage
FROM golang:1.22.2 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /app/main .

# Final stage
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /app/main .

# Set default port and allow it to be overridden by an environment variable
ARG APP_PORT=8080
ENV APP_PORT=${APP_PORT}
EXPOSE ${APP_PORT}

# Command to run the binary
CMD ["sh", "-c", "./main"]