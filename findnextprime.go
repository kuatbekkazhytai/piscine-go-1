package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		if nb%2 == 0 {
			nb++
		}
		res := nb
		for i := nb; !IsPrime(i); i += 2 {
			res = i
		}
		return res
	}
}
