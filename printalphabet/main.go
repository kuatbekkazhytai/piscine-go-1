package main

import "github.com/01-edu/z01"

func main() {
	var aRune rune = 'z'
	var endRune rune = '\n'
	for iRune := 'a'; iRune <= aRune; iRune++ {
		z01.PrintRune(iRune)
	}
	z01.PrintRune(endRune)
}
