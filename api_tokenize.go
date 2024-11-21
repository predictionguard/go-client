package client

import (
	"context"
	"fmt"
	"net/http"
)

// TokenizeInput represents input for tokenizing text.
type TokenizeInput struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

// TokenData represents a single token of information.
type TokenData struct {
	ID    int    `json:"id"`
	Start int    `json:"start"`
	Stop  int    `json:"stop"`
	Text  string `json:"text"`
}

// Tokenize represents the result for the toxicity call.
type Tokenize struct {
	ID      string      `json:"id"`
	Object  string      `json:"object"`
	Created Time        `json:"created"`
	Data    []TokenData `json:"data"`
}

// Tokenize provides the set of tokens the model server calculates.
func (cln *Client) Tokenize(ctx context.Context, input TokenizeInput) (Tokenize, error) {
	url := fmt.Sprintf("%s/tokenize", cln.host)

	body := TokenizeInput{
		Model: input.Model,
		Input: input.Input,
	}

	var resp Tokenize
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Tokenize{}, err
	}

	return resp, nil
}
