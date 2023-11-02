package bot

import (
	"github.com/Li-Khan/go-nasa-bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
)

type Bot struct {
	*tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
}

var (
	bot     *Bot
	botOnce = &sync.Once{}
)

func Init(token string) {
	botOnce.Do(func() {
		b, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			logger.Error.Fatalf("NewBot(): tgbotapi.NewBotAPI(token) failed - %v", err)
		}
		//bot.Debug = true
		logger.Info.Printf("Authorized on account %s", b.Self.UserName)
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates := b.GetUpdatesChan(u)
		bot = &Bot{
			BotAPI:  b,
			Updates: updates,
		}
	})
}

func Get() *Bot {
	if bot == nil {
		logger.Error.Fatal("Get(): bot is nil")
	}
	return bot
}
