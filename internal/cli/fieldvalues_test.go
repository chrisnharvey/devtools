package cli

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestFieldValues_GetString(t *testing.T) {
	tests := []struct {
		name     string
		flagName string
		flagVal  string
		expected string
	}{
		{
			name:     "returns string value when flag exists",
			flagName: "test-flag",
			flagVal:  "test-value",
			expected: "test-value",
		},
		{
			name:     "returns empty string when flag doesn't exist",
			flagName: "nonexistent-flag",
			flagVal:  "",
			expected: "",
		},
		{
			name:     "returns empty string when flag exists but has no value",
			flagName: "empty-flag",
			flagVal:  "",
			expected: "",
		},
		{
			name:     "returns string with spaces",
			flagName: "space-flag",
			flagVal:  "value with spaces",
			expected: "value with spaces",
		},
		{
			name:     "returns string with special characters",
			flagName: "special-flag",
			flagVal:  "value@#$%^&*()",
			expected: "value@#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)

			if tt.flagName != "nonexistent-flag" {
				flagSet.String(tt.flagName, "", "test flag")
				if tt.flagVal != "" {
					err := flagSet.Set(tt.flagName, tt.flagVal)
					if err != nil {
						t.Error(err)
					}
				}
			}

			fv := NewFieldValues(flagSet)
			result := fv.GetString(tt.flagName)

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFieldValues_GetFile(t *testing.T) {
	tests := []struct {
		name     string
		flagName string
		flagVal  string
		expected string
	}{
		{
			name:     "returns file path when flag exists",
			flagName: "file-flag",
			flagVal:  "/path/to/file.txt",
			expected: "/path/to/file.txt",
		},
		{
			name:     "returns empty string when flag doesn't exist",
			flagName: "nonexistent-flag",
			flagVal:  "",
			expected: "",
		},
		{
			name:     "returns relative file path",
			flagName: "rel-file-flag",
			flagVal:  "./relative/path.txt",
			expected: "./relative/path.txt",
		},
		{
			name:     "returns file path with spaces",
			flagName: "space-file-flag",
			flagVal:  "/path/to/file with spaces.txt",
			expected: "/path/to/file with spaces.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)

			if tt.flagName != "nonexistent-flag" {
				flagSet.String(tt.flagName, "", "test file flag")
				if tt.flagVal != "" {
					err := flagSet.Set(tt.flagName, tt.flagVal)
					if err != nil {
						t.Error(err)
					}
				}
			}

			fv := NewFieldValues(flagSet)
			result := fv.GetFile(tt.flagName)

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFieldValues_GetBool(t *testing.T) {
	tests := []struct {
		name     string
		flagName string
		flagVal  bool
		setFlag  bool
		expected bool
	}{
		{
			name:     "returns true when flag is set to true",
			flagName: "bool-flag",
			flagVal:  true,
			setFlag:  true,
			expected: true,
		},
		{
			name:     "returns false when flag is set to false",
			flagName: "bool-flag",
			flagVal:  false,
			setFlag:  true,
			expected: false,
		},
		{
			name:     "returns false when flag doesn't exist",
			flagName: "nonexistent-flag",
			flagVal:  false,
			setFlag:  false,
			expected: false,
		},
		{
			name:     "returns false when flag exists but not set",
			flagName: "unset-bool-flag",
			flagVal:  false,
			setFlag:  false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)

			if tt.flagName != "nonexistent-flag" {
				flagSet.Bool(tt.flagName, false, "test bool flag")
				if tt.setFlag {
					var err error
					if tt.flagVal {
						err = flagSet.Set(tt.flagName, "true")
					} else {
						err = flagSet.Set(tt.flagName, "false")
					}

					if err != nil {
						t.Error(err)
					}
				}
			}

			fv := NewFieldValues(flagSet)
			result := fv.GetBool(tt.flagName)

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewFieldValues(t *testing.T) {
	tests := []struct {
		name    string
		flagSet *pflag.FlagSet
	}{
		{
			name:    "creates FieldValues with valid flagSet",
			flagSet: pflag.NewFlagSet("test", pflag.ContinueOnError),
		},
		{
			name:    "creates FieldValues with nil flagSet",
			flagSet: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := NewFieldValues(tt.flagSet)

			assert.NotNil(t, fv)
			assert.Equal(t, tt.flagSet, fv.flagSet)
		})
	}
}
