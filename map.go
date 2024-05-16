package piscine

func Map(f func(int) bool, a []int) []bool {
	var answer []bool
	for _, c := range a {
		answer = append(answer, f(c))
	}
	return answer
}
