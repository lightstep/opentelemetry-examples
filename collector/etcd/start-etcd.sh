#!/bin/zsh

set -e -x

FLAGS=${@}

FLAGS=$(echo -n "${FLAGS} \n
    --name=${NAME} \n
    --initial-advertise-peer-urls=http://${NAME}:2380 \n
    --listen-peer-urls=http://0.0.0.0:2380 \n
    --listen-client-urls=http://0.0.0.0:2379 \n
    --advertise-client-urls=http://${NAME}:2379 \n
    --heartbeat-interval=250 \n
    --election-timeout=1250 \n
    --initial-cluster-state=new \n
    --listen-metrics-urls=http://0.0.0.0:5050"
)

echo -n "Running etcd ${FLAGS}"
/usr/local/bin/etcd $(echo -n ${FLAGS})
