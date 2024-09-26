package client

import (
	"context"
	"fmt"
	"net/http"
)

// CompletionInput represents the full potential input options for completion.
type CompletionInput struct {
	Model           string
	Prompt          string
	MaxTokens       int
	Temperature     *float32
	TopP            *float64
	TopK            *int
	InputExtension  *InputExtension
	OutputExtension *OutputExtension
}

// CompletionChoice represents a choice for the completion call.
type CompletionChoice struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
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

	// -------------------------------------------------------------------------

	type inputExtension struct {
		BlockPromptInjection bool          `json:"block_prompt_injection"`
		PII                  string        `json:"pii"`
		PIIReplaceMethod     ReplaceMethod `json:"pii_replace_method"`
	}

	type outputExtension struct {
		Factuality bool `json:"factuality"`
		Toxicity   bool `json:"toxicity"`
	}

	// -------------------------------------------------------------------------

	body := struct {
		Model           string           `json:"model"`
		Prompt          string           `json:"prompt"`
		MaxTokens       int              `json:"max_tokens"`
		Temperature     *float32         `json:"temperature,omitempty"`
		TopP            *float64         `json:"top_p,omitempty"`
		TopK            *int             `json:"top_k,omitempty"`
		InputExtension  *inputExtension  `json:"input,omitempty"`
		OutputExtension *outputExtension `json:"output,omitempty"`
	}{
		Model:       input.Model,
		Prompt:      input.Prompt,
		MaxTokens:   input.MaxTokens,
		Temperature: input.Temperature,
		TopP:        input.TopP,
		TopK:        input.TopK,
	}

	// -------------------------------------------------------------------------

	if input.InputExtension != nil {
		if (input.InputExtension.BlockPromptInjection || input.InputExtension.PII != PII{} || input.InputExtension.PIIReplaceMethod != ReplaceMethod{}) {
			body.InputExtension = &inputExtension{
				BlockPromptInjection: input.InputExtension.BlockPromptInjection,
				PII:                  input.InputExtension.PII.name,
				PIIReplaceMethod:     input.InputExtension.PIIReplaceMethod,
			}
		}
	}

	if input.OutputExtension != nil {
		if input.OutputExtension.Factuality || input.OutputExtension.Toxicity {
			body.OutputExtension = &outputExtension{
				Factuality: input.OutputExtension.Factuality,
				Toxicity:   input.OutputExtension.Toxicity,
			}
		}
	}

	// -------------------------------------------------------------------------

	var resp Completion
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Completion{}, err
	}

	return resp, nil
}
