package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	chr := []rune(s)
	for i := 0; i < len(chr); i++ {
		z01.PrintRune(chr[i])
	}
}
