#!/bin/bash

export INFLUX_UDP_ADDR="192.168.1.200:8089"
export KAFKA_BROKERS="192.168.1.201:9092"
export MONGODB_DATABASE_NAME="zinnion"
export MONGODB_CONNECTION="192.168.1.100:27017"
export MONGODB_USERNAME="mongo-admin"
export MONGODB_PASSWORD="Br@sa154"

if pgrep -x "winter" > /dev/null
then
    echo "Running"
else
    echo "Stopped"
    /root/work/src/github.com/maurodelazeri/winter/winter --venues $(hostname -s | tr '[:lower:]' '[:upper:]') 2>&1 | tee -a /var/log/winter.log
fi
