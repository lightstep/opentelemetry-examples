# Use an official Go runtime as a parent image
FROM golang:1.19

# Set the working directory to the app directory
WORKDIR /app

# Copy the application files into the container
COPY . .

# Download the Go module dependencies
RUN go mod init opentelementry-examples-kafka
RUN go mod tidy

# Build the Go application
RUN go build -o producer .

# Start the producer when the container starts
CMD ["./producer"]
