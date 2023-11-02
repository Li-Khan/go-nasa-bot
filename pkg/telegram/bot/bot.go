package bot

import (
	"github.com/Li-Khan/go-nasa-bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	*tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
}

func NewBot(token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	//bot.Debug = true
	logger.Info.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	return &Bot{
		BotAPI:  bot,
		Updates: updates,
	}, nil
}
