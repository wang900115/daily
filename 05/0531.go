package kata

import "strconv"

func CountBits(n uint) int {
	count := 0
	binary := strconv.FormatInt(int64(n), 2)
	for _, char := range binary {
		if char == '1' {
			count++
		}
	}
	return count
}
