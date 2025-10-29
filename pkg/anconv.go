package piscine

import "unicode"

func AnConvert(s string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', 'h', 'H'}

	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {
		// Check for " a " or " A " pattern safely
		if i+2 < len(runes) && runes[i] == ' ' && (runes[i+1] == 'a' || runes[i+1] == 'A') && runes[i+2] == ' ' {
			// Find the next non-space character after " a "
			j := i + 3
			for j < len(runes) && unicode.IsSpace(runes[j]) {
				j++
			}

			// If next word starts with a vowel or h/H → change "a" → "an"
			if j < len(runes) && isVowel(runes[j], vowels) {
				newRunes = append(newRunes, ' ')
				if runes[i+1] == 'A' {
					newRunes = append(newRunes, 'A', 'n')
				} else {
					newRunes = append(newRunes, 'a', 'n')
				}
				newRunes = append(newRunes, ' ')
				i += 2 // Skip the processed " a "
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
