package client_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/predictionguard/go-client"
)

func Test_Chat(t *testing.T) {
	service := newService(t)
	defer service.Teardown()

	runTests(t, chatOKTests(service.Client), "chat-ok")
}

func chatOKTests(cln *client.Client) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Chat{
				ID:      "chat-ShL1yk0N0h1lzmrJDQCpCz3WQFQh9",
				Object:  "chat_completion",
				Created: client.ToTime(1715628729),
				Model:   client.Models.NeuralChat7B,
				Choices: []struct {
					Index   int                `json:"index"`
					Message client.ChatMessage `json:"message"`
					Status  string             `json:"status"`
				}{
					{
						Index: 0,
						Message: client.ChatMessage{
							Role:    client.Roles.Assistant,
							Content: "The world, in general, is full of both beauty and challenges. It can be considered as a mixed bag with various aspects to explore, understand, and appreciate. There are countless achievements in terms of scientific advancements, medical breakthroughs, and technological innovations. On the other hand, the world often encounters issues related to inequality, conflicts, environmental degradation, and moral complexities.\n\nPersonally, it's essential to maintain a balance and perspective while navigating these dimensions. It means trying to find the silver lining behind every storm, practicing gratitude, and embracing empathy to connect with and help others. Actively participating in making the world a better place by supporting causes close to one's heart can also provide a sense of purpose and hope.",
							Output:  "",
						},
						Status: "success",
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := []client.ChatMessage{
					{
						Role:    client.Roles.User,
						Content: "How do you feel about the world in general",
					},
				}

				resp, err := cln.Chat(ctx, client.Models.NeuralChat7B, input, 1000, 1.1)
				if err != nil {
					return err
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
