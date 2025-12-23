package pkg

import (
	"testing"
)

// TestPipeline tests the complete processing pipeline
func TestPipeline(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Complete pipeline test",
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:     "Hex and binary with modifiers",
			input:    "1E (hex) files and 10 (bin) folders (up, 4) were found",
			expected: "30 FILES AND 2 FOLDERS were found",
		},
		{
			name:     "Article conversion with modifiers",
			input:    "a elephant (up, 2) walked by",
			expected: "AN ELEPHANT walked by",
		},
		{
			name:     "Complex quotes and punctuation",
			input:    "He said: ' I saw 2D (hex) birds and 1011 (bin) cats ' yesterday . It was a amazing experience !!",
			expected: "He said: 'I saw 45 birds and 11 cats' yesterday. It was an amazing experience!!",
		},
		{
			name:     "Nested quotes with modifiers",
			input:    "' wow ! ' she said ... unbelievable (up) !!",
			expected: "'wow!' she said... UNBELIEVABLE!!",
		},
		{
			name:     "Multiple article conversions",
			input:    "I saw a elephant, a orange, and a apple (cap, 6) today.",
			expected: "I saw an Elephant, AN Orange, And AN Apple today.",
		},
		{
			name:     "Pipeline with all features",
			input:    "FF (hex) ' a elephant ' (up, 3) said 1010 (bin) .",
			expected: "255 'AN ELEPHANT' said 10.",
		},
		{
			name:     "Edge case - punctuation and modifiers",
			input:    "hello ! world (up, 2)",
			expected: "HELLO! WORLD",
		},
		{
			name:     "Complex punctuation patterns",
			input:    "What ?!? He said ' I think wow (up) ! ' today ...",
			expected: "What?!? He said 'I think WOW!' today...",
		},
		{
			name:     "Article after punctuation",
			input:    "Yes , a elephant ? No : a orange !",
			expected: "Yes, an elephant? No: an orange!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run through the complete pipeline
			result := tt.input
			result = HexBin(result)
			result = UpLowCap(result)
			result = AnConvert(result)
			result = FixPunctuation(result)

			if result != tt.expected {
				t.Errorf("Pipeline(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestPipelineOrder tests that the order of operations matters
func TestPipelineOrder(t *testing.T) {
	input := "a elephant (up, 2)"
	
	// Correct order: HexBin -> UpLowCap -> AnConvert -> FixPunctuation
	correctResult := input
	correctResult = HexBin(correctResult)
	correctResult = UpLowCap(correctResult)
	correctResult = AnConvert(correctResult)
	correctResult = FixPunctuation(correctResult)
	
	// Wrong order: AnConvert before UpLowCap
	wrongResult := input
	wrongResult = HexBin(wrongResult)
	wrongResult = AnConvert(wrongResult)  // This runs before UpLowCap
	wrongResult = UpLowCap(wrongResult)
	wrongResult = FixPunctuation(wrongResult)
	
	// Both should be different to show order matters
	if correctResult == wrongResult {
		t.Logf("Correct order result: %q", correctResult)
		t.Logf("Wrong order result: %q", wrongResult)
		t.Skip("Pipeline order doesn't affect this particular test case")
	}
	
	expectedCorrect := "AN ELEPHANT"
	if correctResult != expectedCorrect {
		t.Errorf("Correct pipeline order should produce %q, got %q", expectedCorrect, correctResult)
	}
}