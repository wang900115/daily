package kata

import "strings"

func Accum(s string) string {
	var result []string

	for i, char := range s {
		repeated := strings.Repeat(strings.ToLower(string(char)), i)
		formatted := strings.ToUpper(string(char)) + repeated
		result = append(result, formatted)
	}

	return strings.Join(result, "-")
}
