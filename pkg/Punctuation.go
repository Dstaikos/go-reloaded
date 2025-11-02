package pkg

import "unicode"

func FixPunctuation(s string) string {
	// First pass: handle quotes
	s = fixQuotes(s)
	
	// Second pass: handle punctuation
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

		// If we hit a punctuation rune, collect the whole punctuation sequence
		if isPunct(r) {
			j := i
			for j < len(runes) && isPunct(runes[j]) {
				j++
			}

			// Remove any spaces before punctuation and append the punctuation unit
			removeTrailingSpaces(&newRunes)
			newRunes = append(newRunes, runes[i:j]...)

			// Skip any spaces after the punctuation in the input
			k := j
			for k < len(runes) && unicode.IsSpace(runes[k]) {
				k++
			}

			// If there is something after the punctuation, ensure exactly one space
			// unless the next rune is another punctuation (but allow space before quotes)
			if k < len(runes) {
				next := runes[k]
				if !isPunct(next) {
					newRunes = append(newRunes, ' ')
				}
			}

			// Advance i to the last consumed position (loop will ++ it)
			i = k - 1
			continue
		}

		// Default: copy rune
		newRunes = append(newRunes, r)
	}

	return string(newRunes)
}

func fixQuotes(s string) string {
	runes := []rune(s)

	isQuote := func(r rune) bool {
		return r == '\'' || r == '\u2019'
	}

	result := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {
		if isQuote(runes[i]) {
			// Find matching closing quote
			j := i + 1
			for j < len(runes) && !isQuote(runes[j]) {
				j++
			}

			if j >= len(runes) {
				// No matching quote, copy as is
				result = append(result, runes[i])
				continue
			}

			// Remove trailing spaces before opening quote
			removeTrailingSpaces(&result)

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

			i = j // Skip to after closing quote
			continue
		}

		// Default: copy rune
		result = append(result, runes[i])
	}

	return string(result)
}