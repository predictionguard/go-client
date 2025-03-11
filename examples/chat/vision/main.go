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

	d := client.D{
		"model": "llava-1.5-7b-hf",
		"messages": []client.D{
			{
				"role": client.Roles.User,
				"content": []client.D{
					{
						"type": "text",
						"text": "Is this a picture of a rose?",
					},
					{
						"type": "image_url",
						"image_url": client.D{
							"url": fmt.Sprintf("data:image/png;base64,%s", base64),
						},
					},
				},
			},
		},
		"max_tokens":  1000,
		"temperature": 0.1,
		"top_p":       0.1,
		"top_k":       50,
		"input": client.D{
			"pii":                client.PIIs.Replace,
			"pii_replace_method": client.ReplaceMethods.Random,
		},
		"output": client.D{
			"factuality": false,
			"toxicity":   true,
		},
	}

	// -------------------------------------------------------------------------

	const url = "https://api.predictionguard.com/chat/completions"

	var resp client.ChatVision
	if err := cln.Do(ctx, http.MethodPost, url, d, &resp); err != nil {
		return fmt.Errorf("do: %w", err)
	}

	for i, choice := range resp.Choices {
		fmt.Printf("choice %d: %s\n", i, choice.Message.Content)
	}

	return nil
}
