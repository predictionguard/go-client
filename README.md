# Prediction Guard Go Client

[![CircleCI](https://dl.circleci.com/status-badge/img/circleci/Cy6tWW4wpE69Ftb8vdTAN9/E2TBj5h2YvKmwX36hcykvy/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/circleci/Cy6tWW4wpE69Ftb8vdTAN9/E2TBj5h2YvKmwX36hcykvy/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/predictionguard/go-client)](https://goreportcard.com/report/github.com/predictionguard/go-client)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/predictionguard/go-client)](https://pkg.go.dev/github.com/predictionguard/go-client)

Copyright 2024 Prediction Guard
bill@predictionguard.com

### Description

This Module provides functionality developed to simplify interfacing with [Prediction Guard API](https://www.predictionguard.com/) in Go.

### Requirements

To access the API, contact us [here](https://www.predictionguard.com/getting-started) to get an enterprise access token. You will need this access token to continue.

### Usage

```go
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

	input := client.ChatInput{
		Model: client.Models.NeuralChat7B,
		Messages: []client.ChatInputMessage{
			{
				Role:    client.Roles.User,
				Content: "How do you feel about the world in general",
			},
		},
		MaxTokens:   1000,
		Temperature: 0.1,
		TopP:        0.1,
		Options: &client.ChatInputOptions{
			Factuality:       true,
			Toxicity:         true,
			PII:              client.PIIs.Replace,
			PIIReplaceMethod: client.ReplaceMethods.Random,
		},
	}

	resp, err := cln.Chat(ctx, input)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return nil
}
```
Take a look at the `examples` directory for more examples.

### Docs

You can find the SDK and Prediction Guard docs using these links.

[SDK Docs](https://pkg.go.dev/github.com/predictionguard/go-client)

[PG API Docs](https://docs.predictionguard.com/docs/getting-started/welcome)

### Getting started

Once you have your api key you can use the `makefile` to run curl commands for the different api endpoints. For example, `make curl-injection` will connect to the injection endpoint and return the injection response. The `makefile` also allows you to run the different examples such as `make go-injection` to run the Go injection example.

#### Licensing

```
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
Copyright 2024 Prediction Guard
