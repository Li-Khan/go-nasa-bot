package rephraseText

import (
	"context"
	"fmt"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	goHttp "github.com/Li-Khan/go-nasa-bot/pkg/service/http"
	"net/http"
	"time"
)

// RunRequest represents the structure of the request to start the Copy.ai workflow.
// It includes start variables and metadata.
type RunRequest struct {
	RunStartVariables `json:"startVariables"`
	RunMetadata       `json:"metadata"`
}

// RunStartVariables holds the original text for rephrasing.
type RunStartVariables struct {
	OriginalText string `json:"original_text"`
}

// RunMetadata contains metadata for the Copy.ai API request.
type RunMetadata struct {
	API bool `json:"api"`
}

// RunResponse holds the response data after initiating the workflow.
type RunResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID string `json:"id"`
	} `json:"data"`
}

// run sends a request to Copy.ai API to initiate the rephrasing process.
// It returns the ID associated with the workflow run.
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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	responseHttp, err := requestHttp.DoWithContext(ctx)
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
