#!/bin/bash

go get github.com/jinzhu/configor
go get github.com/kataras/iris
go get github.com/eclipse/paho.mqtt.golang

rm -rf logging-gateway

go build

ENV=$1 ./logging-gateway
