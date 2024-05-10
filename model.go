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
// Chat

// ChatMessage represents the role of the sender and the content to process.
type ChatMessage struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// Choice represent the choices that are provided for you to choose from.
type ChatChoice struct {
	Index   int         `json:"index"`
	Message ChatMessage `json:"message"`
}

// Chat represents the result for the chat completion call.
type Chat struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created Time         `json:"created"`
	Model   string       `json:"model"`
	Choices []ChatChoice `json:"choices"`
}

// ChoiceSSE represent the choices that are provided for you to choose from.
type ChatChoiceSSE struct {
	Index int `json:"index"`
	Delta struct {
		Content string `json:"content"`
	} `json:"delta"`
	Text         string  `json:"generated_text"`
	Probs        float32 `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

// ChatSSE represents the result for the chat completion call.
type ChatSSE struct {
	Chat
	Choices []ChatChoiceSSE `json:"choices"`
}

// =============================================================================
// Completion

// CompletionChoice represent the choices that are provided for you to choose from.
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
