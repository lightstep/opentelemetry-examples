FROM golang:1.19-rc-alpine3.15 as builder

ARG ETCD_VERSION=3.4.19

RUN apk --no-cache add ca-certificates git make zsh bash

WORKDIR /usr/local/src

RUN git clone https://github.com/etcd-io/etcd \
	&& cd etcd && git checkout 4636a5fab43f94bdb35962c90492474f50ac8682 \
    && make

FROM alpine:latest

ARG ETCD_NODE_NAME=$ETCD_NODE_NAME

RUN apk --no-cache add zsh sed

COPY --from=builder /usr/local/src/etcd/bin/etcd /usr/local/bin/etcd
COPY --from=builder /usr/local/src/etcd/bin/etcdctl /usr/local/bin/etcdctl
COPY start-etcd.sh /opt/start-etcd.sh
COPY etcd-config.yaml /etc/etcd.conf

RUN chmod u+x /opt/start-etcd.sh

ENTRYPOINT ["/opt/start-etcd.sh"]
