package pkg

import "unicode"

// AnConvert changes "a" to "an" before vowel-starting words
// Runs multiple passes to handle cases like "a a apple" → "an an apple"
func AnConvert(s string) string {
	for {
		newS := anConvertOnce(s)
		if newS == s {
			break // No more changes, conversion complete
		}
		s = newS
	}
	return s
}

// anConvertOnce performs one pass of "a" to "an" conversion
func anConvertOnce(s string) string {
	// Define vowels and 'h' that require "an" instead of "a"
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', 'h', 'H'}

	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {
		// Look for standalone "a" or "A" followed by space
		// Handles both " a " (normal) and "a " (after punctuation/start)
		if (i+2 < len(runes) && runes[i] == ' ' && (runes[i+1] == 'a' || runes[i+1] == 'A') && runes[i+2] == ' ') ||
		   (i+1 < len(runes) && (runes[i] == 'a' || runes[i] == 'A') && runes[i+1] == ' ' && (i == 0 || !unicode.IsLetter(runes[i-1]))) {
			
			// Determine pattern type and locate the "a"
			isSpaceA := runes[i] == ' '
			aIndex := i
			if isSpaceA {
				aIndex = i + 1
			}
			
			// Find start of next word after "a "
			j := aIndex + 2
			for j < len(runes) && unicode.IsSpace(runes[j]) {
				j++
			}

			// Convert "a" → "an" if next word starts with vowel/h
			if j < len(runes) && isVowel(runes[j], vowels) {
				// Rebuild the converted sequence
				if isSpaceA {
					newRunes = append(newRunes, ' ')
				}
				if runes[aIndex] == 'A' {
					newRunes = append(newRunes, 'A', 'n')
				} else {
					newRunes = append(newRunes, 'a', 'n')
				}
				newRunes = append(newRunes, ' ')
				
				// Skip processed characters
				if isSpaceA {
					i += 2
				} else {
					i += 1
				}
				continue
			}
		}
		newRunes = append(newRunes, runes[i])
	}

	return string(newRunes)
}

func isVowel(r rune, vowels []rune) bool {
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}