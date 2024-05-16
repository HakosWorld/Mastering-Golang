package piscine

func AppendRange(min, max int) []int {
	var rage []int
	for i := min; i < max; i++ {
		rage = append(rage, i)
	}
	return rage
}
