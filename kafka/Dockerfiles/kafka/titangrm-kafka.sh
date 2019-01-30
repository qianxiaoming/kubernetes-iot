#!/bin/bash

export KAFKA_BROKER_ID=$((${HOSTNAME##*-} + 1))
export KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://$(hostname -f):9092

. /etc/confluent/docker/run
