package sendAdminErrorMessage

import (
	"fmt"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/pkg/logger"
	goBot "github.com/Li-Khan/go-nasa-bot/pkg/telegram/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handle(format string, v ...any) {
	text := fmt.Sprintf(format, v...)
	logger.Error.Println(text)
	cfg := configBot.Get().Telegram
	msg := tgbotapi.NewMessage(cfg.AdminID, "Error: "+text)
	bot := goBot.Get()
	_, err := bot.Send(msg)
	if err != nil {
		logger.Error.Printf("sendAdminMessage.Handle(): bot.Send(%+v) failed", msg)
		return
	}
}
