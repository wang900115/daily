package kata

import (
	"strings"
)

func DuplicateEncode(word string) string {

	res := ""

	mapping := make(map[string]int)

	for _, char := range word {
		mapping[strings.ToLower(string(char))]++
	}

	for _, char := range word {
		if mapping[strings.ToLower(string(char))] > 1 {
			res += ")"
		} else {
			res += "("
		}
	}

	return res
}
