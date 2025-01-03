package main

import (
	"context"
	"fmt"
	"log"
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
	host := "https://api.predictionguard.com"
	apiKey := os.Getenv("PREDICTIONGUARD_API_KEY")

	logger := func(ctx context.Context, msg string, v ...any) {
		s := fmt.Sprintf("msg: %s", msg)
		for i := 0; i < len(v); i = i + 2 {
			s = s + fmt.Sprintf(", %s: %v", v[i], v[i+1])
		}
		log.Println(s)
	}

	cln := client.New(logger, host, apiKey)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	input := client.CompletionInput{
		Model:       "neural-chat-7b-v3-3",
		Prompt:      "Will I lose my hair",
		MaxTokens:   1000,
		Temperature: client.Ptr[float32](0.1),
		TopP:        client.Ptr(0.1),
		TopK:        client.Ptr(50),
	}

	resp, err := cln.Completions(ctx, input)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	fmt.Println(resp.Choices[0].Text)

	return nil
}
