package service

import (
  "context"
  "fmt"
	elasticsearch "gopkg.in/olivere/elastic.v6"
)

func EsInitialize(url string, index string) {
  ctx := context.Background()
  es, err := elasticsearch.NewClient(elasticsearch.SetURL(url))

  if err != nil {
    fmt.Println("Error creating the client: %s", err)
  } else {
    exists, err := es.IndexExists(index).Do(ctx)

    if err == nil && !exists {
      createIndex, indexErr := es.CreateIndex(index).Do(ctx)

      if indexErr != nil {
        fmt.Println("bawal")
      }
    } else {
      fmt.Println("ache")
    }
  }
}

func save(siteId string, payload string) {
  fmt.Println(siteId)
  fmt.Println(payload)
}
