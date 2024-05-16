package piscine

func IsLower(s string) bool {
	str := []rune(s)
	c := 0
	for i := 0; i < len(s); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			c++
		}
	}
	if c == len(s) {
		return true
	} else {
		return false
	}
}
