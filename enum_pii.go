package client

import "fmt"

type piiSet struct {
	Block   PII
	Replace PII
}

// PIIs represents the set of PIIs that can be used.
var PIIs = piiSet{
	Block:   newPII("block"),
	Replace: newPII("replace"),
}

// Parse parses the string value and returns a PII if one exists.
func (piiSet) Parse(value string) (PII, error) {
	pii, exists := piis[value]
	if !exists {
		return PII{}, fmt.Errorf("invalid pii %q", value)
	}

	return pii, nil
}

// MustParse parses the string value and returns a pii if one exists. If
// an error occurs the function panics.
func (piiSet) MustParse(value string) PII {
	pii, err := PIIs.Parse(value)
	if err != nil {
		panic(err)
	}

	return pii
}

// =============================================================================

// Set of known PIIs.
var piis = make(map[string]PII)

// PII represents a PII in the system.
type PII struct {
	name string
}

func newPII(pii string) PII {
	p := PII{pii}
	piis[pii] = p
	return p
}

// String returns the name of the PII.
func (p PII) String() string {
	return p.name
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (p *PII) UnmarshalText(data []byte) error {
	pii, err := PIIs.Parse(string(data))
	if err != nil {
		return err
	}

	p.name = pii.name
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (p PII) MarshalText() ([]byte, error) {
	return []byte(p.name), nil
}

// Equal provides support for the go-cmp package and testing.
func (p PII) Equal(p2 PII) bool {
	return p.name == p2.name
}
