package piscine_go

func StrRev(s string) string {
	bytes := []byte(s)
	var len int = 0
	var tempByte byte
	for v := range bytes {
		len++
		v = v
	}
	for i := 0; i < len/2; i++ {
		tempByte = bytes[i]
		bytes[i] = bytes[len-i-1]
		bytes[len-i-1] = tempByte
	}
	return string(bytes)
}
