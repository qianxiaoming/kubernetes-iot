#!/bin/bash

export KAFKA_BROKER_ID=$((${HOSTNAME##*-} + 1))
echo "broker.id=$KAFKA_BROKER_ID"

if [ -z $DEFAULT_PORT ]; then
  export DEFAULT_PORT=9092
fi
export KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://$(hostname -f):$DEFAULT_PORT
echo "advertised.listeners=$KAFKA_ADVERTISED_LISTENERS"

if [ -z $ZK_PREFIX ]; then
  export ZK_PREFIX=zookeeper
fi
if [ -z $ZK_CLUSTER ]; then
  export ZK_CLUSTER=zookeeper-ensemble
fi
if [ -z $ZK_PORT ]; then
  export ZK_PORT=2181
fi
if [ -z $ZK_CHROOT ]; then
  export ZK_CHROOT=iot-default
fi

ZK_SERVERS=
for i in {0..255} # max number of nodes in zookeeper cluster is 255
do
  ping -c 2 $ZK_PREFIX-$i.$ZK_CLUSTER > /dev/null 2>&1
  if [ $? -eq 0 ]; then
    if [ -z "$ZK_SERVERS" ]; then
      ZK_SERVERS="$ZK_PREFIX-$i.$ZK_CLUSTER:$ZK_PORT"
    else
      ZK_SERVERS=$ZK_SERVERS",$ZK_PREFIX-$i.$ZK_CLUSTER:$ZK_PORT"
    fi
  else
    break
  fi
done
export KAFKA_ZOOKEEPER_CONNECT=$ZK_SERVERS"/"$ZK_CHROOT
echo "zookeeper.connect=$KAFKA_ZOOKEEPER_CONNECT"

. /etc/confluent/docker/run
