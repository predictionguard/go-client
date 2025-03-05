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
		"model":    "bridgetower-large-itm-mlm-itc",
		"truncate": false,
		"input": [][]int{
			{0, 3293, 83, 19893, 118963, 25, 7, 3034, 5, 2},
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
