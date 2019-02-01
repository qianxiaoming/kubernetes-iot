#!/bin/bash

for node in $*
do
  ssh $node "rm -rf /disk1/pv/kafka && mkdir -p /disk1/pv/kafka"
done
