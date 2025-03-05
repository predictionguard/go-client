package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/predictionguard/go-client"
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

	d := client.D{
		"model": "neural-chat-7b-v3-3",
		"messages": []client.D{
			{
				"role":    client.Roles.User,
				"content": "How do you feel about the world in general",
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
			"factuality": true,
			"toxicity":   true,
		},
	}

	// -------------------------------------------------------------------------

	const url = "https://api.predictionguard.com/chat/completions"

	var resp client.Chat
	if err := cln.Do(ctx, http.MethodPost, url, d, &resp); err != nil {
		return fmt.Errorf("do: %w", err)
	}

	fmt.Println(resp.Choices[0].Message)

	return nil
}
