package client_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/predictionguard/go-client"
)

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

/*
FACT
{"checks":[{"score":0.7879658937454224,"index":0,"status":"success"}],"created":1715633327,"id":"fact-qpmMj2oorfdELPkD4z5J8KLjFqIki","object":"factuality_check"}%

FACT - NO KEY
HTTP/2 403
date: Mon, 13 May 2024 20:50:59 GMT
content-length: 0

TRANS
{"translations":[{"score":-100,"translation":"","model":"openai","status":"error: couldn’t get translation"},{"score":0.5008206963539124,"translation":"La lluvia en España se queda principalmente en la llanura","model":"deepl","status":"success"},{"score":0.5381188988685608,"translation":"La lluvia en España permanece principalmente en la llanura","model":"google","status":"success"},{"score":0.48437628149986267,"translation":"La lluvia en España se queda principalmente en la llanura.","model":"nous_hermes_llama2","status":"success"}],"best_translation":"La lluvia en España permanece principalmente en la llanura","best_score":0.5381188988685608,"best_translation_model":"google","created":1715633371,"id":"translation-090df50bb3424396adf1b8f19228ad3a","object":"translation"}%

RPI
{"checks":[{"new_prompt":"My email is * and my number is *.","index":0,"status":"success"}],"created":"1715633389","id":"pii-jzp3PZFSn9DWFe5D2aWd0Lgk45qG41U0","object":"pii_check"}%

DI
{"checks":[{"probability":0.5,"index":0,"status":"success"}],"created":"1715633407","id":"injection-1LZa9ftEJFALPpKRXIPQz5Hx2QMspL7B","object":"injection_check"}%

TOX
{"checks":[{"score":0.7072361707687378,"index":0,"status":"success"}],"created":1715633419,"id":"toxi-v0uq3q9cFid7PMU4spM7gT3XEmeg2","object":"toxicity_check"}%
*/
