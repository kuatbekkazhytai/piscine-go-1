package piscine_go

func Swap(a *int, b *int) {
	*a += *b
	*b = *a - *b
	*a = *a - *b
}
