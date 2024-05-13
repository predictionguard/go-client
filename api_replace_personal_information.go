package client

import (
	"context"
	"fmt"
	"net/http"
)

// ReplacePersonalInformation replaces personal information such as names,
// SSNs, and emails in a given text.
type ReplacePersonalInformation struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created Time   `json:"created"`
	Checks  []struct {
		Text   string `json:"new_prompt"`
		Index  int    `json:"index"`
		Status string `json:"status"`
	} `json:"checks"`
}

// ReplacePersonalInformation replaces personal information such as names, SSNs,
// and emails in a given text.
func (cln *Client) ReplacePersonalInformation(ctx context.Context, prompt string, method ReplaceMethod) (ReplacePersonalInformation, error) {
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

	var resp ReplacePersonalInformation
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return ReplacePersonalInformation{}, err
	}

	return resp, nil
}
