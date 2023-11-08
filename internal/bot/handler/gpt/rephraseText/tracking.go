package rephraseText

import (
	"context"
	"fmt"
	configBot "github.com/Li-Khan/go-nasa-bot/config/bot"
	goHttp "github.com/Li-Khan/go-nasa-bot/pkg/service/http"
	"net/http"
	"time"
)

// TrackingResponse represents the structure of the response received while tracking the workflow progress.
type TrackingResponse struct {
	Status string `json:"status"`
	Data   struct {
		Status string `json:"status"`
		Input  struct {
			OriginalText string `json:"original_text"`
		} `json:"input"`
		Output struct {
			SummarizeText                 string `json:"summarize_text"`
			CheckLengthAndIterateIfNeeded string `json:"check_length_and_iterate_if_needed"`
			Output                        string `json:"output"`
		} `json:"output"`
		Metadata struct {
			API bool `json:"api"`
		} `json:"metadata"`
		CreatedAt     time.Time `json:"createdAt"`
		ID            string    `json:"id"`
		WorkflowRunID string    `json:"workflowRunId"`
		WorkflowID    string    `json:"workflowId"`
		Credits       int       `json:"credits"`
	} `json:"data"`
}

// tracking checks the status of the Copy.ai workflow and retrieves the rephrased text.
func tracking(cfg *configBot.Config, id string) (string, error) {
	client := goHttp.NewClient()
	requestHttp, err := client.Request(http.MethodGet, fmt.Sprintf(url+"/%s", cfg.CopyAI.WorkflowID, id))
	if err != nil {
		return "", err
	}
	requestHttp.AddHeader("x-copy-ai-api-key", cfg.CopyAI.APIKey)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	responseHttp, err := requestHttp.DoWithContext(ctx)
	if err != nil {
		return "", err
	}
	defer responseHttp.Close()

	response := TrackingResponse{}
	if err = responseHttp.UnmarshalJSON(&response); err != nil {
		return "", err
	}
	return response.Data.Output.SummarizeText, nil
}
