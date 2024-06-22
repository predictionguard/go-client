package client

import (
	"context"
	"fmt"
	"net/http"
)

// ReplacePIICheck represents the result for the pii call.
type ReplacePIICheck struct {
	NewPrompt string `json:"new_prompt"`
	Index     int    `json:"index"`
	Status    string `json:"status"`
}

// ReplacePII represents the result for the pii call.
type ReplacePII struct {
	ID      string            `json:"id"`
	Object  string            `json:"object"`
	Created Time              `json:"created"`
	Checks  []ReplacePIICheck `json:"checks"`
}

// ReplacePII replaces personal information such as names, SSNs, and emails in a
// given text.
func (cln *Client) ReplacePII(ctx context.Context, prompt string, method ReplaceMethod) (ReplacePII, error) {
	url := fmt.Sprintf("%s/PII", cln.host)

	body := struct {
		Prompt  string        `json:"prompt"`
		Replace bool          `json:"replace"`
		Method  ReplaceMethod `json:"replace_method"`
	}{
		Prompt:  prompt,
		Replace: true,
		Method:  method,
	}

	var resp ReplacePII
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return ReplacePII{}, err
	}

	return resp, nil
}
