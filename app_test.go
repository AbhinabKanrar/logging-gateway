package main

import (
  "testing"
  "net/url"
)

func TestCreateClientOptions(t *testing.T) {
  uri, err := url.Parse("tcp://127.0.0.1")

  if err != nil {
		t.Fatal("invalid url")
	}

  if CreateClientOptions("cid", uri) == nil {
    t.Fatal("mqtt client cant be created")
  }
}
