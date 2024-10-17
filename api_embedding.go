package client

import (
	"context"
	"fmt"
	"net/http"
)

// EmbeddingInput represents the input to generate embeddings.
type EmbeddingInput struct {
	Text  string
	Image Base64Encoder
}

// EmbeddingData represents the vector data points.
type EmbeddingData struct {
	Index     int       `json:"index"`
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
}

// Embedding represents the result for the embedding call.
type Embedding struct {
	ID      string          `json:"id"`
	Object  string          `json:"object"`
	Created Time            `json:"created"`
	Model   string          `json:"model"`
	Data    []EmbeddingData `json:"data"`
}

// Embedding converts text, text + image, and image to a numerical representation
// that is useful for search and retrieval. When you have both text and image,
// the use case would be like a video frame plus the transcription or an image
// plus a caption. The response should include the output vector.
func (cln *Client) Embedding(ctx context.Context, model string, input []EmbeddingInput) (Embedding, error) {
	return cln.embedding(ctx, model, input, nil, nil)
}

// EmbeddingWithTruncate behaves like Embedding but provides an option to set
// a truncation direction for models that support truncation.
func (cln *Client) EmbeddingWithTruncate(ctx context.Context, model string, input []EmbeddingInput, direction Direction) (Embedding, error) {
	truncate := true
	return cln.embedding(ctx, model, input, &truncate, &direction)
}

func (cln *Client) embedding(ctx context.Context, model string, input []EmbeddingInput, truncate *bool, truncateDir *Direction) (Embedding, error) {
	url := fmt.Sprintf("%s/embeddings", cln.host)

	type embeddingInput struct {
		Text  string `json:"text"`
		Image string `json:"image"`
	}

	data := make([]embeddingInput, len(input))
	for i, inp := range input {
		var base64 string
		if inp.Image != nil {
			var err error
			base64, err = inp.Image.EncodeBase64(ctx)
			if err != nil {
				return Embedding{}, fmt.Errorf("base64: %w", err)
			}
		}

		data[i] = embeddingInput{
			Text:  inp.Text,
			Image: base64,
		}
	}

	var direction *string
	if truncateDir != nil {
		direction = &truncateDir.value
	}

	body := struct {
		Model       string           `json:"model"`
		Truncate    *bool            `json:"truncate,omitempty"`
		TruncateDir *string          `json:"truncation_direction,omitempty"`
		Input       []embeddingInput `json:"input"`
	}{
		Model:       model,
		Truncate:    truncate,
		TruncateDir: direction,
		Input:       data,
	}

	var resp Embedding
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Embedding{}, err
	}

	return resp, nil
}
