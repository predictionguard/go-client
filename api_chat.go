package client

import (
	"context"
	"fmt"
	"net/http"
)

// ChatMessage represents the role of the sender and the content to process.
type ChatMessage struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
	Output  string `json:"output"`
}

// Chat represents the result for the chat completion call.
type Chat struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created Time   `json:"created"`
	Model   Model  `json:"model"`
	Choices []struct {
		Index   int         `json:"index"`
		Message ChatMessage `json:"message"`
		Status  string      `json:"status"`
	} `json:"choices"`
}

// Chat generate chat completions based on a conversation history.
func (cln *Client) Chat(ctx context.Context, model Model, messages []ChatMessage, maxTokens int, temperature float32) (Chat, error) {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	body := struct {
		Model       string        `json:"model"`
		Messages    []ChatMessage `json:"messages"`
		MaxTokens   int           `json:"max_tokens"`
		Temperature float32       `json:"temperature"`
	}{
		Model:       model.name,
		Messages:    messages,
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}

	var resp Chat
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Chat{}, err
	}

	return resp, nil
}

// =============================================================================

// ChatSSE represents the result for the chat completion call.
type ChatSSE struct {
	Chat
	Choices []struct {
		Index int `json:"index"`
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		Text         string  `json:"generated_text"`
		Probs        float32 `json:"logprobs"`
		FinishReason string  `json:"finish_reason"`
	} `json:"choices"`
}

// ChatSSE generate chat completions based on a conversation history.
func (cln *Client) ChatSSE(ctx context.Context, model string, input []ChatMessage, maxTokens int, temperature float32, ch chan ChatSSE) error {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	body := struct {
		Model       string        `json:"model"`
		Messages    []ChatMessage `json:"messages"`
		MaxTokens   int           `json:"max_tokens"`
		Temperature float32       `json:"temperature"`
		Stream      bool          `json:"stream"`
	}{
		Model:       model,
		Messages:    input,
		MaxTokens:   maxTokens,
		Temperature: temperature,
		Stream:      true,
	}

	sse := newSSEClient[ChatSSE](cln)

	if err := sse.do(ctx, http.MethodPost, url, body, ch); err != nil {
		return err
	}

	return nil
}
