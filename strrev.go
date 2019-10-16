package strrev

func StrRev(s string) string {
	runes := []rune(s)
	var len int = 0
	var tempRune rune
	for v := range runes {
		len++
		v = v
	}
	var l int
	if len%2 == 0 {
		l = l / 2
	} else {
		l = l/2 - 1
	}
	for i := 0; i <= l; i++ {
		tempRune = runes[i]
		runes[i] = runes[l+i+1]
		runes[l+i+1] = tempRune
	}
	return string(runes)
}
