package service

import (
  "fmt"
	elasticsearch "github.com/elastic/go-elasticsearch"
)

func EsInitialize(topic string) {
  cfg := elasticsearch.Config{
    Addresses: []string{
      "http://localhost:9200",
      "http://localhost:9201",
    },
  }

  es, err := elasticsearch.NewClient(cfg)

  if err != nil {
    fmt.Println("Error creating the client: %s", err)
  }
}
