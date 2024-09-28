package client

import "fmt"

type capabilitySet struct {
	Completion           Capability
	ChatCompletion       Capability
	ChatCompletionVision Capability
	Embedding            Capability
}

// Capabilities represents the set of model capabilities.
var Capabilities = capabilitySet{
	Completion:           newCapability("completion"),
	ChatCompletion:       newCapability("chat_completion"),
	ChatCompletionVision: newCapability("chat_completion_vision"),
	Embedding:            newCapability("embedding"),
}

// Parse parses the string value and returns a capability if one exists.
func (capabilitySet) Parse(value string) (Capability, error) {
	capability, exists := capabilities[value]
	if !exists {
		return Capability{}, fmt.Errorf("invalid capability %q", value)
	}

	return capability, nil
}

// MustParse parses the string value and returns a capability if one exists.
// If an error occurs the function panics.
func (capabilitySet) MustParse(value string) Capability {
	capability, err := Capabilities.Parse(value)
	if err != nil {
		panic(err)
	}

	return capability
}

// =============================================================================

// Set of known capabilities.
var capabilities = make(map[string]Capability)

// Capability represents a capability in the system.
type Capability struct {
	value string
}

func newCapability(capability string) Capability {
	c := Capability{capability}
	capabilities[capability] = c
	return c
}

// String returns the name of the capability.
func (c Capability) String() string {
	return c.value
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (c *Capability) UnmarshalText(data []byte) error {
	capability, err := Capabilities.Parse(string(data))
	if err != nil {
		return err
	}

	c.value = capability.value
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (c Capability) MarshalText() ([]byte, error) {
	return []byte(c.value), nil
}

// Equal provides support for the go-cmp package and testing.
func (c Capability) Equal(c2 Capability) bool {
	return c.value == c2.value
}
