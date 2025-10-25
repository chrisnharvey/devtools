package charcount

import (
	"fmt"
	"unicode/utf8"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type CharCount struct{}

func New() *CharCount {
	return &CharCount{}
}

func (c *CharCount) Execute(values field.Values) error {
	input := values.GetString("string")

	// Count characters (runes) in the string
	charCount := utf8.RuneCountInString(input)

	// Output the character count
	fmt.Printf("%d\n", charCount)
	return nil
}

func (c *CharCount) GetUse() string {
	return "charcount"
}

func (c *CharCount) GetDescription() string {
	return "Count the number of characters in a string"
}

func (c *CharCount) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "string",
			Type:        field.TextArea,
			Description: "Input string to count characters",
			Required:    true,
		},
	}
}
