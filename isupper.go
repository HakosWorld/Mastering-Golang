package piscine

func IsUpper(s string) bool {
	str := []rune(s)
	c := 0
	for i := 0; i < len(s); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			c++
		}
	}
	if c == len(s) {
		return true
	} else {
		return false
	}
}
