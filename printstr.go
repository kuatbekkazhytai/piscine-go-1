package printstr

import "github.com/01-edu/z01"

func PrintStr(str string) {
	bytes := []byte(str)
	for i, v := range bytes {
		z01.PrintRune(bytes[i])
	}
}
