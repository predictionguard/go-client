package client

import (
	"context"
	"fmt"
	"net/http"
)

// Injection represents the result for the detect injection call.
type Injection struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created Time   `json:"created"`
	Checks  []struct {
		Probability float64 `json:"probability"`
		Index       int     `json:"index"`
		Status      string  `json:"status"`
	} `json:"checks"`
}

// Injection detects potential prompt injection attacks in a given prompt.
func (cln *Client) Injection(ctx context.Context, prompt string) (Injection, error) {
	url := fmt.Sprintf("%s/injection", cln.host)

	body := struct {
		Prompt string `json:"prompt"`
		Detect bool   `json:"detect"`
	}{
		Prompt: prompt,
		Detect: true,
	}

	var resp Injection
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Injection{}, err
	}

	return resp, nil
}
