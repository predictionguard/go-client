package client

import (
	"context"
	"fmt"
	"net/http"
)

// ReplacePI represents the result for the pii call.
type ReplacePI struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created Time   `json:"created"`
	Checks  []struct {
		Text   string `json:"new_prompt"`
		Index  int    `json:"index"`
		Status string `json:"status"`
	} `json:"checks"`
}

// ReplacePI replaces personal information such as names, SSNs, and emails in a
// given text.
func (cln *Client) ReplacePI(ctx context.Context, prompt string, method ReplaceMethod) (ReplacePI, error) {
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

	var resp ReplacePI
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return ReplacePI{}, err
	}

	return resp, nil
}
