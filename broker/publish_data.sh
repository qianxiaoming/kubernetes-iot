#!/bin/bash

while read line
do
  echo -n $line | mosquitto_pub -h 10.107.156.75 -t mqtt-data -s
done < data.json

