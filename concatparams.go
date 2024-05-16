package piscine

func ConcatParams(args []string) string {
	str := ""
	len := 0
	for i := range args {
		len = i
	}
	for i := 0; i <= len; i++ {
		if i != len {
			str = str + args[i] + "\n"
		} else {
			str = str + args[i]
		}
	}
	return str
}
