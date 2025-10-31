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
			// unless the next rune is another punctuation or a closing single quote (attach to punctuation).
			if k < len(runes) {
				next := runes[k]
				if !isPunct(next) && next != '\'' && next != '\u2019' {
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

	// Find all quote positions
	quotes := []int{}
	for i, r := range runes {
		if isQuote(r) {
			quotes = append(quotes, i)
		}
	}

	// Process quotes in pairs
	result := make([]rune, 0, len(runes))
	lastPos := 0

	for i := 0; i < len(quotes); i += 2 {
		if i+1 >= len(quotes) {
			// Odd number of quotes, copy rest as is
			result = append(result, runes[lastPos:]...)
			break
		}

		start := quotes[i]
		end := quotes[i+1]

		// Copy everything before this quote pair, but remove trailing spaces
		beforeQuote := runes[lastPos:start]
		// Skip leading spaces in beforeQuote
		for len(beforeQuote) > 0 && unicode.IsSpace(beforeQuote[0]) {
			beforeQuote = beforeQuote[1:]
		}
		result = append(result, beforeQuote...)
		removeTrailingSpaces(&result)

		// Add opening quote
		result = append(result, runes[start])

		// Find content between quotes (trim spaces)
		contentStart := start + 1
		for contentStart < end && unicode.IsSpace(runes[contentStart]) {
			contentStart++
		}
		contentEnd := end - 1
		for contentEnd >= contentStart && unicode.IsSpace(runes[contentEnd]) {
			contentEnd--
		}

		// Add content if any
		if contentStart <= contentEnd {
			result = append(result, runes[contentStart:contentEnd+1]...)
		}

		// Add closing quote
		result = append(result, runes[end])

		lastPos = end + 1
	}

	// Copy any remaining content after last quote pair
	if lastPos < len(runes) {
		remaining := runes[lastPos:]
		// Skip leading spaces if they exist
		for len(remaining) > 0 && unicode.IsSpace(remaining[0]) {
			remaining = remaining[1:]
		}
		result = append(result, remaining...)
	}

	return string(result)
}
