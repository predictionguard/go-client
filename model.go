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

// Message represents the role of the sender and the content to process.
type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// Choice represent the choices that are provided for you to choose from.
type Choice struct {
	Message Message `json:"message"`
	Index   int     `json:"index"`
}

// ChatCompletionRequest represents the result for the chat completion call.
type ChatCompletion struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created Time     `json:"created"`
	Choices []Choice `json:"choices"`
}
