package kata

import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	fmt.Println(s)
	words := strings.Split(s, " ")
	minLen := len(words[0])

	for _, word := range words {
		if len(word) < minLen {
			minLen = len(word)
		}
	}
	return minLen
}
