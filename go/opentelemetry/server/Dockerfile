FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY server.go .
CMD ["go", "run", "server.go"]