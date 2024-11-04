package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ModelCapabilities struct {
	ChatCompletion     bool `json:"chat_completion"`
	ChatWithImage      bool `json:"chat_with_image"`
	Completion         bool `json:"completion"`
	Embedding          bool `json:"embedding"`
	EmbeddingWithImage bool `json:"embedding_with_image"`
	Tokenize           bool `json:"tokenize"`
}

type ModelData struct {
	ID               string            `json:"id"`
	Object           string            `json:"object"`
	Created          time.Time         `json:"created"`
	OwnedBy          string            `json:"owned_by"`
	Description      string            `json:"description"`
	MaxContextLength int               `json:"max_context_length"`
	PromptFormat     string            `json:"prompt_format"`
	Capabilities     ModelCapabilities `json:"capabilities"`
}

type ModelResponse struct {
	Object string      `json:"object"`
	Data   []ModelData `json:"data"`
}

// Capability returns the set of models for the specified capability.
func (cln *Client) Capability(ctx context.Context, capability Capability) (ModelResponse, error) {
	url := fmt.Sprintf("%s/models/%s", cln.host, capability)

	var resp ModelResponse
	if err := cln.do(ctx, http.MethodGet, url, nil, &resp); err != nil {
		return ModelResponse{}, err
	}

	return resp, nil
}
