package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if IsUpper(string(runes[i])) {
			runes[i] = runes[i] + 32
		}
	}
	return string(runes)
}
