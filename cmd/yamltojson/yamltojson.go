package yamltojson

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type YAMLToJSON struct{}

func New() *YAMLToJSON {
	return &YAMLToJSON{}
}

func (y *YAMLToJSON) Execute(values field.Values) error {
	input := values.GetString("yaml")

	// Parse YAML
	var data interface{}
	if err := yaml.Unmarshal([]byte(input), &data); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Convert to JSON with pretty formatting
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %w", err)
	}

	fmt.Println(string(jsonBytes))
	return nil
}

func (y *YAMLToJSON) GetUse() string {
	return "yamltojson"
}

func (y *YAMLToJSON) GetDescription() string {
	return "Convert YAML to JSON format"
}

func (y *YAMLToJSON) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "yaml",
			Type:        field.TextArea,
			Description: "YAML content to convert to JSON",
			Required:    true,
		},
	}
}
