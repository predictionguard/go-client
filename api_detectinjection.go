package client

// DetectInjection need description.
type DetectInjection struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created Time   `json:"created"`
	Checks  []struct {
		Probability float64 `json:"probability"`
		Index       int     `json:"index"`
		Status      string  `json:"status"`
	} `json:"checks"`
}
