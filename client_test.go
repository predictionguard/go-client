package client_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/predictionguard/go-client"
)

func Test_Client(t *testing.T) {
	service := newService(t)
	defer service.Teardown()

	runTests(t, chatTests(service.Client), "chat")
	runTests(t, completionTests(service.Client), "completion")
	runTests(t, factualityTests(service.Client), "factuality")
	runTests(t, injectionTests(service.Client), "injection")
	runTests(t, replacepiTests(service.Client), "replacepi")
	runTests(t, toxicityTests(service.Client), "toxicity")
	runTests(t, translateTests(service.Client), "translate")
}

func chatTests(cln *client.Client) []table {
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

func completionTests(cln *client.Client) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Completion{
				ID:      "cmpl-3gbwD5tLJxklJAljHCjOqMyqUZvv4",
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

func factualityTests(cln *client.Client) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Factuality{
				ID:      "fact-GK9kueuMw0NQLc0sYEIVlkGsPH31R",
				Object:  "factuality_check",
				Created: client.ToTime(1715730425),
				Checks: []struct {
					Score  float64 `json:"score"`
					Index  int     `json:"index"`
					Status string  `json:"status"`
				}{
					{
						Score:  0.7879658937454224,
						Index:  0,
						Status: "success",
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				reference := "The President shall receive in full for his services during the term for which he shall have been elected compensation in the aggregate amount of 400,000 a year, to be paid monthly, and in addition an expense allowance of 50,000 to assist in defraying expenses relating to or resulting from the discharge of his official duties. Any unused amount of such expense allowance shall revert to the Treasury pursuant to section 1552 of title 31, United States Code. No amount of such expense allowance shall be included in the gross income of the President. He shall be entitled also to the use of the furniture and other effects belonging to the United States and kept in the Executive Residence at the White House."
				text := "The president of the united states can take a salary of one million dollars"

				resp, err := cln.Factuality(ctx, reference, text)
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

func injectionTests(cln *client.Client) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Injection{
				ID:      "injection-Nb817UlEMTog2YOe1JHYbq2oUyZAW7Lk",
				Object:  "injection_check",
				Created: client.ToTime(1715729859),
				Checks: []struct {
					Probability float64 `json:"probability"`
					Index       int     `json:"index"`
					Status      string  `json:"status"`
				}{
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

				resp, err := cln.Injection(ctx, prompt)
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

func replacepiTests(cln *client.Client) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.ReplacePI{
				ID:      "pii-ax9rE9ld3W5yxN1Sz7OKxXkMTMo736jJ",
				Object:  "pii_check",
				Created: client.ToTime(1715730803),
				Checks: []struct {
					Text   string `json:"new_prompt"`
					Index  int    `json:"index"`
					Status string `json:"status"`
				}{
					{
						Text:   "My email is * and my number is *.",
						Index:  0,
						Status: "success",
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				prompt := "My email is bill@ardanlabs.com and my number is 954-123-4567."

				resp, err := cln.ReplacePI(ctx, prompt, client.ReplaceMethods.Mask)
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

func toxicityTests(cln *client.Client) []table {
	table := []table{
		{
			Name: "basic",
			ExpResp: client.Toxicity{
				ID:      "toxi-vRvkxJHmAiSh3NvuuSc48HQ669g7y",
				Object:  "toxicity_check",
				Created: client.ToTime(1715731131),
				Checks: []struct {
					Score  float64 `json:"score"`
					Index  int     `json:"index"`
					Status string  `json:"status"`
				}{
					{
						Score:  0.7072361707687378,
						Index:  0,
						Status: "success",
					},
				},
			},
			ExcFunc: func(ctx context.Context) any {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				text := "Every flight I have is late and I am very angry. I want to hurt someone."

				resp, err := cln.Toxicity(ctx, text)
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

func translateTests(cln *client.Client) []table {
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
				Translations: []struct {
					Score       float64 `json:"score"`
					Translation string  `json:"translation"`
					Model       string  `json:"model"`
					Status      string  `json:"status"`
				}{
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

				resp, err := cln.Translate(ctx, text, source, target)
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

// =============================================================================

type table struct {
	Name    string
	ExpResp any
	ExcFunc func(ctx context.Context) any
	CmpFunc func(got any, exp any) string
}

func runTests(t *testing.T, table []table, testName string) {
	log := func(diff string, got any, exp any) {
		t.Log("DIFF")
		t.Logf("%s", diff)
		t.Log("GOT")
		t.Logf("%#v", got)
		t.Log("EXP")
		t.Logf("%#v", exp)
		t.Fatal("Should get the expected response")
	}

	for _, tt := range table {
		f := func(t *testing.T) {
			gotResp := tt.ExcFunc(context.Background())

			diff := tt.CmpFunc(gotResp, tt.ExpResp)
			if diff != "" {
				log(diff, gotResp, tt.ExpResp)
			}
		}

		t.Run(testName+"-"+tt.Name, f)
	}
}

// =============================================================================

type service struct {
	Client   *client.Client
	Teardown func()
	server   *httptest.Server
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

	client := client.New(logger, srv.URL, "")

	s := service{
		Client: client,
		Teardown: func() {
			t.Log("******************** LOGS ********************")
			t.Log(buf.String())
			t.Log("******************** LOGS ********************\n")

			srv.Close()
		},
		server: srv,
	}

	mux.HandleFunc("/chat/completions", s.chat)
	mux.HandleFunc("/completions", s.completion)
	mux.HandleFunc("/injection", s.injection)
	mux.HandleFunc("/factuality", s.factuality)
	mux.HandleFunc("/PII", s.replacePI)
	mux.HandleFunc("/toxicity", s.toxicity)
	mux.HandleFunc("/translate", s.translate)

	return &s
}

func (s *service) chat(w http.ResponseWriter, r *http.Request) {
	resp := `{"id":"chat-ShL1yk0N0h1lzmrJDQCpCz3WQFQh9","object":"chat_completion","created":1715628729,"model":"Neural-Chat-7B","choices":[{"index":0,"message":{"role":"assistant","content":"The world, in general, is full of both beauty and challenges. It can be considered as a mixed bag with various aspects to explore, understand, and appreciate. There are countless achievements in terms of scientific advancements, medical breakthroughs, and technological innovations. On the other hand, the world often encounters issues related to inequality, conflicts, environmental degradation, and moral complexities.\n\nPersonally, it's essential to maintain a balance and perspective while navigating these dimensions. It means trying to find the silver lining behind every storm, practicing gratitude, and embracing empathy to connect with and help others. Actively participating in making the world a better place by supporting causes close to one's heart can also provide a sense of purpose and hope.","output":null},"status":"success"}]}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) completion(w http.ResponseWriter, r *http.Request) {
	resp := `{"id":"cmpl-3gbwD5tLJxklJAljHCjOqMyqUZvv4","object":"text_completion","created":1715632193,"choices":[{"text":"after weight loss surgery? While losing weight can improve the appearance of your hair and make it appear healthier, some people may experience temporary hair loss in the process.","index":0,"status":"success","model":"Neural-Chat-7B"}]}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) injection(w http.ResponseWriter, r *http.Request) {
	resp := `{"checks":[{"probability":0.5,"index":0,"status":"success"}],"created":"1715729859","id":"injection-Nb817UlEMTog2YOe1JHYbq2oUyZAW7Lk","object":"injection_check"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) factuality(w http.ResponseWriter, r *http.Request) {
	resp := `{"checks":[{"score":0.7879658937454224,"index":0,"status":"success"}],"created":1715730425,"id":"fact-GK9kueuMw0NQLc0sYEIVlkGsPH31R","object":"factuality_check"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) replacePI(w http.ResponseWriter, r *http.Request) {
	resp := `{"checks":[{"new_prompt":"My email is * and my number is *.","index":0,"status":"success"}],"created":"1715730803","id":"pii-ax9rE9ld3W5yxN1Sz7OKxXkMTMo736jJ","object":"pii_check"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) toxicity(w http.ResponseWriter, r *http.Request) {
	resp := `{"checks":[{"score":0.7072361707687378,"index":0,"status":"success"}],"created":1715731131,"id":"toxi-vRvkxJHmAiSh3NvuuSc48HQ669g7y","object":"toxicity_check"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s *service) translate(w http.ResponseWriter, r *http.Request) {
	resp := `{"translations":[{"score":-100,"translation":"","model":"openai","status":"error: couldn't get translation"},{"score":0.5008206963539124,"translation":"La lluvia en España se queda principalmente en la llanura","model":"deepl","status":"success"},{"score":0.5381188988685608,"translation":"La lluvia en España permanece principalmente en la llanura","model":"google","status":"success"},{"score":0.48437628149986267,"translation":"La lluvia en España se queda principalmente en la llanura.","model":"nous_hermes_llama2","status":"success"}],"best_translation":"La lluvia en España permanece principalmente en la llanura","best_score":0.5381188988685608,"best_translation_model":"google","created":1715731416,"id":"translation-0210cae4da704099b58471876ffa3d2e","object":"translation"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
