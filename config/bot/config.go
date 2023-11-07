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
		APIKey  string `env:"NASA_API_KEY"`
		ApodURL string `env:"NASA_APOD_URL"`
	}
	CopyAI struct {
		APIKey     string `env:"COPY_AI_API_KEY"`
		WorkflowID string `env:"COPY_AI_WORKFLOW_ID"`
	}
}

func Get() *Config {
	onceConfig.Do(func() {
		var err error
		if err = godotenv.Load(cfgPath); err != nil {
			logger.Error.Fatalf("Get(): godotenv.Load(cfgPath) failed - %v", err)
		}
		if err = env.Unmarshal(config); err != nil {
			logger.Error.Fatalf("Get(): env.Unmarshal(config) failed - %v", err)
		}
	})
	return config
}
