package basicatoi

func BasicAtoi(s string) int {
	bytes := []byte(s)

	var res int = 0
	for b := range bytes {
		res *= 10
		res += b - 48
	}
	return res
}
