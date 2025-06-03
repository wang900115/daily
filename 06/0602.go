package kata

func Solution(word string) string {
	res := ""

	for _, char := range word {
		res = string(char) + res
	}

	return res
}
