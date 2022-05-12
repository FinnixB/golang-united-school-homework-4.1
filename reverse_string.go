package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)

	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}

	output = string(runes)

	return output
}
