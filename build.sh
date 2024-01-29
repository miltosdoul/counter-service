#!/usr/bin/env bash

docker build -t counter-service-test .
docker run -it -d \
    -e LISTEN_PORT=8081 \
    -p 8081:8081 \
    counter-service-test