package kata

import "strings"

func ToCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '_'
	})

	if len(words) == 0 {
		return ""
	}

	result := words[0]

	for _, word := range words[1:] {
		if len(word) > 0 {
			result += strings.ToUpper(string(word[0])) + word[1:]
		}
	}

	return result
}
