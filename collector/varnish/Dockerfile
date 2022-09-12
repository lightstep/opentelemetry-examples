FROM golang:1.19 as builder

RUN apt update
RUN apt -y install git
RUN git clone https://github.com/jonnenauha/prometheus_varnish_exporter.git
WORKDIR /go/prometheus_varnish_exporter
RUN git checkout 1.6.1
RUN go build


FROM varnish:7.1

COPY --from=builder /go/prometheus_varnish_exporter/prometheus_varnish_exporter /usr/local/bin
