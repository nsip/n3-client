#!/bin/bash

echo "building command line - privacy ..."
cd ./cli-privacy/
go build -o privacy && mv privacy ../build/ && echo "OK: cli [privacy] is built and put into [./build] directory ... :)"
cd -

echo "building n3client ..."
if [ ! -f ./preprocess/util/jq ]; then
    echo "jq-linux64 not found!, let's get one"
    curl -o jq -L https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && mv jq ./preprocess/util/ && chmod 777 ./preprocess/util/jq
fi

go build && mv n3-client ./build/ && echo "OK: [n3-client] is built and put into [./build] directory ..... :)"

echo "Successful: go into [./build] and run ......................... :)"