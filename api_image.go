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

// ImageFile represents image data that will be read from a file.
type ImageFile struct {
	path   string
	base64 string
}

// NewImageFile constructs a ImageFile that can read an image from disk.
func NewImageFile(imagePath string) (ImageFile, error) {
	if _, err := os.Stat(imagePath); err != nil {
		return ImageFile{}, fmt.Errorf("file doesn't exist: %w", err)
	}

	img := ImageFile{
		path: imagePath,
	}

	return img, nil
}

// EncodeBase64 reads the specified image from disk and converts the image
// to a base64 string.
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

// ImageNetwork represents image data that will be read from the network.
type ImageNetwork struct {
	url    url.URL
	base64 string
}

// NewImageNetwork constructs a ImageNetwork that can read an image from the network.
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

// EncodeBase64 reads the specified image from the network and converts the
// image to a base64 string.
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

// ImageBase64 represents image data that is already in base64.
type ImageBase64 struct {
	base64 string
}

// NewImageBase64 constructs a ImageBase64 with an encoded image.
func NewImageBase64(base64 string) ImageBase64 {
	return ImageBase64{
		base64: base64,
	}
}

// EncodeBase64 returns the base64 image provided during construction.
func (img ImageBase64) EncodeBase64(ctx context.Context) (string, error) {
	return img.base64, nil
}
