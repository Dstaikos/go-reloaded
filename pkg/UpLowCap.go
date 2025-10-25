package piscine

import (
	"unicode"
)

func UpLowCap(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {

		// Detect "(up"
		if i+3 < len(runes) && runes[i] == '(' && runes[i+1] == 'u' && runes[i+2] == 'p' {

			// ---------- CASE 1: (up) ----------
			if i+3 < len(runes) && runes[i+3] == ')' {
				removeTrailingSpaces(&newRunes)
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
					removeTrailingSpaces(&newRunes)
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

// applyLow uppercases the previous `count` words
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

// applyUp uppercases the previous `count` words
func applyCap(newRunes *[]rune, count int) {
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
}
