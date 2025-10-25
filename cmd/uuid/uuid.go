package uuid

import (
	"fmt"
	"strconv"

	"github.com/chrisnharvey/devtools/pkg/field"
	"github.com/google/uuid"
)

type UUID struct{}

func (u *UUID) GetName() string {
	return "UUID Generator"
}

func New() *UUID {
	return &UUID{}
}

func (u *UUID) Execute(values field.Values) error {
	countStr := values.GetString("count")
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		count = 1
	}

	version := values.GetString("version")
	if version == "" {
		version = "4"
	}

	for i := 0; i < count; i++ {
		var id uuid.UUID
		var err error

		switch version {
		case "1":
			id, err = uuid.NewUUID()
		case "4":
			id = uuid.New()
		default:
			id = uuid.New()
		}

		if err != nil {
			return fmt.Errorf("failed to generate UUID: %w", err)
		}

		fmt.Println(id.String())
	}

	return nil
}

func (u *UUID) GetUse() string {
	return "uuid"
}

func (u *UUID) GetDescription() string {
	return "Generate UUID (Universally Unique Identifier)"
}

func (u *UUID) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "version",
			Type:        field.String,
			Description: "UUID version (1 for timestamp-based, 4 for random, default: 4)",
			Required:    false,
		},
		{
			Name:        "count",
			Type:        field.String,
			Description: "Number of UUIDs to generate (default: 1)",
			Required:    false,
		},
	}
}
