package jsonprettify

import (
	"encoding/json"
	"fmt"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type JsonPrettify struct{}

func (j *JsonPrettify) GetName() string {
	return "JSON Prettify"
}

func New() *JsonPrettify {
	return &JsonPrettify{}
}

func (j *JsonPrettify) Execute(values field.Values) error {
	input := values.GetString("json")

	// Parse JSON to validate it
	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	// Marshal back to prettified JSON with indentation
	prettified, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to prettify JSON: %w", err)
	}

	// Output the prettified JSON
	fmt.Printf("%s\n", prettified)
	return nil
}

func (j *JsonPrettify) GetUse() string {
	return "jsonprettify"
}

func (j *JsonPrettify) GetDescription() string {
	return "Prettify JSON with proper indentation and formatting"
}

func (j *JsonPrettify) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "json",
			Type:        field.TextArea,
			Description: "JSON to prettify",
			Required:    true,
		},
	}
}
