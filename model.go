package client

// Error represents an error in the system.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface.
func (err *Error) Error() string {
	return err.Message
}
