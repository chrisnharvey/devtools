package urldecode

import (
	"fmt"
	"net/url"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type UrlDecode struct{}

func New() *UrlDecode {
	return &UrlDecode{}
}

func (u *UrlDecode) Execute(values field.Values) error {
	input := values.GetString("string")

	// URL decode the input string
	decoded, err := url.QueryUnescape(input)
	if err != nil {
		return fmt.Errorf("failed to decode URL: %w", err)
	}

	// Output the decoded string
	fmt.Printf("%s\n", decoded)
	return nil
}

func (u *UrlDecode) GetUse() string {
	return "urldecode"
}

func (u *UrlDecode) GetDescription() string {
	return "URL decode a string"
}

func (u *UrlDecode) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "string",
			Type:        field.TextArea,
			Description: "URL encoded string to decode",
			Required:    true,
		},
	}
}
