package kata

import "strings"

func GetCount(str string) (count int) {
	for _, char := range str {
		if strings.Contains("aeiou", string(char)) {
			count++
		}
	}
	return count
}
