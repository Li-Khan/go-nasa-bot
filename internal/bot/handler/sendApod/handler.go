package sendApod

import (
	"fmt"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/entity"
	goBot "github.com/Li-Khan/go-nasa-bot/pkg/telegram/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handle(apod *entity.Apod) error {
	cfg := configBot.Get().Telegram
	msg := tgbotapi.NewMessage(cfg.ChatID, fmt.Sprintf("%+v", apod))
	bot := goBot.Get()
	_, err := bot.Send(msg)
	return err
}
