package piscine

func FindNextPrime(nb int) int {
	n := nb + 1
	for !(IsPrime(nb)) {
		if IsPrime(n) {
			return n
		} else {
			n++
		}
	}
	return nb
}
