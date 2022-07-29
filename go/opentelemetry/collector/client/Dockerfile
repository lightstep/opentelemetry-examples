FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY client.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/client .
CMD ["./client"]