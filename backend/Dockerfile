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

# Build the Go app
RUN go build -o bike-guide-api .

# Stage 2: Create the final image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bike-guide-api .

# Expose port 8080 (or the port your Go app uses)
EXPOSE 8080

# Command to run the binary
CMD ["./bike-guide-api"]
