package client

import "fmt"

type replaceMethodSet struct {
	Random   ReplaceMethod
	Fake     ReplaceMethod
	Category ReplaceMethod
	Mask     ReplaceMethod
}

// ReplaceMethods represents the set of replace methods that can be used.
var ReplaceMethods = replaceMethodSet{
	Random:   newReplaceMethod("random"),
	Fake:     newReplaceMethod("fake"),
	Category: newReplaceMethod("category"),
	Mask:     newReplaceMethod("mask"),
}

// Parse parses the string value and returns a replace method if one exists.
func (replaceMethodSet) Parse(value string) (ReplaceMethod, error) {
	replaceMethod, exists := replaceMethods[value]
	if !exists {
		return ReplaceMethod{}, fmt.Errorf("invalid replace method %q", value)
	}

	return replaceMethod, nil
}

// MustParse parses the string value and returns a replace method if one exists.
// If an error occurs the function panics.
func (replaceMethodSet) MustParse(value string) ReplaceMethod {
	replaceMethod, err := ReplaceMethods.Parse(value)
	if err != nil {
		panic(err)
	}

	return replaceMethod
}

// =============================================================================

// Set of known replace methods.
var replaceMethods = make(map[string]ReplaceMethod)

// ReplaceMethod represents a replace method in the system.
type ReplaceMethod struct {
	name string
}

func newReplaceMethod(replaceMethod string) ReplaceMethod {
	rm := ReplaceMethod{replaceMethod}
	replaceMethods[replaceMethod] = rm
	return rm
}

// String returns the name of the replace method.
func (rm ReplaceMethod) String() string {
	return rm.name
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (rm *ReplaceMethod) UnmarshalText(data []byte) error {
	replaceMethod, err := ReplaceMethods.Parse(string(data))
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
