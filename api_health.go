package client

import (
	"context"
	"net/http"
)

// HealthCheck validates the PG API Service is available.
func (cln *Client) HealthCheck(ctx context.Context) (string, error) {
	var resp string
	if err := cln.do(ctx, http.MethodGet, cln.host, nil, &resp); err != nil {
		return "", err
	}

	return resp, nil
}
