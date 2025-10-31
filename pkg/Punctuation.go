package pkg

import "unicode"

func FixPunctuation(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))

	isPunct := func(r rune) bool {
		switch r {
		case '.', ',', '!', '?', ':', ';':
			return true
		}
		return false
	}

	isQuote := func(r rune) bool {
		return r == '\'' || r == '’'
	}

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// Handle quoted spans: ensure opening and closing quotes are adjacent
		// to the inner content (no spaces immediately inside the quotes).
		if isQuote(r) {
			// find matching closing quote
			j := i + 1
			for j < len(runes) && !isQuote(runes[j]) {
				j++
			}
			// if no matching quote found, treat as normal rune
			if j >= len(runes) {
				newRunes = append(newRunes, r)
				continue
			}

			// inner content boundaries without leading/trailing spaces
			innerStart := i + 1
			for innerStart < j && unicode.IsSpace(runes[innerStart]) {
				innerStart++
			}
			innerEnd := j - 1
			for innerEnd >= innerStart && unicode.IsSpace(runes[innerEnd]) {
				innerEnd--
			}

			// Append opening quote
			newRunes = append(newRunes, runes[i])

			// Append inner content (normalize punctuation inside the quoted span)
			if innerStart <= innerEnd {
				innerProcessed := FixPunctuation(string(runes[innerStart : innerEnd+1]))
				newRunes = append(newRunes, []rune(innerProcessed)...)
			}

			// Append closing quote
			newRunes = append(newRunes, runes[j])

			// advance i to the closing quote (loop will ++ it)
			i = j
			continue
		}

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
				if !isPunct(next) && next != '\'' && next != '’' {
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
