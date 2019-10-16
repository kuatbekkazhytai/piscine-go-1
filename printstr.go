package printstr

import "github.com/01-edu/z01"

func PrintStr(str string) {
	bytes := []byte(str)
	for i := 0; i < len(str); i++ {
		z01.PrintRune(bytes[i])
	}
}
