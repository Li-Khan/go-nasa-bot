package sendApod

import (
	"fmt"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/entity"
	goBot "github.com/Li-Khan/go-nasa-bot/pkg/telegram/bot"
	gt "github.com/bas24/googletranslatefree"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handle(apod *entity.Apod) error {
	cfg := configBot.Get().Telegram
	photo := tgbotapi.NewPhoto(cfg.ChatID, tgbotapi.FileURL(apod.Hdurl))
	photo.Caption = fmt.Sprintf("<b>%s</b>\n\n%s\n\nAuthor: %s\nPhoto: %s\nDate: %s", apod.Title, apod.Explanation, apod.Copyright, apod.Hdurl, apod.Date)
	photo.ParseMode = "HTML"
	var err error
	photo.Caption, err = gt.Translate(photo.Caption, "en", "kk")
	if err != nil {
		return err
	}
	bot := goBot.Get()
	_, err = bot.Send(photo)
	return err
}
