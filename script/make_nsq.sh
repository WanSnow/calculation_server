#!/bin/bash

HOST=127.0.0.1
PORT=9697

docker run --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd
docker run --name nsqd -p 4150:4150 -p 4151:4151 \
    nsqio/nsq /nsqd \
    --broadcast-address=${HOST} \
    --lookupd-tcp-address=${HOST}:${PORT}