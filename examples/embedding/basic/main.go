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

	image, err := client.NewImageNetwork("https://predictionguard.com/lib_eltrNYEjQbpUWFRI/oy2r533pndpk0q8q.png?w=1024&dpr=2")
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	input := []client.EmbeddingInput{
		{
			Text:  "This is prediction guard.",
			Image: image,
		},
	}

	resp, err := cln.Embedding(ctx, "bridgetower-large-itm-mlm-itc", input)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	for _, data := range resp.Data {
		fmt.Print(data.Embedding)
	}

	return nil
}
