package client

// Error represents an error in the system.
type Error struct {
	Message string `json:"error"`
}

// Error implements the error interface.
func (err *Error) Error() string {
	return err.Message
}

// =============================================================================
// ChatCompletion

// Message represents the role of the sender and the content to process.
type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// Choice represent the choices that are provided for you to choose from.
type ChatCompletionChoice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
}

// ChatCompletionRequest represents the result for the chat completion call.
type ChatCompletion struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created Time                   `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
}

// ChoiceSSE represent the choices that are provided for you to choose from.
type ChatCompletionChoiceSSE struct {
	Index int `json:"index"`
	Delta struct {
		Content string `json:"content"`
	} `json:"delta"`
	Text         string  `json:"generated_text"`
	Probs        float32 `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

// ChatCompletionSSE represents the result for the chat completion call.
type ChatCompletionSSE struct {
	ChatCompletion
	Choices []ChatCompletionChoiceSSE `json:"choices"`
}

// =============================================================================
// Completion

// ChoiceSSE represent the choices that are provided for you to choose from.
type CompletionChoice struct {
	Text   string `json:"text"`
	Index  int    `json:"index"`
	Status string `json:"status"`
	Model  string `json:"model"`
}

// CompletionRequest represents the result for the completion call.
type Completion struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created Time               `json:"created"`
	Choices []CompletionChoice `json:"choices"`
}
