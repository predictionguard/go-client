package client

import (
	"context"
	"fmt"
	"net/http"
)

// InputExtension represents options for pre-processing.
type InputExtension struct {
	BlockPromptInjection bool
	PII                  PII
	PIIReplaceMethod     ReplaceMethod
}

// OutputExtension represents options for post-processing.
type OutputExtension struct {
	Factuality bool
	Toxicity   bool
}

// =============================================================================
// Basic Chat Completions

// ChatInputTypes defines behavior any chat input type must implement. The
// method doesn't need to do anything, it just needs to exist.
type ChatInputTypes interface {
	ChatInputType()
}

// ChatInputMessage represents a role and content related to a chat.
type ChatInputMessage struct {
	Role    Role
	Content string
}

// ChatInputMulti represents the full potential input options for chat.
type ChatInputMulti struct {
	Model           string
	Messages        []ChatInputMessage
	MaxTokens       *int
	Temperature     *float32
	TopP            *float64
	TopK            *int
	InputExtension  *InputExtension
	OutputExtension *OutputExtension
}

// ChatInputType implements the ChatInputTypes interface.
func (ChatInputMulti) ChatInputType() {}

// ChatInput represents the full potential input options for chat.
type ChatInput struct {
	Model           string
	Message         string
	MaxTokens       *int
	Temperature     *float32
	TopP            *float64
	TopK            *int
	InputExtension  *InputExtension
	OutputExtension *OutputExtension
}

// ChatInputType implements the ChatInputTypes interface.
func (ChatInput) ChatInputType() {}

// ChatMessage represents the role of the sender and the response.
type ChatMessage struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// ChatChoice represents a choice for the chat call.
type ChatChoice struct {
	Index   int         `json:"index"`
	Message ChatMessage `json:"message"`
}

// Chat represents the result for the chat call.
type Chat struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created Time         `json:"created"`
	Model   string       `json:"model"`
	Choices []ChatChoice `json:"choices"`
}

// Chat generate chat completions based on a conversation history.
func (cln *Client) Chat(ctx context.Context, input ChatInputTypes) (Chat, error) {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	// -------------------------------------------------------------------------

	type chatMessage struct {
		Role    Role   `json:"role"`
		Content string `json:"content"`
		Output  string `json:"output"`
	}

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

	var inputMulti ChatInputMulti

	switch v := input.(type) {
	case ChatInput:
		inputMulti = ChatInputMulti{
			Model: v.Model,
			Messages: []ChatInputMessage{
				{
					Role:    Roles.User,
					Content: v.Message,
				},
			},
			MaxTokens:       v.MaxTokens,
			Temperature:     v.Temperature,
			TopP:            v.TopP,
			TopK:            v.TopK,
			InputExtension:  v.InputExtension,
			OutputExtension: v.OutputExtension,
		}

	case ChatInputMulti:
		inputMulti = v
	}

	// -------------------------------------------------------------------------

	inputs := make([]chatMessage, len(inputMulti.Messages))
	for i, inp := range inputMulti.Messages {
		inputs[i] = chatMessage{
			Role:    inp.Role,
			Content: inp.Content,
		}
	}

	body := struct {
		Model           string           `json:"model"`
		Messages        []chatMessage    `json:"messages"`
		MaxTokens       *int             `json:"max_tokens,omitempty"`
		Temperature     *float32         `json:"temperature,omitempty"`
		TopP            *float64         `json:"top_p,omitempty"`
		TopK            *int             `json:"top_k,omitempty"`
		InputExtension  *inputExtension  `json:"input,omitempty"`
		OutputExtension *outputExtension `json:"output,omitempty"`
	}{
		Model:       inputMulti.Model,
		Messages:    inputs,
		MaxTokens:   inputMulti.MaxTokens,
		Temperature: inputMulti.Temperature,
		TopP:        inputMulti.TopP,
		TopK:        inputMulti.TopK,
	}

	// -------------------------------------------------------------------------

	if inputMulti.InputExtension != nil {
		if (inputMulti.InputExtension.BlockPromptInjection || inputMulti.InputExtension.PII != PII{} || inputMulti.InputExtension.PIIReplaceMethod != ReplaceMethod{}) {
			body.InputExtension = &inputExtension{
				BlockPromptInjection: inputMulti.InputExtension.BlockPromptInjection,
				PII:                  inputMulti.InputExtension.PII.value,
				PIIReplaceMethod:     inputMulti.InputExtension.PIIReplaceMethod,
			}
		}
	}

	if inputMulti.OutputExtension != nil {
		if inputMulti.OutputExtension.Factuality || inputMulti.OutputExtension.Toxicity {
			body.OutputExtension = &outputExtension{
				Factuality: inputMulti.OutputExtension.Factuality,
				Toxicity:   inputMulti.OutputExtension.Toxicity,
			}
		}
	}

	// -------------------------------------------------------------------------

	var resp Chat
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Chat{}, err
	}

	return resp, nil
}

// =============================================================================
// Streaming Chat Completions

// ChatSSEInput represents the full potential input options for SSE chat.
type ChatSSEInput struct {
	Model          string
	Messages       []ChatInputMessage
	MaxTokens      *int
	Temperature    *float32
	TopP           *float64
	TopK           *int
	InputExtension *InputExtension
}

// ChatSSEDelta represents content for the sse call.
type ChatSSEDelta struct {
	Content string `json:"content"`
}

