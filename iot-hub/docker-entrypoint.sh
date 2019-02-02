#!/bin/bash
set -e

if [ $IOT_PROTOCOL = "MQTT" ]; then
  echo "Starting mosquitto broker for mqtt protocol..."
  exec /usr/sbin/mosquitto -c /mosquitto/config/mosquitto.conf
else
  echo "No IoT protocol specified via IOT_PROTOCOL!"
  exit 1
fi
