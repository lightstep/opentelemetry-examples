FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go .
CMD ["go", "run", "client.go"]