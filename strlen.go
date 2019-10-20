package piscine

func StrLen(str string) int {
	runes := []rune(str)
	var count int = 0
	for i := range runes {
		count++
		i = i
	}
	return count
}
