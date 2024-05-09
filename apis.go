package client

import (
	"context"
	"fmt"
	"net/http"
)

// HealthCheck validates the PG API Service is available.
func (cln *Client) HealthCheck(ctx context.Context) (string, error) {
	var resp string
	if err := cln.rawRequest(ctx, http.MethodGet, cln.host, nil, &resp); err != nil {
		return "", err
	}

	return resp, nil
}

// ChatCompletions generate chat completions based on a conversation history.
func (cln *Client) ChatCompletions(ctx context.Context, model string, messages []Message, maxTokens int, temperature float32) (ChatCompletion, error) {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	body := struct {
		Model       string    `json:"model"`
		Messages    []Message `json:"messages"`
		MaxTokens   int       `json:"max_tokens"`
		Temperature float32   `json:"temperature"`
	}{
		Model:       model,
		Messages:    messages,
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}

	var resp ChatCompletion
	if err := cln.rawRequest(ctx, http.MethodPost, url, body, &resp); err != nil {
		return ChatCompletion{}, err
	}

	return resp, nil
}

// ChatCompletionsSSE generate chat completions based on a conversation history.
func (cln *Client) ChatCompletionsSSE(ctx context.Context, model string, messages []Message, maxTokens int, temperature float32, ch chan ChatCompletion) error {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	body := struct {
		Model    string    `json:"model"`
		Messages []Message `json:"messages"`
		Stream   bool      `json:"stream"`
	}{
		Model:    model,
		Messages: messages,
		Stream:   true,
	}

	sse := sseClient[ChatCompletion]{
		Client: cln,
	}

	if err := sse.rawRequest(ctx, http.MethodPost, url, body, ch); err != nil {
		return err
	}

	return nil
}
