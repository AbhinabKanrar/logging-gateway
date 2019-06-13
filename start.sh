#!/bin/bash

go get github.com/jinzhu/configor
go get github.com/eclipse/paho.mqtt.golang

ENV=$1 go run app.go
