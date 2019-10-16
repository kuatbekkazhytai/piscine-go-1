package strrev

func StrRev(s string) string {
	runes := []rune(s)
	var len int = 0
	var tempRune rune
	for v := range runes {
		len++
		v = v
	}
	for i := 0; i <= len/2; i++ {
		tempRune = runes[i]
		runes[i] = runes[i*2]
		runes[i*2] = tempRune
	}
	return string(runes)
}
