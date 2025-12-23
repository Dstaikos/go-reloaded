package pkg

import (
	"testing"
)

func TestUpLowCap(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple uppercase",
			input:    "Ready, set, go (up)!",
			expected: "Ready, set, GO!",
		},
		{
			name:     "Simple lowercase",
			input:    "I should stop SHOUTING (low)",
			expected: "I should stop shouting",
		},
		{
			name:     "Simple capitalize",
			input:    "Welcome to the Brooklyn bridge (cap)",
			expected: "Welcome to the Brooklyn Bridge",
		},
		{
			name:     "Uppercase with count",
			input:    "This is so exciting (up, 2)",
			expected: "This is SO EXCITING",
		},
		{
			name:     "Lowercase with count",
			input:    "HELLO WORLD TEST (low, 3)",
			expected: "hello world test",
		},
		{
			name:     "Capitalize with count",
			input:    "hello world test (cap, 2)",
			expected: "hello World Test",
		},
		{
			name:     "Article conversion - up",
			input:    "a (up) elephant",
			expected: "AN elephant",
		},
		{
			name:     "Article conversion - cap",
			input:    "a (cap) elephant",
			expected: "An elephant",
		},
		{
			name:     "Article conversion - low",
			input:    "A (low) elephant",
			expected: "an elephant",
		},
		{
			name:     "Article conversion with count",
			input:    "a elephant (up, 2)",
			expected: "A ELEPHANT",
		},
		{
			name:     "Mixed alphanumeric words",
			input:    "test123 word456 (up, 2)",
			expected: "TEST123 WORD456",
		},
		{
			name:     "Numbers only (should be consumed but unchanged)",
			input:    "hello 123 world (up, 3)",
			expected: "HELLO 123 WORLD",
		},
		{
			name:     "Punctuation should be ignored",
			input:    "hello ! world (up, 2)",
			expected: "HELLO ! WORLD",
		},
		{
			name:     "Zero count modifier",
			input:    "hello world (up, 0)",
			expected: "hello world",
		},
		{
			name:     "Count larger than available words",
			input:    "hello (up, 5)",
			expected: "HELLO",
		},
		{
			name:     "Chained modifiers",
			input:    "test (up) (low)",
			expected: "test",
		},
		{
			name:     "No modifiers",
			input:    "This is just normal text",
			expected: "This is just normal text",
		},
		{
			name:     "Invalid modifier format",
			input:    "hello ( up ) world",
			expected: "hello ( up ) world",
		},
		{
			name:     "Quotes with modifiers",
			input:    "' the END is NEAR ' (low, 3)",
			expected: "' the end is near '",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UpLowCap(tt.input)
			if result != tt.expected {
				t.Errorf("UpLowCap(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}