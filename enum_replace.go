package client

import "fmt"

type replaceMethodSet struct {
	Random   ReplaceMethod
	Fake     ReplaceMethod
	Category ReplaceMethod
	Mask     ReplaceMethod
}

var ReplaceMethods = replaceMethodSet{
	Random:   newReplaceMethod("random"),
	Fake:     newReplaceMethod("fake"),
	Category: newReplaceMethod("category"),
	Mask:     newReplaceMethod("mask"),
}

func (replaceMethodSet) Parse(value string) (ReplaceMethod, error) {
	replaceMethod, exists := replaceMethods[value]
	if !exists {
		return ReplaceMethod{}, fmt.Errorf("invalid replace method %q", value)
	}

	return replaceMethod, nil
}

func (replaceMethodSet) MustParse(value string) ReplaceMethod {
	replaceMethod, err := ReplaceMethods.Parse(value)
	if err != nil {
		panic(err)
	}

	return replaceMethod
}

// =============================================================================

var replaceMethods = make(map[string]ReplaceMethod)

type ReplaceMethod struct {
	value string
}

func newReplaceMethod(replaceMethod string) ReplaceMethod {
	rm := ReplaceMethod{replaceMethod}
	replaceMethods[replaceMethod] = rm
	return rm
}

func (rm ReplaceMethod) String() string {
	return rm.value
}

func (rm *ReplaceMethod) UnmarshalText(data []byte) error {
	replaceMethod, err := ReplaceMethods.Parse(string(data))
	if err != nil {
		return err
	}

	rm.value = replaceMethod.value
	return nil
}

func (rm ReplaceMethod) MarshalText() ([]byte, error) {
	return []byte(rm.value), nil
}

func (rm ReplaceMethod) Equal(rm2 ReplaceMethod) bool {
	return rm.value == rm2.value
}
