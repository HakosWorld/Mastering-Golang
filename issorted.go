package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	bol := true
	bol2 := true
	for c := 0; c <= len(a)-2; c++ {
		a1 := a[c]
		a2 := a[c+1]
		if f(a1, a2) > 0 {
			bol = false
		}
	}
	for c := 0; c <= len(a)-2; c++ {
		a1 := a[c]
		a2 := a[c+1]
		if f(a1, a2) < 0 {
			bol2 = false
		}
	}
	if bol || bol2 {
		return (true)
	} else {
		return false
	}
}
