package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type HMAC struct{}

func New() *HMAC {
	return &HMAC{}
}

func (h *HMAC) Execute(values field.Values) error {
	message := values.GetString("message")
	key := values.GetString("key")
	algorithm := values.GetString("algorithm")

	if algorithm == "" {
		algorithm = "sha256"
	}

	// Get the hash function based on algorithm
	var hashFunc func() hash.Hash
	switch algorithm {
	case "md5":
		hashFunc = md5.New
	case "sha1":
		hashFunc = sha1.New
	case "sha256":
		hashFunc = sha256.New
	case "sha512":
		hashFunc = sha512.New
	default:
		return fmt.Errorf("unsupported algorithm: %s (supported: md5, sha1, sha256, sha512)", algorithm)
	}

	// Generate HMAC
	mac := hmac.New(hashFunc, []byte(key))
	mac.Write([]byte(message))
	result := mac.Sum(nil)

	// Output as hex string
	fmt.Printf("%x\n", result)
	return nil
}

func (h *HMAC) GetUse() string {
	return "hmac"
}

func (h *HMAC) GetDescription() string {
	return "Generate HMAC (Hash-based Message Authentication Code)"
}

func (h *HMAC) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "message",
			Type:        field.TextArea,
			Description: "Message to generate HMAC for",
			Required:    true,
		},
		{
			Name:        "key",
			Type:        field.String,
			Description: "Secret key for HMAC",
			Required:    true,
		},
		{
			Name:        "algorithm",
			Type:        field.String,
			Description: "Hash algorithm (md5, sha1, sha256, sha512, default: sha256)",
			Required:    false,
		},
	}
}
