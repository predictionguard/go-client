// Package client provides support to access the Prediction Guard API service.
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"path"
	"time"
)

// This provides a default client configuration and is set with reasonable
// defaults. Users can replace this client with application specific settings
// using the WithClient function at the time a Client is constructed.
var defaultClient = http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 15 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// Logger represents a function that will be called to add information
// to the user's application logs.
type Logger func(context.Context, string, ...any)

// Client represents a client that can talk to the PG API service.
type Client struct {
	log    Logger
	host   string
	apiKey string
	http   *http.Client
}

// New constructs a Client that can be used to talk with the PG API service.
func New(log Logger, host string, apiKey string, options ...func(cln *Client)) *Client {
	cln := Client{
		log:    log,
		host:   host,
		apiKey: apiKey,
		http:   &defaultClient,
	}

	for _, option := range options {
		option(&cln)
	}

	return &cln
}

// WithClient adds a custom client for processing requests. It's recommend
// to not use the default client and provide your own.
func WithClient(http *http.Client) func(cln *Client) {
	return func(cln *Client) {
		cln.http = http
	}
}

// =============================================================================

func (cln *Client) rawRequest(ctx context.Context, method string, endpoint string, body io.Reader, v any) error {
	var statusCode int

	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("parsing endpoint: %w", err)
	}
	base := path.Base(u.Path)

	cln.log(ctx, "go-client: rawRequest: started", "method", method, "call", base, "endpoint", endpoint)
	defer func() {
		cln.log(ctx, "go-client: rawRequest: completed", "status", statusCode)
	}()

	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return fmt.Errorf("create request error: %w", err)
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-api-key", cln.apiKey)

	resp, err := cln.http.Do(req)
	if err != nil {
		return fmt.Errorf("do: error: %w", err)
	}
	defer resp.Body.Close()

	// Assign for logging the status code at the end of the function call.
	statusCode = resp.StatusCode

	if statusCode == http.StatusNoContent {
		return nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("copy error: %w", err)
	}

	switch statusCode {
	case http.StatusNoContent:
		return nil

	case http.StatusOK:
		switch d := v.(type) {
		case *string:
			*d = string(data)

		default:
			if err := json.Unmarshal(data, v); err != nil {
				return fmt.Errorf("failed: response: %s, decoding error: %w ", string(data), err)
			}
		}

		return nil

	case http.StatusUnauthorized:
		var err Error
		if err := json.Unmarshal(data, &err); err != nil {
			return fmt.Errorf("failed: response: %s, decoding error: %w ", string(data), err)
		}
		return &err

	default:
		return fmt.Errorf("failed: response: %s", string(data))
	}
}
