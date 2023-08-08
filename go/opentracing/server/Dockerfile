FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY server.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:3.18.3  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/server .
CMD ["./server"]