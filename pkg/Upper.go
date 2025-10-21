package piscine

import (
	"unicode"
)

func Upper(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {

		// Detect "(up"
		if i+3 < len(runes) && runes[i] == '(' && runes[i+1] == 'u' && runes[i+2] == 'p' {

			// ---------- CASE 1: (up) ----------
			if i+3 < len(runes) && runes[i+3] == ')' {
				removeTrailingSpace(&newRunes)
				applyUp(&newRunes, 1)
				i += 3
				continue
			}

			// ---------- CASE 2: (up, x) ----------
			if i+5 < len(runes) && runes[i+3] == ',' && runes[i+4] == ' ' {
				j := i + 5
				num := 0
				hasDigit := false

				for j < len(runes) && unicode.IsDigit(runes[j]) {
					num = num*10 + int(runes[j]-'0')
					j++
					hasDigit = true
				}

				if hasDigit && j < len(runes) && runes[j] == ')' {
					removeTrailingSpace(&newRunes)
					applyUp(&newRunes, num)
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

// applyUp uppercases the previous `count` words
func applyUp(newRunes *[]rune, count int) {
	index := len(*newRunes) - 1

	for w := 0; w < count; w++ {
		// Skip trailing spaces
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

		for i := start; i <= end; i++ {
			if unicode.IsLetter((*newRunes)[i]) {
				(*newRunes)[i] = unicode.ToUpper((*newRunes)[i])
			}
		}
	}
}

// removeTrailingSpace deletes any spaces at the end of the slice
func removeTrailingSpace(newRunes *[]rune) {
	for len(*newRunes) > 0 && unicode.IsSpace((*newRunes)[len(*newRunes)-1]) {
		*newRunes = (*newRunes)[:len(*newRunes)-1]
	}
}
