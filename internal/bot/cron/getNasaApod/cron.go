package getNasaApod

import (
	"context"
	"fmt"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/entity"
	"github.com/Li-Khan/go-nasa-bot/internal/bot/handler/sendApod"
	"github.com/Li-Khan/go-nasa-bot/pkg/logger"
	goHttp "github.com/Li-Khan/go-nasa-bot/pkg/service/http"
	"io/ioutil"
	"net/http"
	"os"
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
	response, err := request.Do(ctx)
	cancel()
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
	file, err := ioutil.ReadFile("./last_date.txt")
	if err != nil {
		fmt.Println(err)
	}
	if string(file) == apod.Date {
		return
	}
	write1, _ := os.Create("./last_date.txt")
	_, err = write1.Write([]byte(apod.Date))
	if err != nil {
		logger.Error.Printf("Cron(): write1.Write([]byte(apod.Date)) failed - %v", err)
		return
	}

	if err = sendApod.Handle(apod); err != nil {
		logger.Error.Printf("Cron(): sendApod.Handle(apod) failed - %v", err)
		return
	}
}
