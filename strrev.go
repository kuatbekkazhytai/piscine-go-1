package strrev

func StrRev(s string) string {
	bytes := []byte(s)
	var len int = 0
	var tempByte byte
	for v := range bytes {
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
		tempByte = bytes[i]
		bytes[i] = bytes[l+i+1]
		bytes[l+i+1] = tempByte
	}
	return string(bytes)
}
