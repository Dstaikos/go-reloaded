package pkg

import (
	"testing"
)

func TestHexBin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple hex conversion",
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			name:     "Simple binary conversion",
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			name:     "Large hex number",
			input:    "FF (hex) is the maximum",
			expected: "255 is the maximum",
		},
		{
			name:     "Large binary number",
			input:    "1111 (bin) in decimal",
			expected: "15 in decimal",
		},
		{
			name:     "Multiple hex conversions",
			input:    "A (hex) and B (hex) make AB",
			expected: "10 and 11 make AB",
		},
		{
			name:     "Multiple binary conversions",
			input:    "101 (bin) and 110 (bin) are different",
			expected: "5 and 6 are different",
		},
		{
			name:     "No conversion needed",
			input:    "This is just normal text",
			expected: "This is just normal text",
		},
		{
			name:     "Invalid hex (should remain unchanged)",
			input:    "GGG (hex) is invalid",
			expected: "GGG is invalid",
		},
		{
			name:     "Invalid binary (should remain unchanged)",
			input:    "123 (bin) is invalid",
			expected: "123 is invalid",
		},
		{
			name:     "Overflow test - very large hex",
			input:    "FFFFFFFFFFFFFFFFFFF (hex)",
			expected: "FFFFFFFFFFFFFFFFFFF (hex)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HexBin(tt.input)
			if result != tt.expected {
				t.Errorf("HexBin(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}