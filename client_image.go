package client

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type ImageFile struct {
	path   string
	base64 string
}

func NewImageFile(imagePath string) (ImageFile, error) {
	if _, err := os.Stat(imagePath); err != nil {
		return ImageFile{}, fmt.Errorf("file doesn't exist: %w", err)
	}

	img := ImageFile{
		path: imagePath,
	}

	return img, nil
}

func (img ImageFile) EncodeBase64(ctx context.Context) (string, error) {
	if img.base64 != "" {
		return img.base64, nil
	}

	data, err := os.ReadFile(img.path)
	if err != nil {
		return "", fmt.Errorf("readfile: %w", err)
	}

	img.base64 = b64.StdEncoding.EncodeToString(data)

	return img.base64, nil
}

// =============================================================================

type ImageNetwork struct {
	url    url.URL
	base64 string
}

func NewImageNetwork(imageURL string) (ImageNetwork, error) {
	url, err := url.Parse(imageURL)
	if err != nil {
		return ImageNetwork{}, fmt.Errorf("url doesn't parse: %w", err)
	}

	img := ImageNetwork{
		url: *url,
	}

	return img, nil
}

func (img ImageNetwork) EncodeBase64(ctx context.Context) (string, error) {
	if img.base64 != "" {
		return img.base64, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, img.url.String(), nil)
	if err != nil {
		return "", fmt.Errorf("create request error: %w", err)
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "image/*")

	resp, err := defaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("readall: %w", err)
	}

	img.base64 = b64.StdEncoding.EncodeToString(data)

	return img.base64, nil
}

// =============================================================================

type ImageBase64 struct {
	base64 string
}

func NewImageBase64(base64 string) ImageBase64 {
	return ImageBase64{
		base64: base64,
	}
}

func (img ImageBase64) EncodeBase64(ctx context.Context) (string, error) {
	return img.base64, nil
}
