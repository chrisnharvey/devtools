package jsonminify

import (
	"encoding/json"
	"fmt"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type JsonMinify struct{}

func New() *JsonMinify {
	return &JsonMinify{}
}

func (j *JsonMinify) Execute(values field.Values) error {
	input := values.GetString("json")

	// Parse JSON to validate it
	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	// Marshal back to minified JSON
	minified, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to minify JSON: %w", err)
	}

	// Output the minified JSON
	fmt.Printf("%s\n", minified)
	return nil
}

func (j *JsonMinify) GetName() string {
	return "jsonminify"
}

func (j *JsonMinify) GetDescription() string {
	return "Minify JSON by removing whitespace and formatting"
}

func (j *JsonMinify) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "json",
			Type:        field.TextArea,
			Description: "JSON to minify",
			Required:    true,
		},
	}
}
