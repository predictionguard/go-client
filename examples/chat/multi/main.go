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

	inputMulti := client.ChatInputMulti{
		Model: "neural-chat-7b-v3-3",
		Messages: []client.ChatInputMessage{
			{
				Role:    client.Roles.User,
				Content: "How do you feel about the world in general",
			},
		},
		MaxTokens:   client.Ptr(1000),
		Temperature: client.Ptr[float32](0.1),
		TopP:        client.Ptr(0.1),
		TopK:        client.Ptr(50),
		InputExtension: &client.InputExtension{
			PII:              client.PIIs.Replace,
			PIIReplaceMethod: client.ReplaceMethods.Random,
		},
		OutputExtension: &client.OutputExtension{
			Factuality: true,
			Toxicity:   true,
		},
	}

	resp, err := cln.Chat(ctx, inputMulti)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return nil
}
