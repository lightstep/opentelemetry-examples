FROM golang:1.19-rc-alpine3.15 as builder

RUN apk --no-cache add ca-certificates git make zsh bash

WORKDIR /usr/local/src/machine

COPY . .

RUN go mod tidy && go build -o machine.executable

FROM alpine:latest

RUN apk --no-cache add zsh sed curl vim

COPY --from=builder /usr/local/src/machine/machine.executable /usr/local/bin/machine.executable

ENTRYPOINT ["/usr/local/bin/machine.executable"]
