package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else {
		res := nb
		for i := 1; i < power; i++ {
			res *= nb
		}
		return res
	}
}
