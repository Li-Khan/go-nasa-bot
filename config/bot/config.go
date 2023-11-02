package bot

import (
	"github.com/Li-Khan/go-nasa-bot/pkg/env"
	"github.com/Li-Khan/go-nasa-bot/pkg/logger"
	"github.com/joho/godotenv"
	"sync"
)

const cfgPath string = "./internal/bot/.env"

var (
	config     = &Config{}
	onceConfig = &sync.Once{}
)

type Config struct {
	Telegram struct {
		BotToken string `env:"TELEGRAM_BOT_TOKEN"`
		ChatID   int64  `env:"TELEGRAM_CHAT_ID"`
	}
	Nasa struct {
		APIToken string `env:"NASA_API_TOKEN"`
		ApodURL  string `env:"NASA_APOD_URL"`
	}
}

func Get() *Config {
	onceConfig.Do(func() {
		var err error
		if err = godotenv.Load(cfgPath); err != nil {
			logger.Error.Fatal(err)
		}
		if err = env.Unmarshal(config); err != nil {
			logger.Error.Fatal(err)
		}
	})
	return config
}
