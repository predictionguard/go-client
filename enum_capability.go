package client

import "fmt"

type capabilitySet struct {
	ChatCompletion     Capability
	ChatWithImage      Capability
	Completion         Capability
	Embedding          Capability
	EmbeddingWithImage Capability
	Tokenize           Capability
}

// Capabilities represents the set of model capabilities.
var Capabilities = capabilitySet{
	ChatCompletion:     newCapability("chat-completion"),
	ChatWithImage:      newCapability("chat-with-image"),
	Completion:         newCapability("completion"),
	Embedding:          newCapability("embedding"),
	EmbeddingWithImage: newCapability("embedding-with-image"),
	Tokenize:           newCapability("tokenize"),
}

func (capabilitySet) Parse(value string) (Capability, error) {
	capability, exists := capabilities[value]
	if !exists {
		return Capability{}, fmt.Errorf("invalid capability %q", value)
	}

	return capability, nil
}

func (capabilitySet) MustParse(value string) Capability {
	capability, err := Capabilities.Parse(value)
	if err != nil {
		panic(err)
	}

	return capability
}

// =============================================================================

var capabilities = make(map[string]Capability)

type Capability struct {
	value string
}

func newCapability(capability string) Capability {
	c := Capability{capability}
	capabilities[capability] = c
	return c
}

func (c Capability) String() string {
	return c.value
}

func (c *Capability) UnmarshalText(data []byte) error {
	capability, err := Capabilities.Parse(string(data))
	if err != nil {
		return err
	}

	c.value = capability.value
	return nil
}

func (c Capability) MarshalText() ([]byte, error) {
	return []byte(c.value), nil
}

func (c Capability) Equal(c2 Capability) bool {
	return c.value == c2.value
}
