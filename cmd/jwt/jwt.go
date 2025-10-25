package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type JWT struct{}

func New() *JWT {
	return &JWT{}
}

func (j *JWT) Execute(values field.Values) error {
	token := values.GetString("token")

	// Split the JWT into its three parts
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return fmt.Errorf("invalid JWT format: expected 3 parts separated by dots, got %d", len(parts))
	}

	// Decode header
	header, err := j.decodeJWTPart(parts[0])
	if err != nil {
		return fmt.Errorf("failed to decode JWT header: %w", err)
	}

	// Decode payload
	payload, err := j.decodeJWTPart(parts[1])
	if err != nil {
		return fmt.Errorf("failed to decode JWT payload: %w", err)
	}

	// Pretty print the results
	fmt.Println("Header:")
	fmt.Println(header)
	fmt.Println()
	fmt.Println("Payload:")
	fmt.Println(payload)
	fmt.Println()
	fmt.Println("Signature:")
	fmt.Println(parts[2])

	return nil
}

func (j *JWT) decodeJWTPart(part string) (string, error) {
	// Add padding if necessary (JWT uses base64url encoding without padding)
	switch len(part) % 4 {
	case 2:
		part += "=="
	case 3:
		part += "="
	}

	// Replace URL-safe characters with standard base64 characters
	part = strings.ReplaceAll(part, "-", "+")
	part = strings.ReplaceAll(part, "_", "/")

	// Decode base64
	decoded, err := base64.StdEncoding.DecodeString(part)
	if err != nil {
		return "", err
	}

	// Pretty print JSON
	var jsonData interface{}
	if err := json.Unmarshal(decoded, &jsonData); err != nil {
		// If it's not valid JSON, return as string
		return string(decoded), nil
	}

	prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return string(decoded), nil
	}

	return string(prettyJSON), nil
}

func (j *JWT) GetName() string {
	return "jwt"
}

func (j *JWT) GetDescription() string {
	return "Decode JWT (JSON Web Token) and display header, payload, and signature"
}

func (j *JWT) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "token",
			Type:        field.TextArea,
			Description: "JWT token to decode",
			Required:    true,
		},
	}
}
