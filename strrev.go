package piscine

func StrRev(s string) string {
	chr := []rune(s)
	for i, j := 0, len(chr)-1; i < j; i, j = i+1, j-1 {
		chr[i], chr[j] = chr[j], chr[i]
	}
	return string(chr)
}
