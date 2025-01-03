package client_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/predictionguard/go-client"
)

func Test_Client(t *testing.T) {
	service := newService(t)
	defer service.Teardown()

	runTests(t, capabilityTests(service), "capability")
	runTests(t, chatTests(service), "chat")
	runTests(t, completionTests(service), "completion")
	runTests(t, embeddingTests(service), "embedding")
	runTests(t, factualityTests(service), "factuality")
	runTests(t, injectionTests(service), "injection")
	runTests(t, replacePIITests(service), "replacePII")
	runTests(t, rerankTests(service), "rerank")
	runTests(t, tokenizeTests(service), "tokenize")
	runTests(t, toxicityTests(service), "toxicity")
	runTests(t, translateTests(service), "translate")
}

func capabilityTests(srv *service) []table {
	created, _ := time.Parse(time.RFC3339, "2024-10-31T00:00:00Z")

	table := []table{
		{
			Name: "basic",
			ExpResp: client.ModelResponse{
				Object: "list",
				Data: []client.ModelData{
					{
						ID:               "llava-1.5-7b-hf",
						Object:           "model",
						Created:          created,
						OwnedBy:          "llava hugging face",
						Description:      "Open-source multimodal chatbot trained by fine-tuning LLaMa/Vicuna.",
						MaxContextLength: 8192,
						PromptFormat:     "llava",
						Capabilities: client.ModelCapabilities{
							ChatCompletion:     true,
							ChatWithImage:      true,
							Completion:         false,
							Embedding:          false,
							EmbeddingWithImage: false,
							Tokenize:           false,
						},
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				resp, err := srv.Client.Capability(ctx, client.Capabilities.ChatCompletion)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "badkey",
			ExpResp: client.ErrUnauthorized,
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				resp, err := srv.BadClient.Capability(ctx, client.Capabilities.ChatCompletion)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotErr, ok := got.(error)
				if !ok {
					return "didn't get an error"
				}
				expErr := exp.(error)

				if !errors.Is(gotErr, expErr) {
					return "diff"
				}

				return ""
			},
		},
	}

	return table
}

func chatTests(srv *service) []table {
	table := []table{
		{
			Name: "basic-multi",
			ExpResp: client.Chat{
				ID:      "chat-ShL1yk0N0h1lzmrJDQCpCz3WQFQh9",
				Object:  "chat.completion",
				Created: client.ToTime(1715628729),
				Model:   "neural-chat-7b-v3-3",
				Choices: []client.ChatChoice{
					{
						Index: 0,
						Message: client.ChatMessage{
							Role:    client.Roles.Assistant,
							Content: "The world, in general, is full of both beauty and challenges. It can be considered as a mixed bag with various aspects to explore, understand, and appreciate. There are countless achievements in terms of scientific advancements, medical breakthroughs, and technological innovations. On the other hand, the world often encounters issues related to inequality, conflicts, environmental degradation, and moral complexities.\n\nPersonally, it's essential to maintain a balance and perspective while navigating these dimensions. It means trying to find the silver lining behind every storm, practicing gratitude, and embracing empathy to connect with and help others. Actively participating in making the world a better place by supporting causes close to one's heart can also provide a sense of purpose and hope.",
						},
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := client.ChatInputMulti{
					Model: "neural-chat-7b-v3-3",
					Messages: []client.ChatInputMessage{
						{
							Role:    client.Roles.User,
							Content: "How do you feel about the world in general",
						},
					},
					MaxTokens:   client.Ptr(1000),
					Temperature: client.Ptr[float32](0.1),
					TopP:        client.Ptr(0.1),
				}

				resp, err := srv.Client.Chat(ctx, input)
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name: "basic-string",
			ExpResp: client.Chat{
				ID:      "chat-ShL1yk0N0h1lzmrJDQCpCz3WQFQh9",
				Object:  "chat.completion",
				Created: client.ToTime(1715628729),
				Model:   "neural-chat-7b-v3-3",
				Choices: []client.ChatChoice{
					{
						Index: 0,
						Message: client.ChatMessage{
							Role:    client.Roles.Assistant,
							Content: "The world, in general, is full of both beauty and challenges. It can be considered as a mixed bag with various aspects to explore, understand, and appreciate. There are countless achievements in terms of scientific advancements, medical breakthroughs, and technological innovations. On the other hand, the world often encounters issues related to inequality, conflicts, environmental degradation, and moral complexities.\n\nPersonally, it's essential to maintain a balance and perspective while navigating these dimensions. It means trying to find the silver lining behind every storm, practicing gratitude, and embracing empathy to connect with and help others. Actively participating in making the world a better place by supporting causes close to one's heart can also provide a sense of purpose and hope.",
						},
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := client.ChatInput{
					Model:       "neural-chat-7b-v3-3",
					Message:     "How do you feel about the world in general",
					MaxTokens:   client.Ptr(1000),
					Temperature: client.Ptr[float32](0.1),
					TopP:        client.Ptr(0.1),
				}

				resp, err := srv.Client.Chat(ctx, input)
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name: "sse",
			ExpResp: []client.ChatSSE{
				{
					ID:      "chat-OoNijY7ZAkVt4t5Zu8nVDHlW8RAJe",
					Object:  "chat.completion.chunk",
					Created: client.ToTime(1715734993),
					Model:   "neural-chat-7b-v3-3",
					Choices: []client.ChatSSEChoice{
						{
							Index: 0,
							Delta: client.ChatSSEDelta{
								Content: " I",
							},
							Text:         "",
							Probs:        0,
							FinishReason: "",
						},
					},
				},
				{
					ID:      "chat-afH2BnyvKPvon2r16DkUWJygbvePY",
					Object:  "chat.completion.chunk",
					Created: client.ToTime(1715734993),
					Model:   "neural-chat-7b-v3-3",
					Choices: []client.ChatSSEChoice{
						{
							Index: 0,
							Delta: client.ChatSSEDelta{
								Content: " believe",
							},
							Text:         "",
							Probs:        -0.8534317,
							FinishReason: "",
						},
					},
				},
				{
					ID:      "chat-Dd6xpFh5TOtLtFeSxALbmfNNGiyvb",
					Object:  "chat.completion.chunk",
					Created: client.ToTime(1715734995),
					Model:   "neural-chat-7b-v3-3",
					Choices: []client.ChatSSEChoice{
						{
							Index: 0,
							Delta: client.ChatSSEDelta{
								Content: "",
							},
							Text:         "I believe",
							Probs:        0,
							FinishReason: "stop",
						},
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := client.ChatSSEInput{
					Model: "neural-chat-7b-v3-3",
					Messages: []client.ChatInputMessage{
						{
							Role:    client.Roles.User,
							Content: "How do you feel about the world in general",
						},
					},
					MaxTokens:   client.Ptr(1000),
					Temperature: client.Ptr[float32](0.1),
					TopP:        client.Ptr(0.1),
				}

				ch := make(chan client.ChatSSE)

				if err := srv.Client.ChatSSE(ctx, input, ch); err != nil {
					return err
				}

				var sse []client.ChatSSE
				for v := range ch {
					sse = append(sse, v)
				}

				return sse
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name: "vision",
			ExpResp: client.ChatVision{
				ID:      "chat-1qKp6k5y1I4McppJvyHqNkaTeJUtT",
				Object:  "chat.completion",
				Created: client.ToTime(1717441090),
				Model:   "llava-1.5-7b-hf",
				Choices: []client.ChatVisionChoice{
					{
						Index: 0,
						Message: client.ChatVisionMessage{
							Role:    client.Roles.Assistant,
							Content: "No, there is no deer in this picture. The image features a man wearing a hat and glasses, smiling for the camera.",
						},
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := client.ChatVisionInput{
					Model:       "llava-1.5-7b-hf",
					Role:        client.Roles.User,
					Question:    "Is there a deer in this picture?",
					Image:       client.ImageBase64{},
					MaxTokens:   1000,
					Temperature: client.Ptr[float32](0.1),
					TopP:        client.Ptr(0.1),
				}

				resp, err := srv.Client.ChatVision(ctx, input)
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "badkey",
			ExpResp: client.ErrUnauthorized,
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				resp, err := srv.BadClient.Chat(ctx, client.ChatInput{Model: "Hermes-2-Pro-Llama-3-8B"})
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotErr, ok := got.(error)
				if !ok {
					return "didn't get an error"
				}
				expErr := exp.(error)

				if !errors.Is(gotErr, expErr) {
					return "diff"
				}

				return ""
			},
		},
	}

	return table
}

func completionTests(srv *service) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Completion{
				ID:      "cmpl-3gbwD5tLJxklJAljHCjOqMyqUZvv4",
				Object:  "text_completion",
				Created: client.ToTime(1715632193),
				Choices: []client.CompletionChoice{
					{
						Text:  "after weight loss surgery? While losing weight can improve the appearance of your hair and make it appear healthier, some people may experience temporary hair loss in the process.",
						Index: 0,
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := client.CompletionInput{
					Model:       "neural-chat-7b-v3-3",
					Prompt:      "Will I lose my hair",
					MaxTokens:   1000,
					Temperature: client.Ptr[float32](0.1),
					TopP:        client.Ptr(0.1),
				}

				resp, err := srv.Client.Completions(ctx, input)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "badkey",
			ExpResp: client.ErrUnauthorized,
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := client.CompletionInput{
					Model:       "neural-chat-7b-v3-3",
					Prompt:      "Will I lose my hair",
					MaxTokens:   1000,
					Temperature: client.Ptr[float32](0.1),
					TopP:        client.Ptr(0.1),
				}

				resp, err := srv.BadClient.Completions(ctx, input)
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotErr, ok := got.(error)
				if !ok {
					return "didn't get an error"
				}
				expErr := exp.(error)

				if !errors.Is(gotErr, expErr) {
					return "diff"
				}

				return ""
			},
		},
	}

	return table
}

func embeddingTests(srv *service) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Embedding{
				ID:      "emb-0qU4sYEutZvkHskxXwzYDgZVOhtLw",
				Object:  "list",
				Created: client.ToTime(1717439154),
				Model:   "bridgetower-large-itm-mlm-itc",
				Data: []client.EmbeddingData{
					{
						Index:  0,
						Object: "embedding",
						Embedding: []float64{
							0.04457271471619606,
						},
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := client.EmbeddingInputs{
					{
						Text:  "This is Bill Kennedy, a decent Go developer.",
						Image: client.ImageBase64{},
					},
				}

				resp, err := srv.Client.Embedding(ctx, "bridgetower-large-itm-mlm-itc", input)
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name: "ints",
			ExpResp: client.Embedding{
				ID:      "emb-0qU4sYEutZvkHskxXwzYDgZVOhtLw",
				Object:  "list",
				Created: client.ToTime(1717439154),
				Model:   "bridgetower-large-itm-mlm-itc",
				Data: []client.EmbeddingData{
					{
						Index:  0,
						Object: "embedding",
						Embedding: []float64{
							0.04457271471619606,
						},
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := client.EmbeddingIntInputs{
					{0, 3293, 83, 19893, 118963, 25, 7, 3034, 5, 2},
				}

				resp, err := srv.Client.Embedding(ctx, "multilingual-e5-large-instruct", input)
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

func factualityTests(srv *service) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Factuality{
				ID:      "fact-GK9kueuMw0NQLc0sYEIVlkGsPH31R",
				Object:  "factuality.check",
				Created: client.ToTime(1715730425),
				Checks: []client.FactualityCheck{
					{
						Score: 0.7879658937454224,
						Index: 0,
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				reference := "The President shall receive in full for his services during the term for which he shall have been elected compensation in the aggregate amount of 400,000 a year, to be paid monthly, and in addition an expense allowance of 50,000 to assist in defraying expenses relating to or resulting from the discharge of his official duties. Any unused amount of such expense allowance shall revert to the Treasury pursuant to section 1552 of title 31, United States Code. No amount of such expense allowance shall be included in the gross income of the President. He shall be entitled also to the use of the furniture and other effects belonging to the United States and kept in the Executive Residence at the White House."
				text := "The president of the united states can take a salary of one million dollars"

				resp, err := srv.Client.Factuality(ctx, reference, text)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "badkey",
			ExpResp: client.ErrUnauthorized,
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				resp, err := srv.BadClient.Factuality(ctx, "", "")
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotErr, ok := got.(error)
				if !ok {
					return "didn't get an error"
				}
				expErr := exp.(error)

				if !errors.Is(gotErr, expErr) {
					return "diff"
				}

				return ""
			},
		},
	}

	return table
}

func injectionTests(srv *service) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Injection{
				ID:      "injection-Nb817UlEMTog2YOe1JHYbq2oUyZAW7Lk",
				Object:  "injection_check",
				Created: client.ToTime(1715729859),
				Checks: []client.InjectionCheck{
					{
						Probability: 0.5,
						Index:       0,
						Status:      "success",
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				prompt := "A short poem may be a stylistic choice or it may be that you have said what you intended to say in a more concise way."

				resp, err := srv.Client.Injection(ctx, prompt)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "badkey",
			ExpResp: client.ErrUnauthorized,
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				resp, err := srv.BadClient.Injection(ctx, "")
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotErr, ok := got.(error)
				if !ok {
					return "didn't get an error"
				}
				expErr := exp.(error)

				if !errors.Is(gotErr, expErr) {
					return "diff"
				}

				return ""
			},
		},
	}

	return table
}

func replacePIITests(srv *service) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.ReplacePII{
				ID:      "pii-ax9rE9ld3W5yxN1Sz7OKxXkMTMo736jJ",
				Object:  "pii_check",
				Created: client.ToTime(1715730803),
				Checks: []client.ReplacePIICheck{
					{
						NewPrompt: "My email is * and my number is *.",
						Index:     0,
						Status:    "success",
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				prompt := "My email is bill@ardanlabs.com and my number is 954-123-4567."

				resp, err := srv.Client.ReplacePII(ctx, prompt, client.ReplaceMethods.Mask)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "badkey",
			ExpResp: client.ErrUnauthorized,
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				resp, err := srv.BadClient.ReplacePII(ctx, "", client.ReplaceMethods.Mask)
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotErr, ok := got.(error)
				if !ok {
					return "didn't get an error"
				}
				expErr := exp.(error)

				if !errors.Is(gotErr, expErr) {
					return "diff"
				}

				return ""
			},
		},
	}

	return table
}

func rerankTests(srv *service) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Rerank{
				ID:      "rerank-837eef1d-90d1-416a-bf8b-948a42998dd7",
				Object:  "list",
				Created: client.ToTime(1732230548),
				Model:   "bge-reranker-v2-m3",
				Results: []client.RerankResult{
					{
						Index:          0,
						RelevanceScore: 0.06572466,
						Text:           "Deep Learning is not pizza.",
					},
					{
						Index:          1,
						RelevanceScore: 0.054098696,
						Text:           "Deep Learning is pizza.",
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := client.RerankInput{
					Model:           "bge-reranker-v2-m3",
					Query:           "What is Deep Learning?",
					Documents:       []string{"Deep Learning is not pizza.", "Deep Learning is pizza."},
					ReturnDocuments: true,
				}

				resp, err := srv.Client.Rerank(ctx, input)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "badkey",
			ExpResp: client.ErrUnauthorized,
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				resp, err := srv.BadClient.Rerank(ctx, client.RerankInput{})
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotErr, ok := got.(error)
				if !ok {
					return "didn't get an error"
				}
				expErr := exp.(error)

				if !errors.Is(gotErr, expErr) {
					return "diff"
				}

				return ""
			},
		},
	}

	return table
}

func tokenizeTests(srv *service) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Tokenize{
				ID:      "token-ab046fcf-945f-421c-b9f0-1c75ff355203",
				Object:  "tokens",
				Created: client.ToTime(1729871708),
				Data: []client.TokenData{
					{
						ID:    0,
						Start: 0,
						Stop:  0,
						Text:  "<s>",
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				input := client.TokenizeInput{
					Model: "Hermes-2-Pro-Mistral-7B",
					Input: "how many tokens exist for this sentence.",
				}

				resp, err := srv.Client.Tokenize(ctx, input)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "badkey",
			ExpResp: client.ErrUnauthorized,
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				resp, err := srv.BadClient.Tokenize(ctx, client.TokenizeInput{})
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotErr, ok := got.(error)
				if !ok {
					return "didn't get an error"
				}
				expErr := exp.(error)

				if !errors.Is(gotErr, expErr) {
					return "diff"
				}

				return ""
			},
		},
	}

	return table
}

func toxicityTests(srv *service) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Toxicity{
				ID:      "toxi-vRvkxJHmAiSh3NvuuSc48HQ669g7y",
				Object:  "toxicity.check",
				Created: client.ToTime(1715731131),
				Checks: []client.ToxicityCheck{
					{
						Score: 0.7072361707687378,
						Index: 0,
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				text := "Every flight I have is late and I am very angry. I want to hurt someone."

				resp, err := srv.Client.Toxicity(ctx, text)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "badkey",
			ExpResp: client.ErrUnauthorized,
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				resp, err := srv.BadClient.Toxicity(ctx, "")
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotErr, ok := got.(error)
				if !ok {
					return "didn't get an error"
				}
				expErr := exp.(error)

				if !errors.Is(gotErr, expErr) {
					return "diff"
				}

				return ""
			},
		},
	}

	return table
}

func translateTests(srv *service) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Translate{
				ID:                   "translation-0210cae4da704099b58471876ffa3d2e",
				Object:               "translation",
				Created:              client.ToTime(1715731416),
				BestTranslation:      "La lluvia en España permanece principalmente en la llanura",
				BestTranslationModel: "google",
				Score:                0.5381188988685608,
				Translations: []client.Translation{
					{
						Score:       -100,
						Translation: "",
						Model:       "openai",
						Status:      "error: couldn't get translation",
					},
					{
						Score:       0.5008206963539124,
						Translation: "La lluvia en España se queda principalmente en la llanura",
						Model:       "deepl",
						Status:      "success",
					},
					{
						Score:       0.5381188988685608,
						Translation: "La lluvia en España permanece principalmente en la llanura",
						Model:       "google",
						Status:      "success",
					},
					{
						Score:       0.48437628149986267,
						Translation: "La lluvia en España se queda principalmente en la llanura.",
						Model:       "nous_hermes_llama2",
						Status:      "success",
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				text := "The rain in Spain stays mainly in the plain"
				source := client.Languages.English
				target := client.Languages.Spanish

				resp, err := srv.Client.Translate(ctx, text, source, target, false)
				if err != nil {
					return fmt.Errorf("ERROR: %w", err)
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "badkey",
			ExpResp: client.ErrUnauthorized,
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				source := client.Languages.English
				target := client.Languages.Spanish

				resp, err := srv.BadClient.Translate(ctx, "", source, target, false)
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotErr, ok := got.(error)
				if !ok {
					return "didn't get an error"
				}
				expErr := exp.(error)

				if !errors.Is(gotErr, expErr) {
					return "diff"
				}

				return ""
			},
		},
	}

	return table
}

// =============================================================================

type table struct {
	Name    string
	ExpResp any
	ExcFunc func(ctx context.Context) any
	CmpFunc func(got any, exp any) string
}

func runTests(t *testing.T, table []table, testName string) {
	for _, tt := range table {
		f := func(t *testing.T) {
			gotResp := tt.ExcFunc(context.Background())

			diff := tt.CmpFunc(gotResp, tt.ExpResp)
			if diff != "" {
				t.Log("DIFF")
				t.Logf("%s", diff)
				t.Log("GOT")
				t.Logf("%#v", gotResp)
				t.Log("EXP")
				t.Logf("%#v", tt.ExpResp)
				t.Fatal("Should get the expected response")
			}
		}

		t.Run(testName+"-"+tt.Name, f)
	}
}

// =============================================================================

type service struct {
	Client    *client.Client
	BadClient *client.Client
	Teardown  func()
	server    *httptest.Server
}

func newService(t *testing.T) *service {
	var buf bytes.Buffer
	logger := func(ctx context.Context, msg string, v ...any) {
		s := fmt.Sprintf("\nmsg: %s", msg)
		for i := 0; i < len(v); i = i + 2 {
			s = s + fmt.Sprintf(", %s: %v", v[i], v[i+1])
		}
		buf.WriteString(s)
	}

	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)

	cln := client.New(logger, srv.URL, "some-key")
	badCln := client.New(logger, srv.URL, "")

	s := service{
		Client:    cln,
		BadClient: badCln,
		Teardown: func() {
			t.Log("******************** LOGS ********************")
			t.Log(buf.String())
			t.Log("******************** LOGS ********************\n")

			srv.Close()
		},
		server: srv,
	}

	mux.HandleFunc("GET /models/{capability}", s.capability)
	mux.HandleFunc("POST /chat/completions", s.chat)
	mux.HandleFunc("POST /completions", s.completion)
	mux.HandleFunc("POST /factuality", s.factuality)
	mux.HandleFunc("POST /embeddings", s.embeddings)
	mux.HandleFunc("POST /injection", s.injection)
	mux.HandleFunc("POST /PII", s.ReplacePII)
	mux.HandleFunc("POST /rerank", s.Rerank)
	mux.HandleFunc("POST /tokenize", s.tokenize)
	mux.HandleFunc("POST /toxicity", s.toxicity)
	mux.HandleFunc("POST /translate", s.translate)

	return &s
}

func (s *service) capability(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := `{"object":"list","data":[{"id":"llava-1.5-7b-hf","object":"model","created":"2024-10-31T00:00:00Z","owned_by":"llava hugging face","description":"Open-source multimodal chatbot trained by fine-tuning LLaMa/Vicuna.","max_context_length":8192,"prompt_format":"llava","capabilities":{"chat_completion":true,"chat_with_image":true,"completion":false,"embedding":false,"embedding_with_image":false,"tokenize":false}}]}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) chat(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var body struct {
		Model  string `json:"model"`
		Stream bool   `json:"stream"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Decoding Failed", http.StatusInternalServerError)
		return
	}

	if body.Stream {
		s.chatSSE(w)
		return
	}

	var resp string
	switch body.Model {
	case "llava-1.5-7b-hf":
		resp = `{"id":"chat-1qKp6k5y1I4McppJvyHqNkaTeJUtT","object":"chat.completion","created":1717441090,"model":"llava-1.5-7b-hf","choices":[{"index":0,"message":{"role":"assistant","content":"No, there is no deer in this picture. The image features a man wearing a hat and glasses, smiling for the camera.","output":null},"status":"success"}]}`

	default:
		resp = `{"id":"chat-ShL1yk0N0h1lzmrJDQCpCz3WQFQh9","object":"chat.completion","created":1715628729,"model":"neural-chat-7b-v3-3","choices":[{"index":0,"message":{"role":"assistant","content":"The world, in general, is full of both beauty and challenges. It can be considered as a mixed bag with various aspects to explore, understand, and appreciate. There are countless achievements in terms of scientific advancements, medical breakthroughs, and technological innovations. On the other hand, the world often encounters issues related to inequality, conflicts, environmental degradation, and moral complexities.\n\nPersonally, it's essential to maintain a balance and perspective while navigating these dimensions. It means trying to find the silver lining behind every storm, practicing gratitude, and embracing empathy to connect with and help others. Actively participating in making the world a better place by supporting causes close to one's heart can also provide a sense of purpose and hope.","output":null},"status":"success"}]}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) chatSSE(w http.ResponseWriter) {
	events := []string{
		`data: {"id":"chat-OoNijY7ZAkVt4t5Zu8nVDHlW8RAJe","object":"chat.completion.chunk","created":1715734993,"model":"neural-chat-7b-v3-3","choices":[{"index":0,"delta":{"content":" I"},"generated_text":null,"logprobs":0,"finish_reason":null}]}`,
		`data: {"id":"chat-afH2BnyvKPvon2r16DkUWJygbvePY","object":"chat.completion.chunk","created":1715734993,"model":"neural-chat-7b-v3-3","choices":[{"index":0,"delta":{"content":" believe"},"generated_text":null,"logprobs":-0.8534317,"finish_reason":null}]}`,
		`data: {"id":"chat-Dd6xpFh5TOtLtFeSxALbmfNNGiyvb","object":"chat.completion.chunk","created":1715734995,"model":"neural-chat-7b-v3-3","choices":[{"index":0,"delta":{},"generated_text":"I believe","logprobs":0,"finish_reason":"stop"}]}`,
		`data: [DONE]`,
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.WriteHeader(http.StatusOK)

	for _, event := range events {
		if _, err := fmt.Fprintln(w, event); err != nil {
			log.Println(err)
			break
		}
		flusher.Flush()
	}
}

func (s *service) completion(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := `{"id":"cmpl-3gbwD5tLJxklJAljHCjOqMyqUZvv4","object":"text_completion","created":1715632193,"choices":[{"text":"after weight loss surgery? While losing weight can improve the appearance of your hair and make it appear healthier, some people may experience temporary hair loss in the process.","index":0,"status":"success","model":"neural-chat-7b-v3-3"}]}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) embeddings(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := `{"id":"emb-0qU4sYEutZvkHskxXwzYDgZVOhtLw","object":"list","created":1717439154,"model":"bridgetower-large-itm-mlm-itc","data":[{"status":"success","index":0,"object":"embedding","embedding":[0.04457271471619606]}]}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) factuality(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := `{"checks":[{"score":0.7879658937454224,"index":0,"status":"success"}],"created":1715730425,"id":"fact-GK9kueuMw0NQLc0sYEIVlkGsPH31R","object":"factuality.check"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) injection(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := `{"checks":[{"probability":0.5,"index":0,"status":"success"}],"created":"1715729859","id":"injection-Nb817UlEMTog2YOe1JHYbq2oUyZAW7Lk","object":"injection_check"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) ReplacePII(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := `{"checks":[{"new_prompt":"My email is * and my number is *.","index":0,"status":"success"}],"created":"1715730803","id":"pii-ax9rE9ld3W5yxN1Sz7OKxXkMTMo736jJ","object":"pii_check"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) Rerank(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := `{"id":"rerank-837eef1d-90d1-416a-bf8b-948a42998dd7","object":"list","created":1732230548,"model":"bge-reranker-v2-m3","results":[{"index":0,"relevance_score":0.06572466,"text":"Deep Learning is not pizza."},{"index":1,"relevance_score":0.054098696,"text":"Deep Learning is pizza."}]}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) tokenize(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := `{"id": "token-ab046fcf-945f-421c-b9f0-1c75ff355203","object": "tokens","created": 1729871708,"model": "multilingual-e5-large-instruct","data": [{"id": 0,"start": 0,"stop": 0,"text": "<s>"}]}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) toxicity(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := `{"checks":[{"score":0.7072361707687378,"index":0,"status":"success"}],"created":1715731131,"id":"toxi-vRvkxJHmAiSh3NvuuSc48HQ669g7y","object":"toxicity.check"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) translate(w http.ResponseWriter, r *http.Request) {
	if v := r.Header.Get("authorization"); v == "Bearer" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := `{"translations":[{"score":-100,"translation":"","model":"openai","status":"error: couldn't get translation"},{"score":0.5008206963539124,"translation":"La lluvia en España se queda principalmente en la llanura","model":"deepl","status":"success"},{"score":0.5381188988685608,"translation":"La lluvia en España permanece principalmente en la llanura","model":"google","status":"success"},{"score":0.48437628149986267,"translation":"La lluvia en España se queda principalmente en la llanura.","model":"nous_hermes_llama2","status":"success"}],"best_translation":"La lluvia en España permanece principalmente en la llanura","best_score":0.5381188988685608,"best_translation_model":"google","created":1715731416,"id":"translation-0210cae4da704099b58471876ffa3d2e","object":"translation"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

// =============================================================================

func ExampleClient_Capability() {
	// examples/capability/main.go

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

	resp, err := cln.Capability(ctx, client.Capabilities.ChatCompletion)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println(resp)
}

func ExampleClient_Chat() {
	// examples/chat/basic/main.go

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

	input := client.ChatInput{
		Model:       "neural-chat-7b-v3-3",
		Message:     "How do you feel about the world in general",
		MaxTokens:   client.Ptr(1000),
		Temperature: client.Ptr[float32](0.1),
		TopP:        client.Ptr(0.1),
		InputExtension: &client.InputExtension{
			PII:              client.PIIs.Replace,
			PIIReplaceMethod: client.ReplaceMethods.Random,
		},
		OutputExtension: &client.OutputExtension{
			Factuality: true,
			Toxicity:   true,
		},
	}

	resp, err := cln.Chat(ctx, input)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)

	inputMulti := client.ChatInputMulti{
		Model: "neural-chat-7b-v3-3",
		Messages: []client.ChatInputMessage{
			{
				Role:    client.Roles.User,
				Content: "How do you feel about the world in general",
			},
		},
		MaxTokens:   client.Ptr(1000),
		Temperature: client.Ptr[float32](0.1),
		TopP:        client.Ptr(0.1),
		InputExtension: &client.InputExtension{
			PII:              client.PIIs.Replace,
			PIIReplaceMethod: client.ReplaceMethods.Random,
		},
		OutputExtension: &client.OutputExtension{
			Factuality: true,
			Toxicity:   true,
		},
	}

	resp, err = cln.Chat(ctx, inputMulti)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func ExampleClient_ChatSSE() {
	// examples/chat/sse/main.go

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

	input := client.ChatSSEInput{
		Model: "neural-chat-7b-v3-3",
		Messages: []client.ChatInputMessage{
			{
				Role:    client.Roles.User,
				Content: "How do you feel about the world in general",
			},
		},
		MaxTokens:   client.Ptr(1000),
		Temperature: client.Ptr[float32](0.1),
		TopP:        client.Ptr(0.1),
	}

	ch := make(chan client.ChatSSE, 100)

	err := cln.ChatSSE(ctx, input, ch)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	for resp := range ch {
		for _, choice := range resp.Choices {
			fmt.Print(choice.Delta.Content)
		}
	}
}

func ExampleClient_ChatVision() {
	// examples/chat/vision/main.go

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

	image, err := client.NewImageNetwork("https://pbs.twimg.com/profile_images/1571574401107169282/ylAgz_f5_400x400.jpg")
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	input := client.ChatVisionInput{
		Model:       "llava-1.5-7b-hf",
		Role:        client.Roles.User,
		Question:    "Is there a deer in this picture?",
		Image:       image,
		MaxTokens:   1000,
		Temperature: client.Ptr[float32](0.1),
		TopP:        client.Ptr(0.1),
	}

	resp, err := cln.ChatVision(ctx, input)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	for i, choice := range resp.Choices {
		fmt.Printf("choice %d: %s\n", i, choice.Message.Content)
	}
}

func ExampleClient_Completions() {
	// examples/completion/main.go

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

	input := client.CompletionInput{
		Model:       "neural-chat-7b-v3-3",
		Prompt:      "Will I lose my hair",
		MaxTokens:   1000,
		Temperature: client.Ptr[float32](0.1),
		TopP:        client.Ptr(0.1),
	}

	resp, err := cln.Completions(ctx, input)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println(resp.Choices[0].Text)
}

func ExampleClient_Embedding() {
	// examples/embedding/basic/main.go

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

	image, err := client.NewImageNetwork("https://pbs.twimg.com/profile_images/1571574401107169282/ylAgz_f5_400x400.jpg")
	if err != nil {
		log.Fatalln("ERROR: %w", err)
	}

	input := client.EmbeddingInputs{
		{
			Text:  "This is Bill Kennedy, a decent Go developer.",
			Image: image,
		},
	}

	resp, err := cln.Embedding(ctx, "bridgetower-large-itm-mlm-itc", input)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	for _, data := range resp.Data {
		fmt.Print(data.Embedding)
	}
}

func ExampleClient_EmbeddingWithTruncate() {
	// examples/embedding/truncate/main.go

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
		log.Fatalln("ERROR:", err)
	}

	for _, data := range resp.Data {
		fmt.Print(data.Embedding)
	}
}

func ExampleClient_Factuality() {
	// examples/factuality/main.go

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

	fact := "The President shall receive in full for his services during the term for which he shall have been elected compensation in the aggregate amount of 400,000 a year, to be paid monthly, and in addition an expense allowance of 50,000 to assist in defraying expenses relating to or resulting from the discharge of his official duties. Any unused amount of such expense allowance shall revert to the Treasury pursuant to section 1552 of title 31, United States Code. No amount of such expense allowance shall be included in the gross income of the President. He shall be entitled also to the use of the furniture and other effects belonging to the United States and kept in the Executive Residence at the White House."
	text := "The president of the united states can take a salary of one million dollars"

	resp, err := cln.Factuality(ctx, fact, text)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	log.Println(resp.Checks[0])
}

func ExampleClient_Injection() {
	// examples/injection/main.go

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

	prompt := "A short poem may be a stylistic choice or it may be that you have said what you intended to say in a more concise way."

	resp, err := cln.Injection(ctx, prompt)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println(resp.Checks[0].Probability)
}

func ExampleClient_ReplacePII() {
	// examples/ReplacePII/main.go

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

	prompt := "My email is bill@ardanlabs.com and my number is 954-123-4567."

	resp, err := cln.ReplacePII(ctx, prompt, client.ReplaceMethods.Mask)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println(resp.Checks[0].NewPrompt)
}

func ExampleClient_Rerank() {
	// examples/rerank/main.go

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

	input := client.RerankInput{
		Model:           "bge-reranker-v2-m3",
		Query:           "What is Deep Learning?",
		Documents:       []string{"Deep Learning is not pizza.", "Deep Learning is pizza."},
		ReturnDocuments: true,
	}

	resp, err := cln.Rerank(ctx, input)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println(resp.Results)
}

func ExampleClient_Tokenize() {
	// examples/tokenize/main.go

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

	input := client.TokenizeInput{
		Model: "Hermes-2-Pro-Mistral-7B",
		Input: "how many tokens exist for this sentence.",
	}

	resp, err := cln.Tokenize(ctx, input)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println(resp.Data)
}

func ExampleClient_Toxicity() {
	// examples/toxicity/main.go

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

	text := "Every flight I have is late and I am very angry. I want to hurt someone."

	resp, err := cln.Toxicity(ctx, text)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println(resp.Checks[0].Score)
}

func ExampleClient_Translate() {
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

	text := "The rain in Spain stays mainly in the plain"

	resp, err := cln.Translate(ctx, text, client.Languages.English, client.Languages.Spanish, false)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println(resp.BestTranslation)
}
