package pkg

import (
	"testing"
)

func TestAnConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple article conversion",
			input:    "There it was. A amazing rock!",
			expected: "There it was. An amazing rock!",
		},
		{
			name:     "Lowercase article conversion",
			input:    "I saw a elephant yesterday",
			expected: "I saw an elephant yesterday",
		},
		{
			name:     "Article before h",
			input:    "A hour passed by",
			expected: "An hour passed by",
		},
		{
			name:     "Article before consonant (no change)",
			input:    "A book is on the table",
			expected: "A book is on the table",
		},
		{
			name:     "Multiple article conversions",
			input:    "I saw a elephant and a orange",
			expected: "I saw an elephant and an orange",
		},
		{
			name:     "Article after punctuation",
			input:    "Yes, a elephant appeared",
			expected: "Yes, an elephant appeared",
		},
		{
			name:     "Article with quoted word",
			input:    "a 'elephant' appeared",
			expected: "an 'elephant' appeared",
		},
		{
			name:     "Article at start of text",
			input:    "a elephant walked by",
			expected: "an elephant walked by",
		},
		{
			name:     "Preserve uppercase from modifiers",
			input:    "A ELEPHANT",
			expected: "AN ELEPHANT",
		},
		{
			name:     "Mixed case scenarios",
			input:    "a elephant, A orange, a apple",
			expected: "an elephant, An orange, an apple",
		},
		{
			name:     "No conversion needed",
			input:    "The cat and the dog",
			expected: "The cat and the dog",
		},
		{
			name:     "Article not standalone",
			input:    "cat elephant",
			expected: "cat elephant",
		},
		{
			name:     "Multiple spaces",
			input:    "a    elephant",
			expected: "an    elephant",
		},
		{
			name:     "Article with all vowels",
			input:    "a apple, a elephant, a igloo, a orange, a umbrella",
			expected: "an apple, an elephant, an igloo, an orange, an umbrella",
		},
		{
			name:     "Complex sentence",
			input:    "I hear something like: a a a!",
			expected: "I hear something like: an an a!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AnConvert(tt.input)
			if result != tt.expected {
				t.Errorf("AnConvert(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}