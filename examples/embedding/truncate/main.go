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

	input := client.EmbeddingInputs{
		{
			Text: "This is Bill Kennedy, a decent Go developer.",
		},
	}

	resp, err := cln.EmbeddingWithTruncate(ctx, "multilingual-e5-large-instruct", input, client.Directions.Right)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	for _, data := range resp.Data {
		fmt.Print(data.Embedding)
	}

	return nil
}
