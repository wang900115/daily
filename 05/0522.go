package kata

import "strconv"

func Digits(n uint64) int {
	str_int := strconv.FormatUint(n, 10)
	return len(str_int)
}
