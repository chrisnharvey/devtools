package jsontoyaml

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type JSONToYAML struct{}

func (j *JSONToYAML) GetName() string {
	return "JSON to YAML"
}

func New() *JSONToYAML {
	return &JSONToYAML{}
}

func (j *JSONToYAML) Execute(values field.Values) error {
	input := values.GetString("json")

	// Parse JSON
	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Convert to YAML
	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to convert to YAML: %w", err)
	}

	fmt.Print(string(yamlBytes))
	return nil
}

func (j *JSONToYAML) GetUse() string {
	return "jsontoyaml"
}

func (j *JSONToYAML) GetDescription() string {
	return "Convert JSON to YAML format"
}

func (j *JSONToYAML) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "json",
			Type:        field.TextArea,
			Description: "JSON content to convert to YAML",
			Required:    true,
		},
	}
}
