package piscine

func IsUpper(s string) bool {
	uper := true
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
		} else {
			uper = false
		}
	}
	return uper
}
