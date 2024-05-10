package client

import "fmt"

// Set of known replace methods.
var replaceMethods = make(map[string]ReplaceMethod)

// Set of possible replace methods.
var (
	ReplaceMethodUser      = newReplaceMethod("user")
	ReplaceMethodAssistant = newReplaceMethod("assistant")
)

// ReplaceMethod represents a replace method in the system.
type ReplaceMethod struct {
	name string
}

func newReplaceMethod(replaceMethod string) ReplaceMethod {
	rm := ReplaceMethod{replaceMethod}
	replaceMethods[replaceMethod] = rm
	return rm
}

// ParseReplaceMethod parses the string value and returns a replace method
// if one exists.
func ParseReplaceMethod(value string) (ReplaceMethod, error) {
	replaceMethod, exists := replaceMethods[value]
	if !exists {
		return ReplaceMethod{}, fmt.Errorf("invalid replace method %q", value)
	}

	return replaceMethod, nil
}

// MustParseReplaceMethod parses the string value and returns a replace method
// if one exists. If an error occurs the function panics.
func MustParseReplaceMethod(value string) Role {
	role, err := ParseRole(value)
	if err != nil {
		panic(err)
	}

	return role
}

// Name returns the name of the role.
func (rm ReplaceMethod) Name() string {
	return rm.name
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (rm *ReplaceMethod) UnmarshalText(data []byte) error {
	replaceMethod, err := ParseReplaceMethod(string(data))
	if err != nil {
		return err
	}

	rm.name = replaceMethod.name
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (rm ReplaceMethod) MarshalText() ([]byte, error) {
	return []byte(rm.name), nil
}

// Equal provides support for the go-cmp package and testing.
func (rm ReplaceMethod) Equal(rm2 ReplaceMethod) bool {
	return rm.name == rm2.name
}
