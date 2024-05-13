package client

import (
	"context"
	"fmt"
	"net/http"
)

// Toxicity represents the result for the toxicity call.
type Toxicity struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created Time   `json:"created"`
	Checks  []struct {
		Score  float64 `json:"score"`
		Index  int     `json:"index"`
		Status string  `json:"status"`
	} `json:"checks"`
}

// Completions retrieve text completions based on the provided input.
func (cln *Client) Toxicity(ctx context.Context, text string) (Toxicity, error) {
	url := fmt.Sprintf("%s/toxicity", cln.host)

	body := struct {
		Text string `json:"text"`
	}{
		Text: text,
	}

	var resp Toxicity
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Toxicity{}, err
	}

	return resp, nil
}
