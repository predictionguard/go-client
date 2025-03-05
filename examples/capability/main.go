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

	url := "https://api.predictionguard.com/models/" + client.Capabilities.ChatCompletion.String()

	var resp client.ModelResponse
	if err := cln.Do(ctx, http.MethodGet, url, nil, &resp); err != nil {
		return fmt.Errorf("do: %w", err)
	}

	fmt.Println(resp)

	return nil
}
