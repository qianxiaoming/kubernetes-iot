#!/bin/bash

for node in $*
do
  ssh $node "rm -rf /disk1/pv/zookeeper && mkdir -p /disk1/pv/zookeeper"
done
