FROM debian:stable-slim

ENV DEBIAN_FRONTEND noninteractive
COPY 50docker-apt-conf /etc/apt/apt.conf.d/

COPY rootfs_prefix/ /usr/src/rootfs_prefix/

RUN apt-get update \
 && apt-get upgrade \
 && apt-get install \
    collectd-core \
    collectd-utils \
    build-essential \
    zsh \
 && make -C /usr/src/rootfs_prefix/ \
 && apt-get --purge remove build-essential \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

COPY collectd.conf /etc/collectd/collectd.conf
COPY run-collectd.sh /usr/local/sbin/run-collectd.sh

RUN chmod u+x /usr/local/sbin/run-collectd.sh

ENV LD_PRELOAD /usr/src/rootfs_prefix/rootfs_prefix.so

EXPOSE 9103

ENTRYPOINT ["/usr/local/sbin/run-collectd.sh", "8"]
