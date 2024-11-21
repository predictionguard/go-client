package client

import (
	"context"
	"fmt"
	"net/http"
)

// RerankInput represents input for reranking documents.
type RerankInput struct {
	Model           string   `json:"model"`
	Query           string   `json:"query"`
	Documents       []string `json:"documents"`
	ReturnDocuments bool     `json:"return_documents"`
}

// RerankResult represents a result for the chat call.
type RerankResult struct {
	Index          int     `json:"index"`
	RelevanceScore float64 `json:"relevance_score"`
	Text           string  `json:"text"`
}

// Rerank represents the result for the rerank call.
type Rerank struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created Time           `json:"created"`
	Model   string         `json:"model"`
	Results []RerankResult `json:"results"`
}

// Rerank sorts text inputs by semantic relevance to a specified query.
func (cln *Client) Rerank(ctx context.Context, input RerankInput) (Rerank, error) {
	url := fmt.Sprintf("%s/rerank", cln.host)

	body := RerankInput{
		Model:           input.Model,
		Query:           input.Query,
		Documents:       input.Documents,
		ReturnDocuments: input.ReturnDocuments,
	}

	var resp Rerank
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Rerank{}, err
	}

	return resp, nil
}
