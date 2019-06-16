package main

import (
  "testing"
  "net/url"
  service "./service"
)

func TestCreateClientOptions(t *testing.T) {
  uri, err := url.Parse("tcp://127.0.0.1")

  if err != nil {
		t.Fatal("invalid url")
	}

  if service.CreateClientOptions("cid", uri) == nil {
    t.Fatal("mqtt client cant be created")
  }
}
