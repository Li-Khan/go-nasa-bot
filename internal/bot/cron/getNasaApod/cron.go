package getNasaApod

import (
	"context"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/handler/telegram/sendAdminErrorMessage"
	"net/http"
	"os"
	"time"

	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/entity"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/handler/telegram/sendApod"
	"github.com/Li-Khan/go-nasa-bot/pkg/file"
	goHttp "github.com/Li-Khan/go-nasa-bot/pkg/service/http"
)

type Error struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Cron is a scheduled task that fetches the Astronomy Picture of the Day (APOD)
// from NASA's API and sends it to a Telegram channel.
func Cron() {
	cfg := configBot.Get().Nasa
	client := goHttp.NewClient()
	request, err := client.Request(http.MethodGet, cfg.ApodURL)
	if err != nil {
		sendAdminErrorMessage.Handle("Cron(): client.Request(http.MethodGet, cfg.Nasa.ApodURL) failed - %v", err)
		return
	}
	request.SetQueryParam("api_key", cfg.APIKey)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response, err := request.DoWithContext(ctx)
	if err != nil {
		sendAdminErrorMessage.Handle("Cron(): request.DoWithTimeout(ctx) failed - %v", err)
		return
	}
	defer response.Close()

	if response.GetStatusCode() != http.StatusOK {
		sendAdminErrorMessage.Handle("Cron(): response status code - %v\nbody - %s", response.GetStatusCode(), string(response.GetBody()))
		return
	}

	apod := &entity.Apod{}
	if err = response.UnmarshalJSON(apod); err != nil {
		sendAdminErrorMessage.Handle("Cron(): response.UnmarshalJSON(&apod) failed - %v", err)
		return
	}
	apod.Normalize()

	lastDate, err := file.OpenAndOverwriteFile("./last_date.txt", apod.Date)
	if err != nil {
		sendAdminErrorMessage.Handle("Cron(): file.OpenAndOverwriteFile('./last_date.txt', apod.Date) failed - %v", err)
		return
	}
	if lastDate == apod.Date {
		return
	}

	if err = sendApod.Handle(apod); err != nil {
		sendAdminErrorMessage.Handle("Cron(): sendApod.Handle(apod) failed - %v", err)
		if err = os.Remove("./last_date.txt"); err != nil {
			sendAdminErrorMessage.Handle("Cron(): os.Remove(\"./last_date.txt\") failed - %v", err)
		}
		return
	}
}
