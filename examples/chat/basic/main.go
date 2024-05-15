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
	apiKey := os.Getenv("PGKEY")

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

	input := []client.ChatMessage{
		{
			Role:    client.Roles.User,
			Content: "How do you feel about the world in general",
		},
	}

	resp, err := cln.Chat(ctx, client.Models.NeuralChat7B, input, 1000, 1.1)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return nil
}
