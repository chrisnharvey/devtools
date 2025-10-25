package base64encode

import (
	"encoding/base64"
	"fmt"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type Base64Encode struct{}

func (b *Base64Encode) GetName() string {
	return "Base64 Encode"
}

func New() *Base64Encode {
	return &Base64Encode{}
}

func (b *Base64Encode) Execute(values field.Values) error {
	input := values.GetString("string")

	// Base64 encode the input string
	encoded := base64.StdEncoding.EncodeToString([]byte(input))

	fmt.Println(encoded)
	return nil
}

func (b *Base64Encode) GetUse() string {
	return "base64encode"
}

func (b *Base64Encode) GetDescription() string {
	return "Base64 encode a string"
}

func (b *Base64Encode) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "string",
			Type:        field.TextArea,
			Description: "Input string to encode",
			Required:    true,
		},
	}
}
