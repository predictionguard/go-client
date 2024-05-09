package client

import (
	"strconv"
	"time"
)

// Time supports the ability to marshal and unmarshal unit epoch time.
type Time struct {
	time.Time
}

// UnmarshalJSON overrides the time.Time implementation so we can unmarshal
// from epoch time.
func (t *Time) UnmarshalJSON(data []byte) error {
	num, err := strconv.Atoi(string(data))
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
