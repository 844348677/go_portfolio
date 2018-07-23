#!/bin/bash
CNAME=$1
CPID=$(sudo docker inspect --format "{{.State.Pid}}" $CNAME) 
nsenter --target "$CPID" --mount --uts --ipc --net --pid
