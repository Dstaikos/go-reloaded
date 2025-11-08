package pkg

import "unicode"

// FixPunctuation normalizes spacing around punctuation marks and quotes
// Two-pass approach: quotes first, then punctuation spacing
func FixPunctuation(s string) string {
	// First pass: process quotes and remove internal spaces
	s = fixQuotes(s)

	// Second pass: fix spacing around punctuation marks
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))

	isPunct := func(r rune) bool {
		switch r {
		case '.', ',', '!', '?', ':', ';':
			return true
		}
		return false
	}

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// Process punctuation sequences
		if isPunct(r) {
			// Collect consecutive punctuation marks
			j := i
			for j < len(runes) && isPunct(runes[j]) {
				j++
			}

			// Remove spaces before punctuation, add punctuation
			removeTrailingSpaces(&newRunes)
			newRunes = append(newRunes, runes[i:j]...)

			// Skip spaces after punctuation in input
			k := j
			for k < len(runes) && unicode.IsSpace(runes[k]) {
				k++
			}

			// Add single space after punctuation (with exceptions)
			if k < len(runes) {
				next := runes[k]
				isQuote := next == '\'' || next == '\u2019'

				// Detect closing quotes (preceded by content AND followed by space/punct/end)
				isClosingQuote := false
				if isQuote {
					if k > 0 {
						charBeforeQuote := runes[k-1]
						// Check what comes after the quote too
						charAfterQuote := rune(0)
						if k+1 < len(runes) {
							charAfterQuote = runes[k+1]
						}
						
						// It's a closing quote if preceded by content AND followed by space/punctuation/end
						precededByContent := unicode.IsLetter(charBeforeQuote) || unicode.IsDigit(charBeforeQuote) || isPunct(charBeforeQuote)
						followedBySpaceOrPunct := k+1 >= len(runes) || unicode.IsSpace(charAfterQuote) || isPunct(charAfterQuote)
						
						if precededByContent && followedBySpaceOrPunct {
							isClosingQuote = true
						}
					}
				}

				// Add space unless next is punctuation or closing quote
				if !isPunct(next) && !(isQuote && isClosingQuote) {
					newRunes = append(newRunes, ' ')
				}
			}

			i = k - 1 // Skip processed characters
			continue
		}

		// Default: copy rune
		newRunes = append(newRunes, r)
	}

	return string(newRunes)
}

// fixQuotes processes quote pairs and removes internal spaces
// Handles nested quotes and preserves spacing after punctuation
func fixQuotes(s string) string {
	runes := []rune(s)

	isQuote := func(r rune) bool {
		return r == '\'' || r == '\u2019'
	}

	isPunct := func(r rune) bool {
		switch r {
		case '.', ',', '!', '?', ':', ';':
			return true
		}
		return false
	}

	result := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {
		if isQuote(runes[i]) {
			// Skip apostrophes in contractions (don't vs 'quote')
			if i > 0 && i < len(runes)-1 && unicode.IsLetter(runes[i-1]) && unicode.IsLetter(runes[i+1]) {
				result = append(result, runes[i])
				continue
			}
			
			// Find matching closing quote
			j := i + 1
			for j < len(runes) && !isQuote(runes[j]) {
				j++
			}

			if j >= len(runes) {
				result = append(result, runes[i]) // Unmatched quote
				continue
			}

			// Only remove trailing spaces if they're not after punctuation
			if len(result) >= 2 && unicode.IsSpace(result[len(result)-1]) && isPunct(result[len(result)-2]) {
				// Don't remove space after punctuation
			} else if len(result) >= 1 && unicode.IsSpace(result[len(result)-1]) {
				// Check if there's content before the space that's not punctuation
				if len(result) >= 2 && !isPunct(result[len(result)-2]) {
					// Keep one space before quote if not after punctuation
					removeTrailingSpaces(&result)
					result = append(result, ' ')
				} else {
					removeTrailingSpaces(&result)
				}
			} else {
				removeTrailingSpaces(&result)
			}

			// Add opening quote
			result = append(result, runes[i])

			// Find content between quotes (trim spaces)
			contentStart := i + 1
			for contentStart < j && unicode.IsSpace(runes[contentStart]) {
				contentStart++
			}
			contentEnd := j - 1
			for contentEnd >= contentStart && unicode.IsSpace(runes[contentEnd]) {
				contentEnd--
			}

			// Add content if any
			if contentStart <= contentEnd {
				result = append(result, runes[contentStart:contentEnd+1]...)
			}

			// Add closing quote
			result = append(result, runes[j])

			// Look ahead for consecutive quote pairs
			k := j + 1

			// Skip spaces and non-quote content to find next quote
			for k < len(runes) && !isQuote(runes[k]) {
				k++
			}

			// If we found another quote, check if content between is only spaces/short words
			if k < len(runes) && isQuote(runes[k]) {
				// Check if content between quotes is "attachable" (spaces + short word)
				trimmed := string(runes[j+1 : k])
				// Remove leading/trailing spaces
				for len(trimmed) > 0 && unicode.IsSpace([]rune(trimmed)[0]) {
					trimmed = trimmed[1:]
				}
				for len(trimmed) > 0 && unicode.IsSpace([]rune(trimmed)[len([]rune(trimmed))-1]) {
					trimmed = trimmed[:len([]rune(trimmed))-1]
				}

				// If only spaces or short content, attach it
				if len(trimmed) == 0 || len([]rune(trimmed)) <= 5 {
					// Add the content without extra spaces
					if len(trimmed) > 0 {
						result = append(result, []rune(trimmed)...)
					}
					i = k - 1 // Will be incremented by loop to k
				} else {
					i = j // Normal case - preserve spaces
				}
			} else {
				i = j // Normal case - no consecutive quote found
			}
			continue
		}

		// Default: copy rune
		result = append(result, runes[i])
	}

	return string(result)
}
