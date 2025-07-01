package kata

import (
	"strings"
)

func duplicate_count(s1 string) int {
	counts := make(map[rune]int)
	for _, r := range strings.ToLower(s1) {
		counts[r]++
	}

	result := 0
	for _, c := range counts {
		if c > 1 {
			result++
		}
	}
	return result
}
