package pkg

import (
	"unicode"
)

// UpLowCap applies case transformations based on (up), (low), and (cap) modifiers
// Also handles special article conversion when modifying "a" before vowel words
func UpLowCap(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {

		// Look for (up) modifiers
		if i+3 < len(runes) && runes[i] == '(' && runes[i+1] == 'u' && runes[i+2] == 'p' {

			// Handle (up) - uppercase previous word
			if i+3 < len(runes) && runes[i+3] == ')' {
				removeTrailingSpaces(&newRunes)
				// Check for article conversion first ("a" → "AN")
				if shouldConvertArticle(&newRunes, runes, i+4, "up") {
					i += 3
					continue
				}
				applyUp(&newRunes, 1)
				i += 3
				continue
			}

			// Handle (up, x) - uppercase previous x words
			if i+5 < len(runes) && runes[i+3] == ',' && runes[i+4] == ' ' {
				j := i + 5
				num := 0
				hasDigit := false

				// Parse the number
				for j < len(runes) && unicode.IsDigit(runes[j]) {
					num = num*10 + int(runes[j]-'0')
					j++
					hasDigit = true
				}

				if hasDigit && j < len(runes) && runes[j] == ')' {
					removeTrailingSpaces(&newRunes)
					// Check for article conversion first
					if shouldConvertArticle(&newRunes, runes, j+1, "up") {
						i = j
						continue
					}
					applyUp(&newRunes, num)
					i = j
					continue
				}
			}
		}

		// Detect "(low)"
		if i+4 < len(runes) && runes[i] == '(' && runes[i+1] == 'l' && runes[i+2] == 'o' && runes[i+3] == 'w' {

			// ---------- CASE 1: (low) ----------
			if i+4 < len(runes) && runes[i+4] == ')' {
				removeTrailingSpaces(&newRunes)
				if shouldConvertArticle(&newRunes, runes, i+5, "low") {
					i += 4
					continue
				}
				applyLow(&newRunes, 1)
				i += 4
				continue
			}

			// ---------- CASE 2: (low, x) ----------
			if i+6 < len(runes) && runes[i+4] == ',' && runes[i+5] == ' ' {
				j := i + 6
				num := 0
				hasDigit := false

				for j < len(runes) && unicode.IsDigit(runes[j]) {
					num = num*10 + int(runes[j]-'0')
					j++
					hasDigit = true
				}

				if hasDigit && j < len(runes) && runes[j] == ')' {
					removeTrailingSpaces(&newRunes)
					if shouldConvertArticle(&newRunes, runes, j+1, "low") {
						i = j
						continue
					}
					applyLow(&newRunes, num)
					i = j
					continue
				}
			}
		}

		// Detect "(cap)"
		if i+4 < len(runes) && runes[i] == '(' && runes[i+1] == 'c' && runes[i+2] == 'a' && runes[i+3] == 'p' {

			// ---------- CASE 1: (cap) ----------
			if i+4 < len(runes) && runes[i+4] == ')' {
				removeTrailingSpaces(&newRunes)
				if shouldConvertArticle(&newRunes, runes, i+5, "cap") {
					i += 4
					continue
				}
				applyCap(&newRunes, 1)
				i += 4
				continue
			}

			// ---------- CASE 2: (cap, x) ----------
			if i+6 < len(runes) && runes[i+4] == ',' && runes[i+5] == ' ' {
				j := i + 6
				num := 0
				hasDigit := false

				for j < len(runes) && unicode.IsDigit(runes[j]) {
					num = num*10 + int(runes[j]-'0')
					j++
					hasDigit = true
				}

				if hasDigit && j < len(runes) && runes[j] == ')' {
					removeTrailingSpaces(&newRunes)
					if shouldConvertArticle(&newRunes, runes, j+1, "cap") {
						i = j
						continue
					}
					applyCap(&newRunes, num)
					i = j
					continue
				}
			}
		}

		// Default case: copy rune
		newRunes = append(newRunes, runes[i])
	}

	return string(newRunes)
}

// ==================== HELPER FUNCTIONS ====================

