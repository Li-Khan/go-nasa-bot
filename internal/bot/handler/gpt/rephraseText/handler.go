package rephraseText

import (
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	"time"
)

const url string = "https://api.copy.ai/api/workflow/%s/run"

// Handle processes the provided text using Copy.ai API to rephrase it.
// It initiates the workflow, monitors its progress, and returns the rephrased text.
func Handle(text string) (string, error) {
	conf := configBot.Get()
	id, err := run(conf, text)
	if err != nil {
		return "", err
	}
	var rephrase string
	for i := 0; ; i++ {
		rephrase, err = tracking(conf, id)
		if err != nil {
			return "", err
		}
		if rephrase != "" {
			break
		}
		time.Sleep(10 * time.Second)
	}
	return rephrase, nil
}
