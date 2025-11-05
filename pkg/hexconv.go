package pkg

import (
	"math/big"
	"unicode"
)

// HexBin converts hexadecimal and binary numbers to decimal
// Processes (hex) and (bin) modifiers that follow numbers
func HexBin(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {

		// Look for (hex) modifier
		if i+5 < len(runes) && runes[i] == '(' && runes[i+1] == 'h' && runes[i+2] == 'e' && runes[i+3] == 'x' && runes[i+4] == ')' {
			removeTrailingSpaces(&newRunes)
			applyHex(&newRunes, 1) // Convert previous hex number to decimal
			i += 4                 // Skip "(hex)" pattern
			continue
		}

		// Look for (bin) modifier
		if i+5 < len(runes) && runes[i] == '(' && runes[i+1] == 'b' && runes[i+2] == 'i' && runes[i+3] == 'n' && runes[i+4] == ')' {
			removeTrailingSpaces(&newRunes)
			applyBin(&newRunes, 1) // Convert previous binary number to decimal
			i += 4                 // Skip "(bin)" pattern
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

// applyHex converts the previous hex number(s) to decimal
func applyHex(newRunes *[]rune, count int) {
	index := len(*newRunes) - 1
	applied := 0

	// Define valid hexadecimal characters
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
		// Find word boundaries
		for index >= 0 && !unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		start := index + 1

		// Extract only hex digits, ignoring punctuation
		tStart, tEnd := trimTokenBounds(*newRunes, start, end, isHex)
		if tStart > tEnd {
			applied++ // No valid hex found, skip this word
			continue
		}

		// Convert hex string to decimal using big.Int for large numbers
		hexStr := string((*newRunes)[tStart : tEnd+1])
		bigInt := new(big.Int)
		bigInt, success := bigInt.SetString(hexStr, 16)
		if success {
			decimalStr := []rune(bigInt.String())

			// Preserve spacing with following letters
			right := (*newRunes)[tEnd+1:]
			if len(right) > 0 && unicode.IsLetter(right[0]) {
				decimalStr = append(decimalStr, ' ')
			}

			// Replace hex number with decimal equivalent
			*newRunes = append((*newRunes)[:tStart], append(decimalStr, right...)...)
		}

		applied++
	}
}

// applyBin converts the previous binary number(s) to decimal
func applyBin(newRunes *[]rune, count int) {
	index := len(*newRunes) - 1
	applied := 0

	// Define valid binary characters
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
		// Find word boundaries
		for index >= 0 && !unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		start := index + 1

		// Extract only binary digits, ignoring punctuation
		tStart, tEnd := trimTokenBounds(*newRunes, start, end, isBin)
		if tStart > tEnd {
			applied++ // No valid binary found, skip this word
			continue
		}

		// Convert binary string to decimal using big.Int for large numbers
		binStr := string((*newRunes)[tStart : tEnd+1])
		bigInt := new(big.Int)
		bigInt, success := bigInt.SetString(binStr, 2)
		if success {
			decimalStr := []rune(bigInt.String())

			// Preserve spacing with following letters
			right := (*newRunes)[tEnd+1:]
			if len(right) > 0 && unicode.IsLetter(right[0]) {
				decimalStr = append(decimalStr, ' ')
			}

			// Replace binary number with decimal equivalent
			*newRunes = append((*newRunes)[:tStart], append(decimalStr, right...)...)
		}

		applied++
	}
}
