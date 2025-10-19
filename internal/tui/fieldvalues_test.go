package tui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldValues_GetString(t *testing.T) {
	tests := []struct {
		name     string
		values   map[string]any
		key      string
		expected string
	}{
		{
			name:     "returns string value when key exists",
			values:   map[string]any{"test-key": "test-value"},
			key:      "test-key",
			expected: "test-value",
		},
		{
			name:     "returns empty string when key doesn't exist",
			values:   map[string]any{},
			key:      "nonexistent-key",
			expected: "",
		},
		{
			name:     "returns empty string when key exists but value is empty string",
			values:   map[string]any{"empty-key": ""},
			key:      "empty-key",
			expected: "",
		},
		{
			name:     "returns string with spaces",
			values:   map[string]any{"space-key": "value with spaces"},
			key:      "space-key",
			expected: "value with spaces",
		},
		{
			name:     "returns string with special characters",
			values:   map[string]any{"special-key": "value@#$%^&*()"},
			key:      "special-key",
			expected: "value@#$%^&*()",
		},
		{
			name:     "returns multiline string",
			values:   map[string]any{"multiline-key": "line1\nline2\nline3"},
			key:      "multiline-key",
			expected: "line1\nline2\nline3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := &FieldValues{values: tt.values}
			result := fv.GetString(tt.key)

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFieldValues_GetFile(t *testing.T) {
	tests := []struct {
		name     string
		values   map[string]any
		key      string
		expected string
	}{
		{
			name:     "returns file path when key exists",
			values:   map[string]any{"file-key": "/path/to/file.txt"},
			key:      "file-key",
			expected: "/path/to/file.txt",
		},
		{
			name:     "returns empty string when key doesn't exist",
			values:   map[string]any{},
			key:      "nonexistent-key",
			expected: "",
		},
		{
			name:     "returns relative file path",
			values:   map[string]any{"rel-file-key": "./relative/path.txt"},
			key:      "rel-file-key",
			expected: "./relative/path.txt",
		},
		{
			name:     "returns file path with spaces",
			values:   map[string]any{"space-file-key": "/path/to/file with spaces.txt"},
			key:      "space-file-key",
			expected: "/path/to/file with spaces.txt",
		},
		{
			name:     "returns home directory path",
			values:   map[string]any{"home-file-key": "~/documents/file.txt"},
			key:      "home-file-key",
			expected: "~/documents/file.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := &FieldValues{values: tt.values}
			result := fv.GetFile(tt.key)

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFieldValues_GetBool(t *testing.T) {
	tests := []struct {
		name     string
		values   map[string]any
		key      string
		expected bool
	}{
		{
			name:     "returns true when key exists and value is true",
			values:   map[string]any{"bool-key": true},
			key:      "bool-key",
			expected: true,
		},
		{
			name:     "returns false when key exists and value is false",
			values:   map[string]any{"bool-key": false},
			key:      "bool-key",
			expected: false,
		},
		{
			name:     "returns false when key doesn't exist",
			values:   map[string]any{},
			key:      "nonexistent-key",
			expected: false,
		},
		{
			name:     "returns false when key exists but map is nil",
			values:   nil,
			key:      "any-key",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := &FieldValues{values: tt.values}
			result := fv.GetBool(tt.key)

			assert.Equal(t, tt.expected, result)
		})
	}
}
