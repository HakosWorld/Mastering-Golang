package piscine

func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}
	SortIntegerTable(numbers)
	return numbers[2]
}
