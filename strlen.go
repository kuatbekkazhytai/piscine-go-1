package strlen

import "github.com/01-edu/z01"

func StrLen(str string) int {
	runes := []rune(str)
	var count int = 0
	for i := range runes {
		count++
	}
	return count
}
