package repeat

// Repeat ...
func Repeat(character string) string {
	res := ""
	for i := 0; i < 5; i++ {
		res = res + character
	}
	return res
}
