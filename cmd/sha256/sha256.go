package sha256

import (
	"crypto/sha256"
	"fmt"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type Sha256 struct{}

func New() *Sha256 {
	return &Sha256{}
}

func (s *Sha256) Execute(values field.Values) error {
	input := values.GetString("string")

	// Generate SHA256 hash
	hash := sha256.Sum256([]byte(input))

	// Convert to hex string and output
	fmt.Printf("%x\n", hash)
	return nil
}

func (s *Sha256) GetName() string {
	return "sha256"
}

func (s *Sha256) GetDescription() string {
	return "Generate SHA256 hash from string or file input"
}

func (s *Sha256) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "string",
			Type:        field.TextArea,
			Description: "Input string",
			Required:    true,
		},
	}
}
