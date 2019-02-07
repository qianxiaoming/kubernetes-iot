#!/bin/bash
set -e

if [ $IOT_PROTOCOL = "MQTT" ]; then
  echo "Starting mosquitto broker for MQTT protocol..."
  exec /usr/sbin/mosquitto -c /mosquitto/config/mosquitto.conf
else
  echo "No IoT broker protocol specified:"
  env | sort
  exit 1
fi
