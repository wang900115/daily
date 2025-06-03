package kata

func Disemvowel(comment string) string {
	res := ""
	for _, char := range comment {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			continue
		}
		if char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			continue
		}
		res += string(char)
	}
	return res
}
