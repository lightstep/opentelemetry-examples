FROM golang:latest

WORKDIR /app
COPY . .

RUN go mod init redis-demo
RUN go mod tidy

RUN go build -o load-generator .

CMD ["./load-generator"]
