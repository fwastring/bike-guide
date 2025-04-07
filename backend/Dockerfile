
# Stage 1: Build the Go binary
FROM golang:1.23.5 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app and specify the output binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Environment=production" -o /app/bike-guide-api ./cmd/app

# Stage 2: Create the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bike-guide-api /usr/local/bin/bike-guide-api

# Expose the required port (example 8080, adjust if necessary)
EXPOSE 8080

# Command to run the binary
CMD ["/usr/local/bin/bike-guide-api"]

