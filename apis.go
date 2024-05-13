package client

import (
	"context"
	"fmt"
	"net/http"
)

// HealthCheck validates the PG API Service is available.
func (cln *Client) HealthCheck(ctx context.Context) (string, error) {
	var resp string
	if err := cln.do(ctx, http.MethodGet, cln.host, nil, &resp); err != nil {
		return "", err
	}

	return resp, nil
}

// =============================================================================

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

// =============================================================================

// Completions retrieve text completions based on the provided input.
func (cln *Client) Completions(ctx context.Context, model Model, text string, maxTokens int, temperature float32) (Completion, error) {
	url := fmt.Sprintf("%s/completions", cln.host)

	body := struct {
		Model       string  `json:"model"`
		Prompt      string  `json:"prompt"`
		MaxTokens   int     `json:"max_tokens"`
		Temperature float32 `json:"temperature"`
	}{
		Model:       model.name,
		Prompt:      text,
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}

	var resp Completion
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Completion{}, err
	}

	return resp, nil
}

// =============================================================================

// Factuality checks the factuality of a given text compared to a reference.
func (cln *Client) Factuality(ctx context.Context, reference string, text string) (Factuality, error) {
	url := fmt.Sprintf("%s/factuality", cln.host)

	body := struct {
		Reference string `json:"reference"`
		Text      string `json:"text"`
	}{
		Reference: reference,
		Text:      text,
	}

	var resp Factuality
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Factuality{}, err
	}

	return resp, nil
}

// =============================================================================

// Translate converts text from one language to another.
func (cln *Client) Translate(ctx context.Context, text string, source Language, target Language) (Translate, error) {
	url := fmt.Sprintf("%s/translate", cln.host)

	body := struct {
		Text   string   `json:"text"`
		Source Language `json:"source_lang"`
		Target Language `json:"target_lang"`
	}{
		Text:   text,
		Source: source,
		Target: target,
	}

	var resp Translate
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Translate{}, err
	}

	return resp, nil
}

// =============================================================================

// ReplacePersonalInformation replaces personal information such as names, SSNs,
// and emails in a given text.
func (cln *Client) ReplacePersonalInformation(ctx context.Context, text string, method ReplaceMethod) (ReplacePersonalInformation, error) {
	url := fmt.Sprintf("%s/PII", cln.host)

	body := struct {
		Prompt  string        `json:"prompt"`
		Replace bool          `json:"replace"`
		Method  ReplaceMethod `json:"replace_method"`
	}{
		Prompt:  text,
		Replace: true,
		Method:  method,
	}

	var resp ReplacePersonalInformation
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return ReplacePersonalInformation{}, err
	}

	return resp, nil
}
