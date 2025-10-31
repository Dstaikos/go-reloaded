package pkg

import "unicode"

// removeTrailingSpace deletes any spaces at the end of the slice

func removeTrailingSpaces(newRunes *[]rune) {
	for len(*newRunes) > 0 && unicode.IsSpace((*newRunes)[len(*newRunes)-1]) {
		*newRunes = (*newRunes)[:len(*newRunes)-1]
	}
}
