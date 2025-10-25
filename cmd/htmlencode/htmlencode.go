package htmlencode

import (
	"fmt"
	"html"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type HTMLEncode struct{}

func New() *HTMLEncode {
	return &HTMLEncode{}
}

func (h *HTMLEncode) Execute(values field.Values) error {
	input := values.GetString("string")

	// HTML encode the input string
	encoded := html.EscapeString(input)

	fmt.Println(encoded)
	return nil
}

func (h *HTMLEncode) GetUse() string {
	return "htmlencode"
}

func (h *HTMLEncode) GetDescription() string {
	return "HTML encode a string (escape HTML entities)"
}

func (h *HTMLEncode) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "string",
			Type:        field.TextArea,
			Description: "Input string to HTML encode",
			Required:    true,
		},
	}
}
