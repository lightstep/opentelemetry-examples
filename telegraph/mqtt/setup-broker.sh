#!/usr/bin/env bash

docker-compose exec mosquitto-broker mosquitto_passwd -c /mosquitto/conf/mosquitto.passwd mosquitto

# You can also use a batch mechanism
# docker-compose exec mosquitto mosquitto_passwd -b /mosquitto/conf/mosquitto.passwd seconduser shoaCh3ohnokeathal6eeH2marei2o
