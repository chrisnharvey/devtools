package timestamp

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chrisnharvey/devtools/pkg/field"
)

type Timestamp struct{}

func (t *Timestamp) GetName() string {
	return "Timestamp Converter"
}

func New() *Timestamp {
	return &Timestamp{}
}

func (t *Timestamp) Execute(values field.Values) error {
	input := values.GetString("timestamp")
	timezone := values.GetString("timezone")

	if input == "" {
		// If no timestamp provided, use current time
		now := time.Now()
		fmt.Printf("Current Unix timestamp: %d\n", now.Unix())
		fmt.Printf("Current time (UTC): %s\n", now.UTC().Format("2006-01-02 15:04:05 MST"))
		fmt.Printf("Current time (Local): %s\n", now.Format("2006-01-02 15:04:05 MST"))
		return nil
	}

	// Parse the timestamp
	ts, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp format: %w", err)
	}

	// Convert to time.Time
	var parsedTime time.Time
	if ts > 1e10 {
		// Milliseconds timestamp
		parsedTime = time.Unix(ts/1000, (ts%1000)*1e6)
	} else {
		// Seconds timestamp
		parsedTime = time.Unix(ts, 0)
	}

	// Set timezone if specified
	if timezone != "" {
		loc, err := time.LoadLocation(timezone)
		if err != nil {
			return fmt.Errorf("invalid timezone: %w", err)
		}
		parsedTime = parsedTime.In(loc)
	}

	// Display results
	fmt.Printf("Unix timestamp: %d\n", ts)
	fmt.Printf("UTC time: %s\n", parsedTime.UTC().Format("2006-01-02 15:04:05 MST"))
	fmt.Printf("Local time: %s\n", parsedTime.Format("2006-01-02 15:04:05 MST"))

	if timezone != "" {
		fmt.Printf("Time in %s: %s\n", timezone, parsedTime.Format("2006-01-02 15:04:05 MST"))
	}

	// Additional formats
	fmt.Printf("ISO 8601: %s\n", parsedTime.UTC().Format("2006-01-02T15:04:05Z"))
	fmt.Printf("RFC 3339: %s\n", parsedTime.UTC().Format(time.RFC3339))
	fmt.Printf(
		"Human readable: %s\n",
		parsedTime.UTC().Format("Monday, January 2, 2006 at 3:04:05 PM MST"),
	)

	return nil
}

func (t *Timestamp) GetUse() string {
	return "timestamp"
}

func (t *Timestamp) GetDescription() string {
	return "Convert Unix timestamp to human readable time (or show current time if no timestamp provided)"
}

func (t *Timestamp) GetFields() []field.Field {
	return []field.Field{
		{
			Name:        "timestamp",
			Type:        field.String,
			Description: "Unix timestamp (seconds or milliseconds, leave empty for current time)",
			Required:    false,
		},
		{
			Name:        "timezone",
			Type:        field.String,
			Description: "Timezone (e.g., 'America/New_York', 'Europe/London', optional)",
			Required:    false,
		},
	}
}
