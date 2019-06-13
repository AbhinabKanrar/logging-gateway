package main

import (
  "fmt"
  "os"
  "strings"
  "github.com/jinzhu/configor"
)

var Config = struct {
    Mqtt struct {
      URL string `required:"true"`
    } `required:"true"`
}{}


func main() {
  var env = os.Getenv("ENV")

  if strings.EqualFold("dev", env) {
    configor.Load(&Config, "config-dev.json")
  } else if strings.EqualFold("qa", env) {
    configor.Load(&Config, "config-qa.json")
  } else {
    configor.Load(&Config, "config-prod.json")
  }

  fmt.Printf("config: %#v", Config.Mqtt.URL)
}
