#!/bin/bash

export CONNECT_REST_ADVERTISED_HOST_NAME=$(hostname -f)
echo "rest.advertised.host.name=$CONNECT_REST_ADVERTISED_HOST_NAME"

. /etc/confluent/docker/run
