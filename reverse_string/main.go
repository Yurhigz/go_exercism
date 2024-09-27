package reverse

func Reverse(input string) string {
	runes := []rune(input)
	reversed := ""
	for i := len(runes) - 1; i >= 0; i-- {
		reversed += string(runes[i])
	}
	return reversed
}
