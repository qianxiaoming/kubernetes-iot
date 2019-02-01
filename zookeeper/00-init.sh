#!/bin/bash

for i in {0..9}
do
  kubectl delete pvc data-zookeeper-$i -n titangrm-iot > /dev/null 2>&1 
  kubectl delete pv  titangrm-iot-zk-pv$((i+1)) > /dev/null 2>&1
  if [ $? -ne 0 ]; then
    break
  fi
done

for node in $*
do
  ssh $node "rm -rf /disk1/pv/zookeeper && mkdir -p /disk1/pv/zookeeper"
done
