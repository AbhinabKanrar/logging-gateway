package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
  "github.com/jinzhu/configor"
	service "./service"
)

var Config = struct {
  Mqtt struct {
    URL string `required:"true"`
    Topic string `required:"true"`
  } `required:"true"`

	ES struct {
    URL string `required:"true"`
		Index string `required:"true"`
  } `required:"true"`
}{}

func main() {
	configor.Load(&Config, fmt.Sprintf("./config/%s.json", os.Getenv("ENV")))

	uri, err := url.Parse(Config.Mqtt.URL)

	if err != nil {
		log.Fatal(err)
	}

	go service.Listen(uri, Config.Mqtt.Topic)
	go service.EsInitialize(Config.ES.URL, Config.ES.Index)

	for{

	}
}
