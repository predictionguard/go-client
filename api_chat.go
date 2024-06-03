package client

import (
	"context"
	"fmt"
	"net/http"
)

// ChatInput represents a role and content related to a chat.
type ChatInput struct {
	Role    Role
	Content string
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
func (cln *Client) Chat(ctx context.Context, model Model, input []ChatInput, maxTokens int, temperature float32) (Chat, error) {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	type chatInput struct {
		Role    Role   `json:"role"`
		Content string `json:"content"`
		Output  string `json:"output"`
	}

	inputs := make([]chatInput, len(input))
	for i, inp := range input {
		inputs[i] = chatInput{
			Role:    inp.Role,
			Content: inp.Content,
		}
	}

	body := struct {
		Model       string      `json:"model"`
		Messages    []chatInput `json:"messages"`
		MaxTokens   int         `json:"max_tokens"`
		Temperature float32     `json:"temperature"`
	}{
		Model:       model.name,
		Messages:    inputs,
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
func (cln *Client) ChatSSE(ctx context.Context, model Model, input []ChatInput, maxTokens int, temperature float32, ch chan ChatSSE) error {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	type chatInput struct {
		Role    Role   `json:"role"`
		Content string `json:"content"`
		Output  string `json:"output"`
	}

	inputs := make([]chatInput, len(input))
	for i, inp := range input {
		inputs[i] = chatInput{
			Role:    inp.Role,
			Content: inp.Content,
		}
	}

	body := struct {
		Model       string      `json:"model"`
		Messages    []chatInput `json:"messages"`
		MaxTokens   int         `json:"max_tokens"`
		Temperature float32     `json:"temperature"`
		Stream      bool        `json:"stream"`
	}{
		Model:       model.name,
		Messages:    inputs,
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
func (cln *Client) ChatVision(ctx context.Context, role Role, question string, image Base64Encoder, maxTokens int, temperature float32) (ChatVision, error) {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	base64, err := image.EncodeBase64(ctx)
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
	}{
		Model: Models.Llava157BHF.name,
		Messages: []message{
			{
				Role: role,
				Content: []content{
					{
						Type: "text",
						Text: question,
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
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}

	var resp ChatVision
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return ChatVision{}, err
	}

	return resp, nil
}
