package client

import (
	"context"
	"fmt"
	"net/http"
)

// CompletionInput represents the full potential input options for completion.
type CompletionInput struct {
	Model       string
	Prompt      string
	MaxTokens   int
	Temperature *float32
	TopP        *float64
	TopK        *int
}

// CompletionChoice represents a choice for the completion call.
type CompletionChoice struct {
	Text   string `json:"text"`
	Index  int    `json:"index"`
	Status string `json:"status"`
	Model  string `json:"model"`
}

// Completion represents the result for the completion call.
type Completion struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created Time               `json:"created"`
	Choices []CompletionChoice `json:"choices"`
}

// Completions retrieve text completions based on the provided input.
func (cln *Client) Completions(ctx context.Context, input CompletionInput) (Completion, error) {
	url := fmt.Sprintf("%s/completions", cln.host)

	body := struct {
		Model       string   `json:"model"`
		Prompt      string   `json:"prompt"`
		MaxTokens   int      `json:"max_tokens"`
		Temperature *float32 `json:"temperature,omitempty"`
		TopP        *float64 `json:"top_p,omitempty"`
		TopK        *int     `json:"top_k,omitempty"`
	}{
		Model:       input.Model,
		Prompt:      input.Prompt,
		MaxTokens:   input.MaxTokens,
		Temperature: input.Temperature,
		TopP:        input.TopP,
		TopK:        input.TopK,
	}

	var resp Completion
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Completion{}, err
	}

	return resp, nil
}
