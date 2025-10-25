package base64decode

import (
	"encoding/base64"
	"fmt"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type Base64Decode struct{}

func (b *Base64Decode) GetName() string {
	return "Base64 Decode"
}

func New() *Base64Decode {
	return &Base64Decode{}
}

func (b *Base64Decode) Execute(values field.Values) error {
	input := values.GetString("string")

	// Base64 decode the input string
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return fmt.Errorf("failed to decode base64: %w", err)
	}

	fmt.Println(string(decoded))
	return nil
}

func (b *Base64Decode) GetUse() string {
	return "base64decode"
}

func (b *Base64Decode) GetDescription() string {
	return "Base64 decode a string"
}

func (b *Base64Decode) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "string",
			Type:        field.TextArea,
			Description: "Base64 encoded string to decode",
			Required:    true,
		},
	}
}
