// Package client provides support to access the Prediction Guard API service.
package client

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

// TODO: Maintain this version when a new tag is created.
const version = "v0.14.0"

// ErrUnauthorized represent a situation where authentication fails.
var ErrUnauthorized = errors.New("api understands the request but refuses to authorize it")

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

// =============================================================================

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
		host:   strings.TrimLeft(host, "/"),
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

func (cln *Client) do(ctx context.Context, method string, endpoint string, body any, v any) error {
	resp, err := do(ctx, cln, method, endpoint, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("client: copy error: %w", err)
	}

	switch d := v.(type) {
	case *string:
		*d = string(data)

	default:
		if err := json.Unmarshal(data, v); err != nil {
			return fmt.Errorf("client: response: %s, decoding error: %w ", string(data), err)
		}
	}

	return nil
}

// =============================================================================

type sseClient[T any] struct {
	*Client
}

func newSSEClient[T any](cln *Client) sseClient[T] {
	return sseClient[T]{
		Client: cln,
	}
}

func (cln *sseClient[T]) do(ctx context.Context, method string, endpoint string, body any, ch chan T) error {
	resp, err := do(ctx, cln.Client, method, endpoint, body)
	if err != nil {
		return err
	}

	go func(ctx context.Context) {
		ticker := time.NewTicker(5 * time.Second)

		defer func() {
			resp.Body.Close()
			ticker.Stop()
			close(ch)
		}()

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()

			if line == "" || line == "data: [DONE]" {
				continue
			}

			var v T
			if err := json.Unmarshal([]byte(line[6:]), &v); err != nil {
				cln.log(ctx, "sseclient: rawRequest:", "Unmarshal", err)
				break
			}

			select {
			case ch <- v:

			case <-ctx.Done():
				cln.log(ctx, "sseclient: rawRequest:", "Context", err)

			case <-ticker.C:
				cln.log(ctx, "sseclient: rawRequest:", "WARNING", "timeout waiting for a receiver")
			}

			if ctx.Err() != nil {
				break
			}

			ticker.Reset(time.Second)
		}

	}(ctx)

	return nil
}

// =============================================================================

func do(ctx context.Context, cln *Client, method string, endpoint string, body any) (*http.Response, error) {
	var statusCode int

	cln.log(ctx, "do: rawRequest: started", "method", method, "endpoint", endpoint)
	defer func() {
		cln.log(ctx, "do: rawRequest: completed", "status", statusCode)
	}()

	var b bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&b).Encode(body); err != nil {
			return nil, fmt.Errorf("encoding: error: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, &b)
	if err != nil {
		return nil, fmt.Errorf("create request error: %w", err)
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("Prediction Guard Go Client: %s", version))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cln.apiKey))

	resp, err := cln.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do: error: %w", err)
	}

	// Assign for logging the status code at the end of the function call.
	statusCode = resp.StatusCode

	if statusCode != http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("readall: error: %w", err)
		}

		switch statusCode {
		case http.StatusForbidden:
			return nil, ErrUnauthorized

		default:
			var err Error
			if err := json.Unmarshal(data, &err); err != nil {
				return nil, fmt.Errorf("decoding: response: %s, error: %w ", string(data), err)
			}

			return nil, fmt.Errorf("error: response: %s", err.Message)
		}
	}

	return resp, nil
}
