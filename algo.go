package main

import "fmt"

func oddAlgo(A []int) {
	j := 0
	n := len(A)
	for i := n - 1; i >= j; i-- {
		if A[i]%2 == 1 {
			// Swap elements if A[i] is odd
			A[i], A[j] = A[j], A[i]
			j++
			if A[j]%2 == 1 {
				// If the element at A[j] is also odd, increment i to maintain the order
				i++
			}
		}
	}
}

func main() {
	// Example usage
	arr := []int{1, 2, 3, 4, 5}
	oddAlgo(arr)
	fmt.Println("Array after oddAlgo:", arr)
}
