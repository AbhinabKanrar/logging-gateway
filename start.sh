#!/bin/bash

go get github.com/jinzhu/configor
go get github.com/eclipse/paho.mqtt.golang
go get gopkg.in/olivere/elastic.v6

rm -rf logging-gateway

go build

ENV=$1 ./logging-gateway
