FROM golang:1.19-rc-alpine3.15 as builder

ARG ETCD_VERSION=3.5.4

RUN apk --no-cache add ca-certificates git make zsh bash

WORKDIR /usr/local/src

RUN git clone https://github.com/etcd-io/etcd \
	&& cd etcd && git checkout 08407ff7600eb16c4445d5f21c4fafaf19412e24 \
    && make

FROM alpine:latest

ARG ETCD_NODE_NAME=$ETCD_NODE_NAME

RUN apk --no-cache add zsh sed curl vim

COPY --from=builder /usr/local/src/etcd/bin/etcd /usr/local/bin/etcd
COPY --from=builder /usr/local/src/etcd/bin/etcdctl /usr/local/bin/etcdctl
COPY start-etcd.sh /opt/start-etcd.sh
COPY etcd.conf /etc/etcd.conf

RUN chmod u+x /opt/start-etcd.sh

EXPOSE 5050 2379 2380

ENTRYPOINT ["/opt/start-etcd.sh"]
