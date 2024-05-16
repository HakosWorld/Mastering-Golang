package piscine

func AlphaCount(s string) int {
	str := []rune(s)
	c := 0
	for i := 0; i < len(s); i++ {
		if (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= 'a' && str[i] <= 'z') {
			c++
		}
	}
	return c
}