// ChatSSEChoice represents a choice for the sse call.
type ChatSSEChoice struct {
	Index        int          `json:"index"`
	Delta        ChatSSEDelta `json:"delta"`
	Text         string       `json:"generated_text"`
	Probs        float32      `json:"logprobs"`
	FinishReason string       `json:"finish_reason"`
}

// ChatSSE represents the result for the sse call.
type ChatSSE struct {
	ID      string          `json:"id"`
	Object  string          `json:"object"`
	Created Time            `json:"created"`
	Model   string          `json:"model"`
	Choices []ChatSSEChoice `json:"choices"`
	Error   string          `json:"error"`
}

// ChatSSE generate chat completions based on a conversation history.
func (cln *Client) ChatSSE(ctx context.Context, input ChatSSEInput, ch chan ChatSSE) error {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	// -------------------------------------------------------------------------

	type chatInput struct {
		Role    Role   `json:"role"`
		Content string `json:"content"`
		Output  string `json:"output"`
	}

	type inputExtension struct {
		BlockPromptInjection bool          `json:"block_prompt_injection"`
		PII                  string        `json:"pii"`
		PIIReplaceMethod     ReplaceMethod `json:"pii_replace_method"`
	}

	// -------------------------------------------------------------------------

	messages := make([]chatInput, len(input.Messages))
	for i, inp := range input.Messages {
		messages[i] = chatInput{
			Role:    inp.Role,
			Content: inp.Content,
		}
	}

	body := struct {
		Model          string          `json:"model"`
		Messages       []chatInput     `json:"messages"`
		MaxTokens      *int            `json:"max_tokens,omitempty"`
		Temperature    *float32        `json:"temperature,omitempty"`
		TopP           *float64        `json:"top_p,omitempty"`
		TopK           *int            `json:"top_k,omitempty"`
		Stream         bool            `json:"stream"`
		InputExtension *inputExtension `json:"input,omitempty"`
	}{
		Model:       input.Model,
		Messages:    messages,
		MaxTokens:   input.MaxTokens,
		Temperature: input.Temperature,
		TopP:        input.TopP,
		TopK:        input.TopK,
		Stream:      true,
	}

	// -------------------------------------------------------------------------

	if input.InputExtension != nil {
		if (input.InputExtension.BlockPromptInjection || input.InputExtension.PII != PII{} || input.InputExtension.PIIReplaceMethod != ReplaceMethod{}) {
			body.InputExtension = &inputExtension{
				BlockPromptInjection: input.InputExtension.BlockPromptInjection,
				PII:                  input.InputExtension.PII.value,
				PIIReplaceMethod:     input.InputExtension.PIIReplaceMethod,
			}
		}
	}

	// -------------------------------------------------------------------------

	sse := newSSEClient[ChatSSE](cln)
	if err := sse.do(ctx, http.MethodPost, url, body, ch); err != nil {
		return err
	}

	return nil
}

// =============================================================================
// Vision Chat Completions

// ChatVisionInput represents the full potential input options for vision chat.
type ChatVisionInput struct {
	Model           string
	Role            Role
	Question        string
	Image           Base64Encoder
	MaxTokens       int
	Temperature     *float32
	TopP            *float64
	TopK            *int
	InputExtension  *InputExtension
	OutputExtension *OutputExtension
}

// ChatVisionMessage represents content for the vision call.
type ChatVisionMessage struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// ChatVisionChoice represents a choice for the vision call.
type ChatVisionChoice struct {
	Index   int               `json:"index"`
	Message ChatVisionMessage `json:"message"`
}

// ChatVision represents the result for the vision call.
type ChatVision struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created Time               `json:"created"`
	Model   string             `json:"model"`
	Choices []ChatVisionChoice `json:"choices"`
}

// ChatVision generate chat completions based on a question and an image.
func (cln *Client) ChatVision(ctx context.Context, input ChatVisionInput) (ChatVision, error) {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	base64, err := input.Image.EncodeBase64(ctx)
	if err != nil {
		return ChatVision{}, fmt.Errorf("base64: %w", err)
	}

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

	type content struct {
		Type     string `json:"type"`
		Text     string `json:"text,omitempty"`
		ImageURL struct {
			URL string `json:"url,omitempty"`
		} `json:"image_url,omitempty"`
	}

	type message struct {
		Role    Role      `json:"role"`
		Content []content `json:"content"`
	}

	// -------------------------------------------------------------------------

	body := struct {
		Model           string           `json:"model"`
		Messages        []message        `json:"messages"`
		MaxTokens       int              `json:"max_tokens"`
		Temperature     *float32         `json:"temperature,omitempty"`
		TopP            *float64         `json:"top_p,omitempty"`
		TopK            *int             `json:"top_k,omitempty"`
		InputExtension  *inputExtension  `json:"input,omitempty"`
		OutputExtension *outputExtension `json:"output,omitempty"`
	}{
		Model: input.Model,
		Messages: []message{
			{
				Role: input.Role,
				Content: []content{
					{
						Type: "text",
						Text: input.Question,
					},
					{
						Type: "image_url",
						ImageURL: struct {
							URL string `json:"url,omitempty"`
						}{
							URL: fmt.Sprintf("data:image/jpeg;base64,%s", base64),
						},
					},
				},
			},
		},
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
				PII:                  input.InputExtension.PII.value,
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

	var resp ChatVision
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return ChatVision{}, err
	}

	return resp, nil
}
