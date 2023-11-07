package rephraseText

import (
	"fmt"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	goHttp "github.com/Li-Khan/go-nasa-bot/pkg/service/http"
	"net/http"
)

type RunRequest struct {
	RunStartVariables `json:"startVariables"`
	RunMetadata       `json:"metadata"`
}

type RunStartVariables struct {
	OriginalText string `json:"original_text"`
}
type RunMetadata struct {
	API bool `json:"api"`
}

type RunResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID string `json:"id"`
	} `json:"data"`
}

func run(cfg *configBot.Config, text string) (string, error) {
	request := RunRequest{
		RunStartVariables: RunStartVariables{
			OriginalText: text,
		},
		RunMetadata: RunMetadata{
			API: true,
		},
	}

	client := goHttp.NewClient()
	requestHttp, err := client.RequestJSON(http.MethodPost, fmt.Sprintf(url, cfg.CopyAI.WorkflowID), request)
	if err != nil {
		return "", err
	}
	requestHttp.AddHeader("Content-Type", "application/json")
	requestHttp.AddHeader("x-copy-ai-api-key", cfg.CopyAI.APIKey)

	responseHttp, err := requestHttp.DoWithTimeout(30)
	if err != nil {
		return "", err
	}
	defer responseHttp.Close()

	response := RunResponse{}
	if err = responseHttp.UnmarshalJSON(&response); err != nil {
		return "", err
	}

	if response.Status != "success" {
		return "", fmt.Errorf("request failed. status returned: %s", response.Status)
	}
	return response.Data.ID, nil
}
