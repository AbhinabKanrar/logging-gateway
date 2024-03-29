package service

import (
  "fmt"
	"log"
	"net/url"
	"time"
  "strings"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func CreateClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetClientID(clientId)
  opts.SetAutoReconnect(true)
  opts.SetKeepAlive(5 * time.Second)
  opts.SetConnectTimeout(10 * time.Second)
  opts.SetMaxReconnectInterval(10 * time.Second)
	return opts
}

func Connect(clientId string, uri *url.URL) mqtt.Client {
	opts := CreateClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()

	for !token.WaitTimeout(3 * time.Second) {
	}

	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func Listen(uri *url.URL, topic string) {
	client := Connect("logging-gateway-broker", uri)

	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
    go save(strings.Split(msg.Topic(),"/")[3], string(msg.Payload()))
	})
}
