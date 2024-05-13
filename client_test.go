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

	return &s
}

func (s *service) chat(w http.ResponseWriter, r *http.Request) {
	resp := `{"id":"chat-ShL1yk0N0h1lzmrJDQCpCz3WQFQh9","object":"chat_completion","created":1715628729,"model":"Neural-Chat-7B","choices":[{"index":0,"message":{"role":"assistant","content":"The world, in general, is full of both beauty and challenges. It can be considered as a mixed bag with various aspects to explore, understand, and appreciate. There are countless achievements in terms of scientific advancements, medical breakthroughs, and technological innovations. On the other hand, the world often encounters issues related to inequality, conflicts, environmental degradation, and moral complexities.\n\nPersonally, it's essential to maintain a balance and perspective while navigating these dimensions. It means trying to find the silver lining behind every storm, practicing gratitude, and embracing empathy to connect with and help others. Actively participating in making the world a better place by supporting causes close to one's heart can also provide a sense of purpose and hope.","output":null},"status":"success"}]}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
