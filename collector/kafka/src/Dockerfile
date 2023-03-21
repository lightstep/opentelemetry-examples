# Use an official Go runtime as a parent image
FROM golang:1.19

# Set the working directory to the app directory
WORKDIR /app

# Copy the application files into the container
COPY . .

# Download the Go module dependencies
 RUN go mod download

# Build the Go application
RUN go build -o producer_consumer .

# Start the producer when the container starts
CMD ["./producer_consumer"]
