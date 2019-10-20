package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else {
		n1 := 0
		n2 := 1
		t := 0
		for i := 0; i < index; i++ {
			t = n2 + n1
			n1 = n2
			n2 = t
		}
		return n1
	}
}
