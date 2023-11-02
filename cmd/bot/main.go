package main

import (
	"context"
	"fmt"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/pkg/logger"
	"github.com/Li-Khan/go-nasa-bot/pkg/service/http"
	goBot "github.com/Li-Khan/go-nasa-bot/pkg/telegram/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func init() {
	logger.Init("configBot")
	configBot.Get()
}

func main() {
	cfg := configBot.Get()
	c := http.NewClient()
	req, err := c.Request("GET", cfg.Nasa.ApodURL)
	if err != nil {
		logger.Error.Fatal(err)
	}
	req.SetQueryParam("api_key", cfg.Nasa.APIToken)
	req.SetQueryParam("date", "2022-11-01")

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

	logger.Info.Printf("%+v", m)

	bot, err := goBot.NewBot(cfg.Telegram.BotToken)
	if err != nil {
		logger.Error.Fatal(err)
	}

	for update := range bot.Updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%+v", m))
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
