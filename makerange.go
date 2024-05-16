package piscine

func MakeRange(min, max int) []int {
	len := max - min
	var rang []int

	if min >= max {
		return nil
	}

	rang = make([]int, len)
	for i := 0; i < len; i++ {
		rang[i] = min
		min++
	}
	return rang
}
