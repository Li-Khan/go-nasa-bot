package sendApod

import (
	"fmt"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/entity"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/handler/gpt/rephraseText"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/handler/youtube/downloadVideo"
	goBot "github.com/Li-Khan/go-nasa-bot/pkg/telegram/bot"
	translategooglefree "github.com/bas24/googletranslatefree"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strings"
)

const msgFormat string = `<b>%s</b>

%s

Автор: %s
%s: %s`

// Handle processes the Astronomy Picture of the Day (APOD) and sends it to a Telegram channel.
// It translates the title to Russian, rephrases the explanation, and sends either a photo or video message based on the media type.
func Handle(apod *entity.Apod) error {
	if apod.Copyright == "" {
		apod.Copyright = "NASA"
	}
	var err error
	apod.Title, err = translategooglefree.Translate(apod.Title, "en", "ru")
	if err != nil {
		return err
	}
	apod.Explanation, err = rephraseText.Handle(apod.Explanation)
	if err != nil {
		return err
	}

	switch apod.MediaType {
	case "image":
		return sendPhoto(apod)
	case "video":
		return sendVideo(apod)
	default:
		return fmt.Errorf("unknown media_type - %s", apod.MediaType)
	}
}

func sendPhoto(apod *entity.Apod) error {
	text := fmt.Sprintf(msgFormat, apod.Title, apod.Explanation, apod.Copyright, "Фото", apod.URL)
	cfg := configBot.Get().Telegram
	photo := tgbotapi.NewPhoto(cfg.ChatID, tgbotapi.FileURL(apod.URL))
	photo.Caption = strings.TrimSpace(text)
	photo.ParseMode = "HTML"
	bot := goBot.Get()
	_, err := bot.Send(photo)
	return err
}

func sendVideo(apod *entity.Apod) error {
	if err := downloadVideo.Handle(apod.URL); err != nil {
		return err
	}
	text := fmt.Sprintf(msgFormat, apod.Title, apod.Explanation, apod.Copyright, "Видео", apod.URL)
	cfg := configBot.Get().Telegram
	video := tgbotapi.NewVideo(cfg.ChatID, tgbotapi.FilePath("./video.mp4"))
	video.ParseMode = "HTML"
	video.Caption = text
	bot := goBot.Get()
	_, err := bot.Send(video)
	if err != nil {
		return err
	}
	return os.Remove("./video.mp4")
}
