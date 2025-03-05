package client

import "fmt"

type piiSet struct {
	Block   PII
	Replace PII
}

var PIIs = piiSet{
	Block:   newPII("block"),
	Replace: newPII("replace"),
}

func (piiSet) Parse(value string) (PII, error) {
	pii, exists := piis[value]
	if !exists {
		return PII{}, fmt.Errorf("invalid pii %q", value)
	}

	return pii, nil
}

func (piiSet) MustParse(value string) PII {
	pii, err := PIIs.Parse(value)
	if err != nil {
		panic(err)
	}

	return pii
}

// =============================================================================

var piis = make(map[string]PII)

type PII struct {
	value string
}

func newPII(pii string) PII {
	p := PII{pii}
	piis[pii] = p
	return p
}

func (p PII) String() string {
	return p.value
}

func (p *PII) UnmarshalText(data []byte) error {
	pii, err := PIIs.Parse(string(data))
	if err != nil {
		return err
	}

	p.value = pii.value
	return nil
}

func (p PII) MarshalText() ([]byte, error) {
	return []byte(p.value), nil
}

func (p PII) Equal(p2 PII) bool {
	return p.value == p2.value
}
