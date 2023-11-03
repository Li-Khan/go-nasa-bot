package main

import (
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/cron/getNasaApod"
	"github.com/Li-Khan/go-nasa-bot/pkg/cron"
	"github.com/Li-Khan/go-nasa-bot/pkg/logger"
	"github.com/Li-Khan/go-nasa-bot/pkg/telegram/bot"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	logger.Init("Bot")
	cfg := configBot.Get()
	bot.Init(cfg.Telegram.BotToken)
}

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go cron.RunByFrequency(getNasaApod.Cron, 60*time.Minute)
	<-stop
}
