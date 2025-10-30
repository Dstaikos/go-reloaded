package piscine

import (
	"strconv"
	"unicode"
)

// ...existing code...

func HexBin(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {

		// Detect "(hex)"
		if i+5 < len(runes) && runes[i] == '(' && runes[i+1] == 'h' && runes[i+2] == 'e' && runes[i+3] == 'x' && runes[i+4] == ')' {

			removeTrailingSpaces(&newRunes)
			applyHex(&newRunes, 1)
			// pattern length is 5, so advance by 4 here; loop's i++ will move past ')'
			i += 4
			continue
		}

		// Detect "(bin)"
		if i+5 < len(runes) && runes[i] == '(' && runes[i+1] == 'b' && runes[i+2] == 'i' && runes[i+3] == 'n' && runes[i+4] == ')' {
			removeTrailingSpaces(&newRunes)
			applyBin(&newRunes, 1)
			// same adjustment as above
			i += 4
			continue
		}

		newRunes = append(newRunes, runes[i])
	}
	return string(newRunes)
}

// trimTokenBounds moves start/end inward to the actual token characters allowed by `allowed`.
// This allows skipping surrounding quotes or punctuation so we parse the numeric token correctly.
func trimTokenBounds(runes []rune, start, end int, allowed func(rune) bool) (int, int) {
	for start <= end && !allowed(runes[start]) {
		start++
	}
	for end >= start && !allowed(runes[end]) {
		end--
	}
	return start, end
}

func applyHex(newRunes *[]rune, count int) {
	index := len(*newRunes) - 1
	applied := 0

	isHex := func(r rune) bool {
		if r >= '0' && r <= '9' {
			return true
		}
		if r >= 'a' && r <= 'f' {
			return true
		}
		if r >= 'A' && r <= 'F' {
			return true
		}
		return false
	}

	for index >= 0 && applied < count {
		for index >= 0 && unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		if index < 0 {
			break
		}

		end := index
		for index >= 0 && !unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		start := index + 1

		// Trim surrounding punctuation/quotes so token starts/ends at hex digits
		tStart, tEnd := trimTokenBounds(*newRunes, start, end, isHex)
		if tStart > tEnd {
			// no hex token found in this span
			applied++
			continue
		}

		hexStr := string((*newRunes)[tStart : tEnd+1])
		value, err := strconv.ParseInt(hexStr, 16, 64)
		if err == nil {
			decimalStr := []rune(strconv.FormatInt(value, 10))

			// Prepare right-side slice after the token (use original end to keep surrounding punctuation)
			right := (*newRunes)[tEnd+1:]
			// Check if we need to add a space after decimal when right begins with a letter
			if len(right) > 0 && unicode.IsLetter(right[0]) {
				decimalStr = append(decimalStr, ' ')
			}

			// Replace only the trimmed token region, preserving anything before tStart and after tEnd
			*newRunes = append((*newRunes)[:tStart], append(decimalStr, right...)...)
		}

		applied++
	}
}

func applyBin(newRunes *[]rune, count int) {
	index := len(*newRunes) - 1
	applied := 0

	isBin := func(r rune) bool {
		return r == '0' || r == '1'
	}

	for index >= 0 && applied < count {
		for index >= 0 && unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		if index < 0 {
			break
		}

		end := index
		for index >= 0 && !unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		start := index + 1

		// Trim surrounding punctuation/quotes so token starts/ends at binary digits
		tStart, tEnd := trimTokenBounds(*newRunes, start, end, isBin)
		if tStart > tEnd {
			// no bin token found in this span
			applied++
			continue
		}

		binStr := string((*newRunes)[tStart : tEnd+1])
		value, err := strconv.ParseInt(binStr, 2, 64)
		if err == nil {
			decimalStr := []rune(strconv.FormatInt(value, 10))

			// Prepare right-side slice after the token
			right := (*newRunes)[tEnd+1:]
			// Add space if next rune is a letter
			if len(right) > 0 && unicode.IsLetter(right[0]) {
				decimalStr = append(decimalStr, ' ')
			}

			// Replace only the trimmed token region
			*newRunes = append((*newRunes)[:tStart], append(decimalStr, right...)...)
		}

		applied++
	}
}
