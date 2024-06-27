package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// Set of models supported by these APIs.
var chatModels = map[Model]bool{
	Models.DeepseekCoder67BInstruct: true,
	Models.Hermes2ProLlama38B:       true,
	Models.Hermes2ProMistral7B:      true,
	Models.LLama3SqlCoder8b:         true,
	Models.Llava157BHF:              true,
	Models.NeuralChat7B:             true,
}

// ChatInput represents the full potential input options for chat.
type ChatInput struct {
	Model       Model
	Messages    []ChatInputMessage
	MaxTokens   int
	Temperature float32
	TopP        float64
	TopK        float64
	Options     *ChatInputOptions
}

// ChatInputMessage represents a role and content related to a chat.
type ChatInputMessage struct {
	Role    Role
	Content string
}

// ChatInputOptions represents options for post and preprocessing the input.
type ChatInputOptions struct {
	Factuality           bool
	Toxicity             bool
	BlockPromptInjection bool
	PII                  PII
	PIIReplaceMethod     ReplaceMethod
}

// ChatMessage represents the role of the sender and the response.
type ChatMessage struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
	Output  string `json:"output"`
}

// ChatChoice represents a choice for the chat call.
type ChatChoice struct {
	Index   int         `json:"index"`
	Message ChatMessage `json:"message"`
	Status  string      `json:"status"`
}

// Chat represents the result for the chat call.
type Chat struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created Time         `json:"created"`
	Model   Model        `json:"model"`
	Choices []ChatChoice `json:"choices"`
}

// Chat generate chat completions based on a conversation history.
func (cln *Client) Chat(ctx context.Context, input ChatInput) (Chat, error) {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	type chatMessage struct {
		Role    Role   `json:"role"`
		Content string `json:"content"`
		Output  string `json:"output"`
	}

	type chatOutputOption struct {
		Factuality bool `json:"factuality"`
		Toxicity   bool `json:"toxicity"`
	}

	type chatInputOption struct {
		BlockPromptInjection bool          `json:"block_prompt_injection"`
		PII                  string        `json:"pii"`
		PIIReplaceMethod     ReplaceMethod `json:"pii_replace_method"`
	}

	if !chatModels[input.Model] {
		return Chat{}, errors.New("model specified is not supported")
	}

	inputs := make([]chatMessage, len(input.Messages))
	for i, inp := range input.Messages {
		inputs[i] = chatMessage{
			Role:    inp.Role,
			Content: inp.Content,
		}
	}

	body := struct {
		Model       string            `json:"model"`
		Messages    []chatMessage     `json:"messages"`
		MaxTokens   int               `json:"max_tokens"`
		Temperature float32           `json:"temperature"`
		TopP        float64           `json:"top_p"`
		TopK        float64           `json:"top_k"`
		Output      *chatOutputOption `json:"output,omitempty"`
		Input       *chatInputOption  `json:"input,omitempty"`
	}{
		Model:       input.Model.name,
		Messages:    inputs,
		MaxTokens:   input.MaxTokens,
		Temperature: input.Temperature,
		TopP:        input.TopP,
		TopK:        input.TopK,
	}

	if input.Options != nil {
		if input.Options.Factuality || input.Options.Toxicity {
			body.Output = &chatOutputOption{
				Factuality: input.Options.Factuality,
				Toxicity:   input.Options.Toxicity,
			}
		}

		if (input.Options.BlockPromptInjection || input.Options.PII != PII{} || input.Options.PIIReplaceMethod != ReplaceMethod{}) {
			body.Input = &chatInputOption{
				BlockPromptInjection: input.Options.BlockPromptInjection,
				PII:                  input.Options.PII.name,
				PIIReplaceMethod:     input.Options.PIIReplaceMethod,
			}
		}
	}

	var resp Chat
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Chat{}, err
	}

	return resp, nil
}

// =============================================================================

// ChatSSEInput represents the full potential input options for SSE chat.
type ChatSSEInput struct {
	Model       Model
	Messages    []ChatInputMessage
	MaxTokens   int
	Temperature float32
	TopP        float64
	TopK        float64
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
	Model   Model           `json:"model"`
	Choices []ChatSSEChoice `json:"choices"`
}

// ChatSSE generate chat completions based on a conversation history.
func (cln *Client) ChatSSE(ctx context.Context, input ChatSSEInput, ch chan ChatSSE) error {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	type chatInput struct {
		Role    Role   `json:"role"`
		Content string `json:"content"`
		Output  string `json:"output"`
	}

	if !chatModels[input.Model] {
		return errors.New("model specified is not supported")
	}

	messages := make([]chatInput, len(input.Messages))
	for i, inp := range input.Messages {
		messages[i] = chatInput{
			Role:    inp.Role,
			Content: inp.Content,
		}
	}

	body := struct {
		Model       string      `json:"model"`
		Messages    []chatInput `json:"messages"`
		MaxTokens   int         `json:"max_tokens"`
		Temperature float32     `json:"temperature"`
		TopP        float64     `json:"top_p"`
		TopK        float64     `json:"top_k"`
		Stream      bool        `json:"stream"`
	}{
		Model:       input.Model.name,
		Messages:    messages,
		MaxTokens:   input.MaxTokens,
		Temperature: input.Temperature,
		TopP:        input.TopP,
		TopK:        input.TopK,
		Stream:      true,
	}

	sse := newSSEClient[ChatSSE](cln)

	if err := sse.do(ctx, http.MethodPost, url, body, ch); err != nil {
		return err
	}

	return nil
}

// =============================================================================

// ChatVisionInput represents the full potential input options for vision chat.
type ChatVisionInput struct {
	Role        Role
	Question    string
	Image       Base64Encoder
	MaxTokens   int
	Temperature float32
	TopP        float64
	TopK        float64
}

// ChatVisionMessage represents content for the vision call.
type ChatVisionMessage struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
	Output  string `json:"output"`
}

// ChatVisionChoice represents a choice for the vision call.
type ChatVisionChoice struct {
	Index   int               `json:"index"`
	Message ChatVisionMessage `json:"message"`
	Status  string            `json:"status"`
}

// ChatVision represents the result for the vision call.
type ChatVision struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created Time               `json:"created"`
	Model   Model              `json:"model"`
	Choices []ChatVisionChoice `json:"choices"`
}

// ChatVision generate chat completions based on a question and an image.
func (cln *Client) ChatVision(ctx context.Context, input ChatVisionInput) (ChatVision, error) {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	base64, err := input.Image.EncodeBase64(ctx)
	if err != nil {
		return ChatVision{}, fmt.Errorf("base64: %w", err)
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

	body := struct {
		Model       string    `json:"model"`
		Messages    []message `json:"messages"`
		MaxTokens   int       `json:"max_tokens"`
		Temperature float32   `json:"temperature"`
		TopP        float64   `json:"top_p"`
		TopK        float64   `json:"top_k"`
	}{
		Model: Models.Llava157BHF.name,
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

	var resp ChatVision
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return ChatVision{}, err
	}

	return resp, nil
}
