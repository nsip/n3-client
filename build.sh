#!/bin/bash

set -e
ORIGINALPATH=`pwd`

rm -rf ./build

mkdir -p build/Win64
cp config.toml ./build/Win64/
cp main_test.txt ./build/Win64/main_test.go
cp main.go ./build/Win64/main.go
mkdir -p ./build/Win64/debug_pub
mkdir -p ./build/Win64/debug_qry

mkdir -p build/Linux64
cp config.toml ./build/Linux64/
cp main_test.txt ./build/Linux64/main_test.go
cp main.go ./build/Linux64/main.go
mkdir -p ./build/Linux64/debug_pub
mkdir -p ./build/Linux64/debug_qry

mkdir -p build/Mac
cp config.toml ./build/Mac/
cp main_test.txt ./build/Mac/main_test.go
cp main.go ./build/Mac/main.go
mkdir -p ./build/Mac/debug_pub
mkdir -p ./build/Mac/debug_qry

mv $GOPATH/src/github.com/cdutwhu/go-wrappers $GOPATH/src/github.com/cdutwhu/go-wrappers_OLD_$$

OK=0
echo "building command line - privacy ..."
go get github.com/cdutwhu/go-gjxy
GOPATH=`go env GOPATH`
cd $GOPATH/src/github.com/cdutwhu/go-gjxy
git fetch
git checkout master #30b39b932d92afa40d71cf77b941bd44110399b1
git pull
cd $ORIGINALPATH

GOPATH=`go env GOPATH`
cd $GOPATH/src/github.com/cdutwhu/go-util
git fetch
git checkout master
git pull
cd $ORIGINALPATH

cd ./cli-privacy/
go get github.com/inconshreveable/mousetrap # for windows spf13
go get

GOARCH=amd64
LDFLAGS="-s -w"

GOOS="linux" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o privacy
mv privacy ../build/Linux64/
GOOS="windows" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o privacy.exe
mv privacy.exe ../build/Win64/
GOOS="darwin" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o privacy
mv privacy ../build/Mac/

echo "OK: cli [privacy] is built and put into [./build] directory ... :)"
OK=1
cd -

OK=0
echo "fetching utils ..."

JQ="jq-linux64"
if [ ! -f ./preprocess/util/$JQ ]; then
    echo "$JQ not found!, let's get one"
    curl -o $JQ -L https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && mv $JQ ./preprocess/util/ && chmod 777 ./preprocess/util/$JQ
fi
cp ./preprocess/util/$JQ ./build/Linux64/jq
cp ./preprocess/util/$JQ ./preprocess/util/jq

JQ="jq-win64.exe"
if [ ! -f ./preprocess/util/$JQ ]; then
    echo "$JQ not found!, let's get one"
    curl -o $JQ -L https://github.com/stedolan/jq/releases/download/jq-1.6/jq-win64.exe && mv $JQ ./preprocess/util/
fi
cp ./preprocess/util/$JQ ./build/Win64/jq.exe

JQ="jq-mac"
if [ ! -f ./preprocess/util/$JQ ]; then
    echo "$JQ not found!, let's get one"
    curl -o $JQ -L https://github.com/stedolan/jq/releases/download/jq-1.6/jq-osx-amd64 && mv $JQ ./preprocess/util/ && chmod 777 ./preprocess/util/$JQ
fi
cp ./preprocess/util/$JQ ./build/Mac/jq

echo "building n3client ..."
go get
GOOS="linux" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o n3-client
mv n3-client ./build/Linux64/
GOOS="windows" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o n3-client.exe
mv n3-client.exe ./build/Win64/
GOOS="darwin" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o n3-client
mv n3-client ./build/Mac/

echo "OK: [n3-client] is built and put into [./build] directory ..... :)"
OK=1

echo $OK

if [ $OK == 1 ]; then
    echo "Successful: go into [./build] and run ......................... :)"
else
    echo "Failed: ....................................................... :("
fi
