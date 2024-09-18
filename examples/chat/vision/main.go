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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	image, err := client.NewImageNetwork("https://predictionguard.com/lib_eltrNYEjQbpUWFRI/oy2r533pndpk0q8q.png?w=1024&dpr=2")
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	input := client.ChatVisionInput{
		Model:       "llava-1.5-7b-hf",
		Role:        client.Roles.User,
		Question:    "Is there a computer in this picture?",
		Image:       image,
		MaxTokens:   1000,
		Temperature: client.Ptr[float32](0.1),
		TopP:        client.Ptr(0.1),
		TopK:        client.Ptr(50),
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
