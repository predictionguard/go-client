package client_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/predictionguard/go-client"
)

func Test_Completion(t *testing.T) {
	service := newService(t)
	defer service.Teardown()

	runTests(t, chatOKTests(service.Client), "completion-ok")
}

func completionOKTests(cln *client.Client) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Completion{
				ID:      "chat-3gbwD5tLJxklJAljHCjOqMyqUZvv4",
				Object:  "text_completion",
				Created: client.ToTime(1715632193),
				Choices: []struct {
					Text   string `json:"text"`
					Index  int    `json:"index"`
					Status string `json:"status"`
					Model  string `json:"model"`
				}{
					{
						Text:   "after weight loss surgery? While losing weight can improve the appearance of your hair and make it appear healthier, some people may experience temporary hair loss in the process.",
						Index:  0,
						Status: "success",
						Model:  "Neural-Chat-7B",
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				resp, err := cln.Completions(ctx, client.Models.NeuralChat7B, "Will I lose my hair", 1000, 1.1)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
	}

	return table
}
