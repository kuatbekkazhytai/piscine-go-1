package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb == 2 {
		return true
	} else {
		for i := 3; i*i <= nb; i += 2 {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
}
