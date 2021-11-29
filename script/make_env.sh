#!/bin/bash

MYSQL_PASSWORD=`cat ./config/config.json `| jq '.mysql_password'
MYSQL_VERSION=5.6.17
REDIS_VERSION=3.2.12

docker run -itd --name mysql-service -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${MYSQL_PASSWORD} mysql:${MYSQL_VERSION}
docker run -itd --name redis-service -p 6379:6379 redis:${REDIS_VERSION}