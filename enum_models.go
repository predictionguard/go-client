package client

import "fmt"

type modelSet struct {
	BridgetowerLargeItmMlmItc Model
	DeepseekCoder67BInstruct  Model
	Hermes2ProLlama38B        Model
	Hermes2ProMistral7B       Model
	LLama3SqlCoder8b          Model
	Llava157BHF               Model
	NeuralChat7B              Model
	NousHermesLlama213B       Model
}

// Models represents the set of models that can be used.
var Models = modelSet{
	BridgetowerLargeItmMlmItc: newModel("bridgetower-large-itm-mlm-itc"),
	DeepseekCoder67BInstruct:  newModel("deepseek-coder-6.7b-instruct"),
	Hermes2ProLlama38B:        newModel("Hermes-2-Pro-Llama-3-8B"),
	Hermes2ProMistral7B:       newModel("Hermes-2-Pro-Mistral-7B"),
	LLama3SqlCoder8b:          newModel("llama-3-sqlcoder-8b"),
	Llava157BHF:               newModel("llava-1.5-7b-hf"),
	NeuralChat7B:              newModel("Neural-Chat-7B"),
	NousHermesLlama213B:       newModel("Nous-Hermes-Llama-213B"),
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

// String returns the name of the Model.
func (r Model) String() string {
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
