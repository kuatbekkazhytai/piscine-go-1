package strrev

func StrRev(s string) string {
	runes = []rune(s)
	var len int = 0
	for v := range runes {
		len++
		v = v
	}
	var resRunes [len]rune
	j := len - 1
	for i := 0; i < len; i++ {
		resRunes[i] = runes[j]
		j--
	}
	return string(resRunes)
}
