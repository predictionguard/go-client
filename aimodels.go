package client

import "fmt"

// Models represents the set of models that can be used.
var Models = struct {
	MetaLlama38BInstruct     Model
	NousHermesLlama213B      Model
	Hermes2ProMistral7B      Model
	NeuralChat7B             Model
	Yi34BChat                Model
	DeepseekCoder67BInstruct Model
}{
	MetaLlama38BInstruct:     newModel("Meta-Llama-38B-Instruct"),
	NousHermesLlama213B:      newModel("Nous-Hermes-Llama-213B"),
	Hermes2ProMistral7B:      newModel("Hermes-2-Pro-Mistral-7B"),
	NeuralChat7B:             newModel("Neural-Chat-7B"),
	Yi34BChat:                newModel("Yi-34B-Chat"),
	DeepseekCoder67BInstruct: newModel("deepseek-coder-6.7b-instruct"),
}

// Set of known models.
var models = make(map[string]Model)

// Model represents a model in the system.
type Model struct {
	name string
}

func newModel(model string) Model {
	r := Model{model}
	models[model] = r
	return r
}

// ParseModel parses the string value and returns a model if one exists.
func ParseModel(value string) (Model, error) {
	model, exists := models[value]
	if !exists {
		return Model{}, fmt.Errorf("invalid Model %q", value)
	}

	return model, nil
}

// MustParseModel parses the string value and returns a model if one exists. If
// an error occurs the function panics.
func MustParseModel(value string) Model {
	Model, err := ParseModel(value)
	if err != nil {
		panic(err)
	}

	return Model
}

// Name returns the name of the Model.
func (r Model) Name() string {
	return r.name
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (r *Model) UnmarshalText(data []byte) error {
	Model, err := ParseModel(string(data))
	if err != nil {
		return err
	}

	r.name = Model.name
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (r Model) MarshalText() ([]byte, error) {
	return []byte(r.name), nil
}

// Equal provides support for the go-cmp package and testing.
func (r Model) Equal(r2 Model) bool {
	return r.name == r2.name
}
