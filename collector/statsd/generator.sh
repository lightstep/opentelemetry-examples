#!/bin/zsh

-x

for (( ; ; ))
do
    echo "sending metric"
    GAUGE_EC01=$(( ( RANDOM % 50 )  + 1 ))
    GAUGE_EC02=$(( ( RANDOM % 50 )  + 1 ))
    echo -n "collectd.ec01.nginx.nginx_connections-active:${GAUGE_EC01}|g" | nc -u -w0 localhost 8125
    echo -n "collectd.ec01.nginx.nginx_connections-active:${GAUGE_EC01}|g" | nc -u -w0 localhost 8125
    echo -n "collectd.ec02.nginx.nginx_connections-active:${GAUGE_EC02}|g" | nc -u -w0 localhost 8125
    echo "Send:"
    echo " [*] ec01: ${GAUGE_EC01}"
    echo " [*] ec02: ${GAUGE_EC02}"
done
