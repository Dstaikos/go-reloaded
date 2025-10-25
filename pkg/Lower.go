package piscine

import (
	"unicode"
)

func Lower(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {

		// Detect "(low"
		if i+4 < len(runes) && runes[i] == '(' && runes[i+1] == 'l' && runes[i+2] == 'o' && runes[i+3] == 'w' {

			// ---------- CASE 1: (low) ----------
			if i+4 < len(runes) && runes[i+4] == ')' {
				removeTrailingSpace(&newRunes)
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
					removeTrailingSpace(&newRunes)
					applyLow(&newRunes, num)
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
func applyLow(newRunes *[]rune, count int) {
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
				(*newRunes)[i] = unicode.ToLower((*newRunes)[i])
			}
		}
	}
}

// removeTrailingSpace deletes any spaces at the end of the slice
//func removeTrailingSpace(newRunes *[]rune) {
//for len(*newRunes) > 0 && unicode.IsSpace((*newRunes)[len(*newRunes)-1]) {
//*newRunes = (*newRunes)[:len(*newRunes)-1]
//}
//}
