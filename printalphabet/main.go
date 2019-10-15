package main

import (
	"os"
	"unicode/utf8"
	"errors"
)

func PrintRune(r rune) error {
	l := utf8.RuneLen(r)
	if l == -1 {
		return errors.New("The rune is not a valid value to encode in UTF-8")
	}
	p := make([]byte, l)
	utf8.EncodeRune(p, r)
	_, err := os.Stdout.Write(p)
	return err
}

func main() {
	var aRune rune = 'z'
	var endRune rune = '\n'
	for iRune := 'a'; iRune <= aRune; iRune++ {
		PrintRune(iRune)
	}
	PrintRune(endRune)
}
