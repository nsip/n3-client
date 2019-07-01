#!/bin/bash

if [ ! -f ./preprocess/util/jq ]; then
    echo "jq-linux64 not found!, let's get one"
    curl -o jq -L https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && mv jq ./preprocess/util/ && chmod 777 ./preprocess/util/jq
fi

go build && mv n3-client ./build/ && echo "[n3-client] is built and put into [./build] directory, go into [./build] and run..."
