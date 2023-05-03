FROM golang:1.17-alpine AS build

WORKDIR /app
COPY . .
RUN go build -o app main.go

FROM alpine:3.14
COPY --from=build /app/app /app
ENTRYPOINT ["/app"]
