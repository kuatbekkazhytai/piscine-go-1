package basicatoi

func BasicAtoi(s string) int {
	runes := []rune(s)

	var res int = 0
	for r := range runes {
		res *= 10
		res += r - '0'
	}
	return res
}
