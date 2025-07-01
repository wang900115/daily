package kata

import "sort"

func TwoToOne(s1 string, s2 string) string {
	charMap := make(map[rune]bool)
	for _, ch := range s1 + s2 {
		charMap[ch] = true
	}

	uniqueChars := make([]rune, 0, len(charMap))

	for ch := range charMap {
		uniqueChars = append(uniqueChars, ch)
	}

	sort.Slice(uniqueChars, func(i, j int) bool {
		return uniqueChars[i] < uniqueChars[j]
	})
	return string(uniqueChars)
}
