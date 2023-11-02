package main

import (
	"context"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/pkg/logger"
	"github.com/Li-Khan/go-nasa-bot/pkg/service/http"
	"log"
)

func init() {
	logger.Init("configBot")
	configBot.Get()
}

func main() {
	cfg := configBot.Get().Nasa
	c := http.NewClient()
	req, err := c.Request("GET", cfg.ApodURL)
	if err != nil {
		logger.Error.Fatal(err)
	}
	req.SetQueryParam("api_key", cfg.APIToken)

	resp, err := req.Do(context.Background())
	if err != nil {
		logger.Error.Fatal(err)
	}
	defer func() { _ = resp.Close() }()

	m := make(map[string]interface{})
	err = resp.UnmarshalJSON(&m)
	if err != nil {
		logger.Error.Fatal(err)
	}

	log.Printf("%+v", m)
}
