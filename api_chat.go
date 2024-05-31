package client

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// ChatMessage represents the role of the sender and the content to process.
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
func (cln *Client) Chat(ctx context.Context, model Model, input []ChatMessage, maxTokens int, temperature float32) (Chat, error) {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	body := struct {
		Model       string        `json:"model"`
		Messages    []ChatMessage `json:"messages"`
		MaxTokens   int           `json:"max_tokens"`
		Temperature float32       `json:"temperature"`
	}{
		Model:       model.name,
		Messages:    input,
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
func (cln *Client) ChatSSE(ctx context.Context, model Model, input []ChatMessage, maxTokens int, temperature float32, ch chan ChatSSE) error {
	url := fmt.Sprintf("%s/chat/completions", cln.host)

	body := struct {
		Model       string        `json:"model"`
		Messages    []ChatMessage `json:"messages"`
		MaxTokens   int           `json:"max_tokens"`
		Temperature float32       `json:"temperature"`
		Stream      bool          `json:"stream"`
	}{
		Model:       model.name,
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

// Base64Encoder defines a method that can read a data source and returns a
// base64 encoded string.
type Base64Encoder interface {
	EncodeBase64(ctx context.Context) (string, error)
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
func (cln *Client) ChatVision(ctx context.Context, role Role, prompt string, image Base64Encoder, maxTokens int, temperature float32) (ChatVision, error) {
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
						Text: prompt,
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

// =============================================================================

// ImageFile represents image data that will be read from a file.
type ImageFile struct {
	path   string
	base64 string
}

// NewImageFile constructs a ImageFile that can read an image from disk.
func NewImageFile(imagePath string) (ImageFile, error) {
	if _, err := os.Stat(imagePath); err != nil {
		return ImageFile{}, fmt.Errorf("file doesn't exist: %w", err)
	}

	img := ImageFile{
		path: imagePath,
	}

	return img, nil
}

// EncodeBase64 reads the specified image from disk and converts the image
// to a base64 string.
func (img ImageFile) EncodeBase64(ctx context.Context) (string, error) {
	if img.base64 != "" {
		return img.base64, nil
	}

	data, err := os.ReadFile(img.path)
	if err != nil {
		return "", fmt.Errorf("readfile: %w", err)
	}

	img.base64 = b64.StdEncoding.EncodeToString(data)

	return img.base64, nil
}

// =============================================================================

// ImageNetwork represents image data that will be read from the network.
type ImageNetwork struct {
	url    url.URL
	base64 string
}

// NewImageNetwork constructs a ImageNetwork that can read an image from the network.
func NewImageNetwork(imageURL string) (ImageNetwork, error) {
	url, err := url.Parse(imageURL)
	if err != nil {
		return ImageNetwork{}, fmt.Errorf("url doesn't parse: %w", err)
	}

	img := ImageNetwork{
		url: *url,
	}

	return img, nil
}

// EncodeBase64 reads the specified image from the network and converts the
// image to a base64 string.
func (img ImageNetwork) EncodeBase64(ctx context.Context) (string, error) {
	if img.base64 != "" {
		return img.base64, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, img.url.String(), nil)
	if err != nil {
		return "", fmt.Errorf("create request error: %w", err)
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "image/*")

	resp, err := defaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("readall: %w", err)
	}

	img.base64 = b64.StdEncoding.EncodeToString(data)

	return img.base64, nil
}
