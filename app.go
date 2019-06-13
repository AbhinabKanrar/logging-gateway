package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
  "strings"
	"time"
  "github.com/jinzhu/configor"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Config = struct {
    Mqtt struct {
      URL string `required:"true"`
      Topic string `required:"true"`
    } `required:"true"`
}{}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetClientID(clientId)
  opts.SetAutoReconnect(true)
  opts.SetKeepAlive(5 * time.Second)
  opts.SetConnectTimeout(10 * time.Second)
  opts.SetMaxReconnectInterval(10 * time.Second)
	return opts
}

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func listen(uri *url.URL, topic string) {
	client := connect("sub-id", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func main() {
  var env = os.Getenv("ENV")

  if strings.EqualFold("dev", env) {
    configor.Load(&Config, "config/dev.json")
  } else if strings.EqualFold("qa", env) {
    configor.Load(&Config, "config/qa.json")
  } else {
    configor.Load(&Config, "config/prod.json")
  }

	uri, err := url.Parse(Config.Mqtt.URL)

	if err != nil {
		log.Fatal(err)
	}

	topic := Config.Mqtt.Topic

	go listen(uri, topic)

  timer := time.NewTicker(1 * time.Hour)
	for t := range timer.C {
    fmt.Println(t)
	}
}
