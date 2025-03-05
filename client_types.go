package client

import (
	"strconv"
	"strings"
	"time"
)

type D map[string]any

// =============================================================================

type Error struct {
	Message string `json:"error"`
}

func (err *Error) Error() string {
	return err.Message
}

// =============================================================================

type Time struct {
	time.Time
}

func ToTime(sec int64) Time {
	return Time{
		Time: time.Unix(sec, 0),
	}
}

func (t *Time) UnmarshalJSON(data []byte) error {
	d := strings.Trim(string(data), "\"")

	num, err := strconv.Atoi(d)
	if err != nil {
		return err
	}

	t.Time = time.Unix(int64(num), 0)

	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	data := strconv.Itoa(int(t.Unix()))
	return []byte(data), nil
}

// =============================================================================

type ModelCapabilities struct {
	ChatCompletion     bool `json:"chat_completion"`
	ChatWithImage      bool `json:"chat_with_image"`
	Completion         bool `json:"completion"`
	Embedding          bool `json:"embedding"`
	EmbeddingWithImage bool `json:"embedding_with_image"`
	Tokenize           bool `json:"tokenize"`
}

type ModelData struct {
	ID               string            `json:"id"`
	Object           string            `json:"object"`
	Created          time.Time         `json:"created"`
	OwnedBy          string            `json:"owned_by"`
	Description      string            `json:"description"`
	MaxContextLength int               `json:"max_context_length"`
	PromptFormat     string            `json:"prompt_format"`
	Capabilities     ModelCapabilities `json:"capabilities"`
}

type ModelResponse struct {
	Object string      `json:"object"`
	Data   []ModelData `json:"data"`
}

// =============================================================================

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatChoice struct {
	Index   int         `json:"index"`
	Message ChatMessage `json:"message"`
}

type Chat struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created Time         `json:"created"`
	Model   string       `json:"model"`
	Choices []ChatChoice `json:"choices"`
}

// =============================================================================

type ChatSSEDelta struct {
	Content string `json:"content"`
}

type ChatSSEChoice struct {
	Index        int          `json:"index"`
	Delta        ChatSSEDelta `json:"delta"`
	Text         string       `json:"generated_text"`
	Probs        float32      `json:"logprobs"`
	FinishReason string       `json:"finish_reason"`
}

type ChatSSE struct {
	ID      string          `json:"id"`
	Object  string          `json:"object"`
	Created Time            `json:"created"`
	Model   string          `json:"model"`
	Choices []ChatSSEChoice `json:"choices"`
	Error   string          `json:"error"`
}

// =============================================================================

type ChatVisionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatVisionChoice represents a choice for the vision call.
type ChatVisionChoice struct {
	Index   int               `json:"index"`
	Message ChatVisionMessage `json:"message"`
}

// ChatVision represents the result for the vision call.
type ChatVision struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created Time               `json:"created"`
	Model   string             `json:"model"`
	Choices []ChatVisionChoice `json:"choices"`
}

// =============================================================================

type CompletionChoice struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

type Completion struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created Time               `json:"created"`
	Model   string             `json:"model"`
	Choices []CompletionChoice `json:"choices"`
}

// =============================================================================

type EmbeddingData struct {
	Index     int       `json:"index"`
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
}

type Embedding struct {
	ID      string          `json:"id"`
	Object  string          `json:"object"`
	Created Time            `json:"created"`
	Model   string          `json:"model"`
	Data    []EmbeddingData `json:"data"`
}

// =============================================================================

type FactualityCheck struct {
	Score float64 `json:"score"`
	Index int     `json:"index"`
}

type Factuality struct {
	ID      string            `json:"id"`
	Object  string            `json:"object"`
	Created Time              `json:"created"`
	Checks  []FactualityCheck `json:"checks"`
}

// =============================================================================

type InjectionCheck struct {
	Probability float64 `json:"probability"`
	Index       int     `json:"index"`
	Status      string  `json:"status"`
}

type Injection struct {
	ID      string           `json:"id"`
	Object  string           `json:"object"`
	Created Time             `json:"created"`
	Checks  []InjectionCheck `json:"checks"`
}

// =============================================================================

type ReplacePIICheck struct {
	NewPrompt string `json:"new_prompt"`
	Index     int    `json:"index"`
	Status    string `json:"status"`
}

type ReplacePII struct {
	ID      string            `json:"id"`
	Object  string            `json:"object"`
	Created Time              `json:"created"`
	Checks  []ReplacePIICheck `json:"checks"`
}

// =============================================================================

type RerankResult struct {
	Index          int     `json:"index"`
	RelevanceScore float64 `json:"relevance_score"`
	Text           string  `json:"text"`
}

type Rerank struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created Time           `json:"created"`
	Model   string         `json:"model"`
	Results []RerankResult `json:"results"`
}

// =============================================================================

type TokenData struct {
	ID    int    `json:"id"`
	Start int    `json:"start"`
	Stop  int    `json:"stop"`
	Text  string `json:"text"`
}

type Tokenize struct {
	ID      string      `json:"id"`
	Object  string      `json:"object"`
	Created Time        `json:"created"`
	Data    []TokenData `json:"data"`
}

// =============================================================================

type ToxicityCheck struct {
	Score float64 `json:"score"`
	Index int     `json:"index"`
}

type Toxicity struct {
	ID      string          `json:"id"`
	Object  string          `json:"object"`
	Created Time            `json:"created"`
	Checks  []ToxicityCheck `json:"checks"`
}

// =============================================================================

type Translation struct {
	Score       float64 `json:"score"`
	Translation string  `json:"translation"`
	Model       string  `json:"model"`
	Status      string  `json:"status"`
}

type Translate struct {
	ID                   string        `json:"id"`
	Object               string        `json:"object"`
	Created              Time          `json:"created"`
	BestTranslation      string        `json:"best_translation"`
	BestTranslationModel string        `json:"best_translation_model"`
	Score                float64       `json:"best_score"`
	Translations         []Translation `json:"translations"`
}
