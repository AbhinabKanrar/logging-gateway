#!/bin/bash

go get github.com/jinzhu/configor
go get github.com/eclipse/paho.mqtt.golang
go get github.com/elastic/go-elasticsearch

rm -rf logging-gateway

go build

ENV=$1 ./logging-gateway
