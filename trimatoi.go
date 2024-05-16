package piscine

func TrimAtoi(s string) int {
	onum := 0
	count := 0
	ars := []rune(s)
	pl := 1
	minuspos := 0
	minushave := false
	firstnum := 0
	for index, num := range ars {
		if num == '-' {
			minuspos = index
			minushave = true
		}
	}
	for _, word := range ars {
		if word >= '0' && word <= '9' {
			for i := '0'; i < word; i++ {
				count++
			}
			onum = onum*10 + count
			count = 0
		}
	}
	if minushave {
		for index, num := range ars {
			if num >= '0' && num <= '9' {
				firstnum = index
				break
			}
		}
	}
	if firstnum > minuspos {
		pl = -1
	}
	return onum * pl
}
