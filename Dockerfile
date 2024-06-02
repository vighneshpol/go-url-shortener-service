# Use the official Go image as a base image
FROM golang:1.22.1-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Optionally, tidy up dependencies
RUN go mod tidy

# Build the application
RUN go build -o go-url-shortener-service

# Use a lightweight base image for the final image
FROM alpine:latest

# Install CA certificates and add a non-root user
RUN apk --no-cache add ca-certificates \
    && adduser -D -g '' appuser

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=build /app/go-url-shortener-service .

# Change ownership of the application directory
RUN chown -R appuser /app

# Switch to the non-root user
USER appuser

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./go-url-shortener-service"]
