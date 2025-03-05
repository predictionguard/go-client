package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/predictionguard/go-client/v2"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger := func(ctx context.Context, msg string, v ...any) {
		s := fmt.Sprintf("msg: %s", msg)
		for i := 0; i < len(v); i = i + 2 {
			s = s + fmt.Sprintf(", %s: %v", v[i], v[i+1])
		}
		log.Println(s)
	}

	cln := client.New(logger, os.Getenv("PREDICTIONGUARD_API_KEY"))

	// -------------------------------------------------------------------------

	image, err := client.NewImageNetwork("https://static.wixstatic.com/media/f54603_b7882b876e2b47d3a38843a58a9829f1~mv2.png")
	if err != nil {
		return fmt.Errorf("newimage: %w", err)
	}

	base64, err := image.EncodeBase64(ctx)
	if err != nil {
		return fmt.Errorf("base64: %w", err)
	}

	// -------------------------------------------------------------------------

	d := client.D{
		"model":              "bridgetower-large-itm-mlm-itc",
		"truncate":           true,
		"truncate_direction": client.Directions.Right,
		"input": []client.D{
			{
				"text":  "A picture of a rose",
				"image": base64,
			},
		},
	}

	// -------------------------------------------------------------------------

	const url = "https://api.predictionguard.com/embeddings"

	var resp client.Embedding
	if err := cln.Do(ctx, http.MethodPost, url, d, &resp); err != nil {
		return fmt.Errorf("do: %w", err)
	}

	for _, data := range resp.Data {
		fmt.Print(data.Embedding)
	}

	return nil
}
