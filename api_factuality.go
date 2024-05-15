package client

import (
	"context"
	"fmt"
	"net/http"
)

// FactualityCheck represents the result for the factuality call.
type FactualityCheck struct {
	Score  float64 `json:"score"`
	Index  int     `json:"index"`
	Status string  `json:"status"`
}

// Factuality represents the result for the factuality call.
type Factuality struct {
	ID      string            `json:"id"`
	Object  string            `json:"object"`
	Created Time              `json:"created"`
	Checks  []FactualityCheck `json:"checks"`
}

// Factuality checks the factuality of a given text compared to a reference.
func (cln *Client) Factuality(ctx context.Context, reference string, text string) (Factuality, error) {
	url := fmt.Sprintf("%s/factuality", cln.host)

	body := struct {
		Reference string `json:"reference"`
		Text      string `json:"text"`
	}{
		Reference: reference,
		Text:      text,
	}

	var resp Factuality
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Factuality{}, err
	}

	return resp, nil
}
