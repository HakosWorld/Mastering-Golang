package piscine

func Index(s string, toFind string) int {
	sLen := len(s)
	toFindLen := len(toFind)
	for i := 0; i <= sLen-toFindLen; i++ {
		found := true
		for j := 0; j < toFindLen; j++ {
			if s[i+j] != toFind[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}
