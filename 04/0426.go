package kata

import "unicode"

func FirstNonRepeating(str string) string {
	count := make(map[rune]int)

	for _, ch := range str {
		lowerCh := unicode.ToLower(ch)
		count[lowerCh]++
	}

	for _, ch := range str {
		lowerCh := unicode.ToLower(ch)
		if count[lowerCh] == 1 {
			return string(ch)
		}
	}

	return ""
}
