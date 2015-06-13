#!/bin/bash

BASE=$(pwd)

cd $GOPATH/src/github.com/toomore/plivo-go
go get -v ./...

cd $GOPATH/bin
cp ./plivoSendMass $BASE
cp ./plivoQuickSend $BASE

cd $BASE
docker build -t toomore/plivo-go-mini .

rm -rf ./plivoSendMass
rm -rf ./plivoQuickSend
