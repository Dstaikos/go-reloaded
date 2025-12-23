package pkg

import (
	"testing"
)

func TestFixPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic punctuation spacing",
			input:    "I was sitting over there ,and then BAMM !!",
			expected: "I was sitting over there, and then BAMM!!",
		},
		{
			name:     "Multiple punctuation marks",
			input:    "I was thinking ... You were right",
			expected: "I was thinking... You were right",
		},
		{
			name:     "Simple quote processing",
			input:    "I am exactly how they describe me: ' awesome '",
			expected: "I am exactly how they describe me: 'awesome'",
		},
		{
			name:     "Multiple words in quotes",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			name:     "Nested quotes",
			input:    "' wow ! ' she said",
			expected: "'wow!' she said",
		},
		{
			name:     "Quote pairs",
			input:    "'   '  hello  '   '",
			expected: "''hello''",
		},
		{
			name:     "Punctuation after quotes",
			input:    "hey!'idiot'",
			expected: "hey! 'idiot'",
		},
		{
			name:     "Complex punctuation",
			input:    "What?!? Really...?",
			expected: "What?!? Really...?",
		},
		{
			name:     "Contractions (should not be processed as quotes)",
			input:    "don't you think!?",
			expected: "don't you think!?",
		},
		{
			name:     "Mixed quotes and punctuation",
			input:    "He said: 'What?!? Really...?' and left.",
			expected: "He said: 'What?!? Really...?' and left.",
		},
		{
			name:     "Empty quotes",
			input:    "He said '' and left",
			expected: "He said '' and left",
		},
		{
			name:     "Quote with punctuation inside",
			input:    "'Hello, world!' she said.",
			expected: "'Hello, world!' she said.",
		},
		{
			name:     "Multiple quote pairs",
			input:    "'hello' and 'world' are words",
			expected: "'hello'and'world' are words",
		},
		{
			name:     "Unmatched quote",
			input:    "He said 'hello world but forgot to close",
			expected: "He said 'hello world but forgot to close",
		},
		{
			name:     "No processing needed",
			input:    "This is normal text.",
			expected: "This is normal text.",
		},
		{
			name:     "Spaces around colons",
			input:    "He said : hello",
			expected: "He said: hello",
		},
		{
			name:     "Spaces around semicolons",
			input:    "First ; second ; third",
			expected: "First; second; third",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FixPunctuation(tt.input)
			if result != tt.expected {
				t.Errorf("FixPunctuation(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}