package piscine_go

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		n *= -1
		z01.PrintRune('-')
	} else if n == 0 {
		z01.PrintRune('0')
	}
	PrintNbr(n / 10)
	r := '0'
	for i := 0; i < n%10; i++ {
		r++
	}
	z01.PrintRune(r)
}
