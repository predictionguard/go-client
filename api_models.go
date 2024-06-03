package client

import (
	"context"
	"strconv"
	"strings"
	"time"
)

// Base64Encoder defines a method that can read a data source and returns a
// base64 encoded string.
type Base64Encoder interface {
	EncodeBase64(ctx context.Context) (string, error)
}

// =============================================================================

// Error represents an error in the system.
type Error struct {
	Message string `json:"error"`
}

// Error implements the error interface.
func (err *Error) Error() string {
	return err.Message
}

// =============================================================================

// Time supports the ability to marshal and unmarshal unit epoch time.
type Time struct {
	time.Time
}

// ToTime returns the local Time corresponding to the given Unix time, sec seconds
// and nsec nanoseconds since January 1, 1970 UTC. It is valid to pass nsec
// outside the range [0, 999999999]. Not all sec values have a corresponding time
// value. One such value is 1<<63-1 (the largest int64 value).
func ToTime(sec int64) Time {
	return Time{
		Time: time.Unix(sec, 0),
	}
}

// UnmarshalJSON overrides the time.Time implementation so we can unmarshal
// from epoch time.
func (t *Time) UnmarshalJSON(data []byte) error {
	d := strings.Trim(string(data), "\"")

	num, err := strconv.Atoi(d)
	if err != nil {
		return err
	}

	t.Time = time.Unix(int64(num), 0)

	return nil
}

// MarshalJSON overrides the time.Time implementation so we can marshal
// to epoch time.
func (t Time) MarshalJSON() ([]byte, error) {
	data := strconv.Itoa(int(t.Unix()))
	return []byte(data), nil
}
