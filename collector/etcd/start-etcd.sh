#!/bin/zsh

if [[ -z ETCD_CONFIG ]]; then
    echo "running with custom config"
    sed -i "s/etcd-name/${ETCD_NODE_NAME}/g" ${ETCD_CONFIG}
    /usr/local/bin/etcd --config-file ${ETCD_CONFIG}
    exit 0
    echo "running with default config"
    /usr/local/bin/etcd
    exit 0
fi


echo "running with custom config"
sed -i "s/etcd-name/${ETCD_NODE_NAME}/g" ${ETCD_CONFIG}
/usr/local/bin/etcd --config-file ${ETCD_CONFIG}
