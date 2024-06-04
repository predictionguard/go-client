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
	host := "https://staging.predictionguard.com"
	apiKey := os.Getenv("PGKEYSTAGE")

	logger := func(ctx context.Context, msg string, v ...any) {
		s := fmt.Sprintf("msg: %s", msg)
		for i := 0; i < len(v); i = i + 2 {
			s = s + fmt.Sprintf(", %s: %v", v[i], v[i+1])
		}
		log.Println(s)
	}

	cln := client.New(logger, host, apiKey)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	image, err := client.NewImageNetwork("https://pbs.twimg.com/profile_images/1571574401107169282/ylAgz_f5_400x400.jpg")
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	// image, err := client.NewImageFile("/Users/bill/Documents/images/pGwOq5tz_400x400.jpg")
	// if err != nil {
	// 	return fmt.Errorf("ERROR: %w", err)
	// }

	input := client.ChatVisionInput{
		Role:        client.Roles.User,
		Question:    "Is there a deer in this picture?",
		Image:       image,
		MaxTokens:   1000,
		Temperature: 0.1,
		TopP:        0.1,
	}

	resp, err := cln.ChatVision(ctx, input)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	for i, choice := range resp.Choices {
		fmt.Printf("choice %d: %s\n", i, choice.Message.Content)
	}

	return nil
}
