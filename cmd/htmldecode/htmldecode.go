package htmldecode

import (
	"fmt"
	"html"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type HTMLDecode struct{}

func New() *HTMLDecode {
	return &HTMLDecode{}
}

func (h *HTMLDecode) Execute(values field.Values) error {
	input := values.GetString("string")

	// HTML decode the input string
	decoded := html.UnescapeString(input)

	fmt.Println(decoded)
	return nil
}

func (h *HTMLDecode) GetUse() string {
	return "htmldecode"
}

func (h *HTMLDecode) GetDescription() string {
	return "HTML decode a string (unescape HTML entities)"
}

func (h *HTMLDecode) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "string",
			Type:        field.TextArea,
			Description: "HTML encoded string to decode",
			Required:    true,
		},
	}
}
