package main

import (
	"fmt"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"github.com/Li-Khan/go-nasa-bot/pkg/logger"
	"log"
)

func init() {
	logger.Init("configBot")
	configBot.Get()
}

func main() {
	log.Printf("%+v", configBot.Get())
	fmt.Println("Hello, World!")
}
