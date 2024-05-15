package client

import (
	"context"
	"fmt"
	"net/http"
)

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
func (cln *Client) Completions(ctx context.Context, model Model, prompt string, maxTokens int, temperature float32) (Completion, error) {
	url := fmt.Sprintf("%s/completions", cln.host)

	body := struct {
		Model       string  `json:"model"`
		Prompt      string  `json:"prompt"`
		MaxTokens   int     `json:"max_tokens"`
		Temperature float32 `json:"temperature"`
	}{
		Model:       model.name,
		Prompt:      prompt,
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}

	var resp Completion
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Completion{}, err
	}

	return resp, nil
}
