package getNasaApod

import (
	"context"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/entity"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/handler/sendApod"
	"github.com/Li-Khan/go-nasa-bot/pkg/file"
	"github.com/Li-Khan/go-nasa-bot/pkg/logger"
	goHttp "github.com/Li-Khan/go-nasa-bot/pkg/service/http"
	"net/http"
	"time"
)

func Cron() {
	cfg := configBot.Get().Nasa
	client := goHttp.NewClient()
	request, err := client.Request(http.MethodGet, cfg.ApodURL)
	if err != nil {
		logger.Error.Printf("Cron(): client.Request(http.MethodGet, cfg.Nasa.ApodURL) failed - %v", err)
		return
	}
	request.SetQueryParam("api_key", cfg.APIToken)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, err := request.Do(ctx)
	if err != nil {
		logger.Error.Printf("Cron(): request.Do(ctx) failed - %v", err)
		return
	}
	defer response.Close()

	apod := &entity.Apod{}
	if err = response.UnmarshalJSON(apod); err != nil {
		logger.Error.Printf("Cron(): response.UnmarshalJSON(&apod) failed - %v", err)
		return
	}

	lastDate, err := file.OpenAndOverwriteFile("./last_date.txt", apod.Date)
	if lastDate == apod.Date {
		return
	}

	if err = sendApod.Handle(apod); err != nil {
		logger.Error.Printf("Cron(): sendApod.Handle(apod) failed - %v", err)
		return
	}
}
