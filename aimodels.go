package client

import "fmt"

type modelSet struct {
	MetaLlama38BInstruct     Model
	NousHermesLlama213B      Model
	Hermes2ProMistral7B      Model
	NeuralChat7B             Model
	Yi34BChat                Model
	DeepseekCoder67BInstruct Model
}

// Models represents the set of models that can be used.
var Models = modelSet{
	MetaLlama38BInstruct:     newModel("Meta-Llama-38B-Instruct"),
	NousHermesLlama213B:      newModel("Nous-Hermes-Llama-213B"),
	Hermes2ProMistral7B:      newModel("Hermes-2-Pro-Mistral-7B"),
	NeuralChat7B:             newModel("Neural-Chat-7B"),
	Yi34BChat:                newModel("Yi-34B-Chat"),
	DeepseekCoder67BInstruct: newModel("deepseek-coder-6.7b-instruct"),
}

// Parse parses the string value and returns a model if one exists.
func (modelSet) Parse(value string) (Model, error) {
	model, exists := models[value]
	if !exists {
		return Model{}, fmt.Errorf("invalid Model %q", value)
	}

	return model, nil
}

// MustParseModel parses the string value and returns a model if one exists. If
// an error occurs the function panics.
func (modelSet) MustParse(value string) Model {
	Model, err := Models.Parse(value)
	if err != nil {
		panic(err)
	}

	return Model
}

// =============================================================================

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

// Name returns the name of the Model.
func (r Model) Name() string {
	return r.name
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (r *Model) UnmarshalText(data []byte) error {
	Model, err := Models.Parse(string(data))
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