// shouldConvertArticle handles special case: "a" + modifier + vowel word → "an"/"AN"/"An"
// Returns true if conversion was applied, false if normal modifier should be used
func shouldConvertArticle(newRunes *[]rune, originalRunes []rune, nextPos int, modifier string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', 'h', 'H'}
	
	// Find last word and check if it's standalone "a" or "A"
	index := len(*newRunes) - 1
	for index >= 0 && unicode.IsSpace((*newRunes)[index]) {
		index--
	}
	if index < 0 || ((*newRunes)[index] != 'a' && (*newRunes)[index] != 'A') {
		return false
	}
	
	// Ensure it's a standalone article, not part of another word
	if index > 0 && !unicode.IsSpace((*newRunes)[index-1]) {
		return false
	}
	
	// Find next word in original input
	for nextPos < len(originalRunes) && unicode.IsSpace(originalRunes[nextPos]) {
		nextPos++
	}
	
	if nextPos >= len(originalRunes) {
		return false
	}
	
	// Check if next word starts with vowel/h
	startsWithVowel := false
	for _, v := range vowels {
		if originalRunes[nextPos] == v {
			startsWithVowel = true
			break
		}
	}
	
	if !startsWithVowel {
		return false
	}
	
	// Apply appropriate article conversion based on modifier
	if modifier == "up" {
		(*newRunes)[index] = 'A'
		*newRunes = append((*newRunes)[:index+1], append([]rune{'N'}, (*newRunes)[index+1:]...)...)
	} else if modifier == "cap" {
		(*newRunes)[index] = 'A'
		*newRunes = append((*newRunes)[:index+1], append([]rune{'n'}, (*newRunes)[index+1:]...)...)
	} else if modifier == "low" {
		(*newRunes)[index] = 'a'
		*newRunes = append((*newRunes)[:index+1], append([]rune{'n'}, (*newRunes)[index+1:]...)...)
	}
	
	return true
}

// applyUp uppercases the previous `count` words, consuming all words regardless of content
func applyUp(newRunes *[]rune, count int) {
	index := len(*newRunes) - 1
	applied := 0

	for index >= 0 && applied < count {
		// Skip trailing spaces
		for index >= 0 && unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		if index < 0 {
			break
		}

		end := index
		hasLetter := false

		// Find word boundaries and check for letters
		for index >= 0 && !unicode.IsSpace((*newRunes)[index]) {
			if unicode.IsLetter((*newRunes)[index]) {
				hasLetter = true
			}
			index--
		}

		start := index + 1

		// Apply uppercase only to letters, but consume word regardless
		if hasLetter {
			for i := start; i <= end; i++ {
				(*newRunes)[i] = unicode.ToUpper((*newRunes)[i])
			}
		}
		// Always increment applied count (consumes word even if no letters)
		applied++
	}
}

// applyLow lowercases the previous `count` words, consuming all words regardless of content
func applyLow(newRunes *[]rune, count int) {
	index := len(*newRunes) - 1
	applied := 0

	for index >= 0 && applied < count {
		for index >= 0 && unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		if index < 0 {
			break
		}

		end := index
		hasLetter := false

		// Find word boundaries and check for letters
		for index >= 0 && !unicode.IsSpace((*newRunes)[index]) {
			if unicode.IsLetter((*newRunes)[index]) {
				hasLetter = true
			}
			index--
		}

		start := index + 1

		// Apply lowercase only to letters, but consume word regardless
		if hasLetter {
			for i := start; i <= end; i++ {
				(*newRunes)[i] = unicode.ToLower((*newRunes)[i])
			}
		}
		// Always increment applied count (consumes word even if no letters)
		applied++
	}
}

// applyCap capitalizes the previous `count` words, consuming all words regardless of content
func applyCap(newRunes *[]rune, count int) {
	index := len(*newRunes) - 1
	applied := 0

	for index >= 0 && applied < count {
		for index >= 0 && unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		if index < 0 {
			break
		}

		end := index
		hasLetter := false

		// Find word boundaries and check for letters
		for index >= 0 && !unicode.IsSpace((*newRunes)[index]) {
			if unicode.IsLetter((*newRunes)[index]) {
				hasLetter = true
			}
			index--
		}

		start := index + 1

		// Apply capitalization only to letters, but consume word regardless
		if hasLetter {
			firstLetterFound := false
			for i := start; i <= end; i++ {
				if unicode.IsLetter((*newRunes)[i]) {
					if !firstLetterFound {
						(*newRunes)[i] = unicode.ToUpper((*newRunes)[i])
						firstLetterFound = true
					} else {
						(*newRunes)[i] = unicode.ToLower((*newRunes)[i])
					}
				}
			}
		}
		// Always increment applied count (consumes word even if no letters)
		applied++
	}
}
