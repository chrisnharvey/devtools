package ulid

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type ULID struct{}

func (u *ULID) GetName() string {
	return "ULID Generator"
}

func New() *ULID {
	return &ULID{}
}

func (u *ULID) Execute(values field.Values) error {
	countStr := values.GetString("count")
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		count = 1
	}

	// Use current time as the timestamp for ULIDs
	timestamp := time.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)

	for i := 0; i < count; i++ {
		id := ulid.MustNew(ulid.Timestamp(timestamp), entropy)
		fmt.Println(id.String())
	}

	return nil
}

func (u *ULID) GetUse() string {
	return "ulid"
}

func (u *ULID) GetDescription() string {
	return "Generate ULID (Universally Unique Lexicographically Sortable Identifier)"
}

func (u *ULID) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "count",
			Type:        field.String,
			Description: "Number of ULIDs to generate (default: 1)",
			Required:    false,
		},
	}
}
