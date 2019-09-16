#!/bin/bash

set -e
ORIGINALPATH=`pwd`

rm -rf ./build/
mkdir build && cp config.toml ./build/ && cp main_test.txt ./build/main_test.go && cp main.go ./build/main.go
mkdir ./build/debug_pub && mkdir ./build/debug_qry

OK=0
echo "building command line - privacy ..."
go get github.com/cdutwhu/go-gjxy
GOPATH=`go env GOPATH`
cd $GOPATH/src/github.com/cdutwhu/go-gjxy
git checkout master #30b39b932d92afa40d71cf77b941bd44110399b1
cd $ORIGINALPATH

cd ./cli-privacy/
go get
go build -o privacy && mv privacy ../build/ && echo "OK: cli [privacy] is built and put into [./build] directory ... :)" && OK=1
cd -

OK=0
echo "building n3client ..."
if [ ! -f ./preprocess/util/jq ]; then
    echo "jq-linux64 not found!, let's get one"
    curl -o jq -L https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && mv jq ./preprocess/util/ && chmod 777 ./preprocess/util/jq
fi
go get
go build && mv n3-client ./build/ && echo "OK: [n3-client] is built and put into [./build] directory ..... :)" && OK=1

echo $OK

if [ $OK == 1 ]; then
    echo "Successful: go into [./build] and run ......................... :)"
else
    echo "Failed: ....................................................... :("
fi
