package client

import (
	"context"
	"fmt"
	"net/http"
)

// Capability returns the set of models for the specified capability.
func (cln *Client) Capability(ctx context.Context, capability Capability) ([]string, error) {
	var url string

	switch capability {
	case Capabilities.ChatCompletion:
		url = fmt.Sprintf("%s/chat/completions", cln.host)

	case Capabilities.ChatCompletionVision:
		url = fmt.Sprintf("%s/chat/completions/vision", cln.host)

	case Capabilities.Completion:
		url = fmt.Sprintf("%s/completions", cln.host)

	case Capabilities.Embedding:
		url = fmt.Sprintf("%s/embeddings", cln.host)
	}

	var resp []string
	if err := cln.do(ctx, http.MethodGet, url, nil, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
