package piscine

import (
	"strconv"
	"unicode"
)

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

func applyHex(newRunes *[]rune, count int) {
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
		for index >= 0 && !unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		start := index + 1

		hexStr := string((*newRunes)[start : end+1])
		value, err := strconv.ParseInt(hexStr, 16, 64)
		if err == nil {
			decimalStr := []rune(strconv.FormatInt(value, 10))

			// Check if we need to add a space after
			after := (*newRunes)[end+1:]
			if len(after) > 0 && unicode.IsLetter(after[0]) {
				decimalStr = append(decimalStr, ' ')
			}

			*newRunes = append((*newRunes)[:start], append(decimalStr, after...)...)
		}

		applied++
	}
}

func applyBin(newRunes *[]rune, count int) {
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
		for index >= 0 && !unicode.IsSpace((*newRunes)[index]) {
			index--
		}
		start := index + 1

		binStr := string((*newRunes)[start : end+1])
		value, err := strconv.ParseInt(binStr, 2, 64)
		if err == nil {
			decimalStr := []rune(strconv.FormatInt(value, 10))

			// Add space if next rune is a letter
			after := (*newRunes)[end+1:]
			if len(after) > 0 && unicode.IsLetter(after[0]) {
				decimalStr = append(decimalStr, ' ')
			}

			*newRunes = append((*newRunes)[:start], append(decimalStr, after...)...)
		}

		applied++
	}
}
