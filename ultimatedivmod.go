package piscine

func UltimateDivMod(a *int, b *int) {
	aa := *a
	bb := *b
	*a = aa / bb
	*b = aa % bb
}
