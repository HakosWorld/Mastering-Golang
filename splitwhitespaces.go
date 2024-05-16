package piscine

func SplitWhiteSpaces(str string) []string {
	count := 0
	ln := 0

	ok := false
	for c := range str {

		if ok && c != 0 && (str[c-1] == '\n' || str[c-1] == '\t' || str[c-1] == ' ') && str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			ln++
		}
		if str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			ok = true
		}
	}
	ln++

	rang := make([]string, ln)
	index := 0
	myStr := ""
	for i, w := range str {
		if w == '\n' || w == ' ' || w == '\t' {
			if myStr != "" {
				rang[index] = myStr
				index++
				myStr = ""
				count = i
			}
		} else {
			if w != ' ' {
				myStr = myStr + string(w)
			}
		}
	}

	if str[count+1:] != "" && str[count+1:] != " " {
		rang[index] = str[count+1:]
	}
	return rang
}
