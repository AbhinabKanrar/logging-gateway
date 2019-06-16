package main

import (
	"log"
	"net/url"
	"os"
  "strings"
  "github.com/jinzhu/configor"
	service "./service"
)

var Config = struct {
	Port string `required:"true"`

  Mqtt struct {
    URL string `required:"true"`
    Topic string `required:"true"`
  } `required:"true"`
}{}

func main() {
  var env = os.Getenv("ENV")

  if strings.EqualFold("dev", env) {
    configor.Load(&Config, "./config/dev.json")
  } else if strings.EqualFold("qa", env) {
    configor.Load(&Config, "./config/qa.json")
  } else {
    configor.Load(&Config, "./config/prod.json")
  }

	uri, err := url.Parse(Config.Mqtt.URL)

	if err != nil {
		log.Fatal(err)
	}

	go service.Listen(uri, Config.Mqtt.Topic)

	for{

	}
}
