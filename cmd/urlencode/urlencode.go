package urlencode

import (
	"fmt"
	"net/url"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type UrlEncode struct{}

func New() *UrlEncode {
	return &UrlEncode{}
}

func (u *UrlEncode) Execute(values field.Values) error {
	input := values.GetString("string")

	// URL encode the input string
	encoded := url.QueryEscape(input)

	// Output the encoded string
	fmt.Printf("%s\n", encoded)
	return nil
}

func (u *UrlEncode) GetName() string {
	return "urlencode"
}

func (u *UrlEncode) GetDescription() string {
	return "URL encode a string"
}

func (u *UrlEncode) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "string",
			Type:        field.TextArea,
			Description: "Input string to URL encode",
			Required:    true,
		},
	}
}
