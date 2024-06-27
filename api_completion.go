package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// Set of models supported by this API.
var completionModels = map[Model]bool{
	Models.DeepseekCoder67BInstruct: true,
	Models.Hermes2ProLlama38B:       true,
	Models.Hermes2ProMistral7B:      true,
	Models.NeuralChat7B:             true,
	Models.NousHermesLlama213B:      true,
}

// CompletionInput represents the full potential input options for completion.
type CompletionInput struct {
	Model       Model
	Prompt      string
	MaxTokens   int
	Temperature float32
	TopP        float64
	TopK        float64
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

	if !completionModels[input.Model] {
		return Completion{}, errors.New("model specified is not supported")
	}

	body := struct {
		Model       string  `json:"model"`
		Prompt      string  `json:"prompt"`
		MaxTokens   int     `json:"max_tokens"`
		Temperature float32 `json:"temperature"`
		TopP        float64 `json:"top_p"`
		TopK        float64 `json:"top_k"`
	}{
		Model:       input.Model.name,
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
