# Prediction Guard Go Client

[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/predictionguard/go-client)](https://pkg.go.dev/github.com/predictionguard/go-client)

> [!WARNING]
> **This module is deprecated and no longer maintained.** Some features are broken or missing, and no further updates will be released. Existing versions remain available, but you should migrate to an OpenAI-compatible or Anthropic-compatible client as described below.

Copyright 2024 Prediction Guard
bill@predictionguard.com

## Migrating

The Prediction Guard API is compatible with both OpenAI-style and Anthropic-style clients. Use whichever SDK matches the functionality you need, pointed at the Prediction Guard API with your existing API key.

### OpenAI-compatible (`openai-go`)

```
go get github.com/openai/openai-go/v3
```

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func main() {
	cln := openai.NewClient(
		option.WithAPIKey(os.Getenv("PREDICTIONGUARD_API_KEY")),
		option.WithBaseURL("https://api.predictionguard.com"),
	)

	resp, err := cln.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Model: "<model-name>",
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("How do you feel about the world in general?"),
		},
		MaxTokens: openai.Int(1000),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
```

### Anthropic-compatible (`anthropic-sdk-go`)

```
go get github.com/anthropics/anthropic-sdk-go
```

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func main() {
	cln := anthropic.NewClient(
		option.WithAPIKey(os.Getenv("PREDICTIONGUARD_API_KEY")),
		option.WithBaseURL("https://api.predictionguard.com"),
	)

	msg, err := cln.Messages.New(context.Background(), anthropic.MessageNewParams{
		Model:     "<model-name>",
		MaxTokens: 1000,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock("How do you feel about the world in general?")),
		},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(msg.Content[0].Text)
}
```

## Docs

For the full list of endpoints and models, see the [Prediction Guard documentation](https://docs.predictionguard.com).

The [API documentation for this module](https://pkg.go.dev/github.com/predictionguard/go-client/v2) remains available for reference but will not be updated.
