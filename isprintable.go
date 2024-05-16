package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] < 32 {
			return false
		}
	}
	return true
}
