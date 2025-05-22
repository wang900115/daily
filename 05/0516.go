package kata

import "strconv"

func NextBigger(n int) int {
	digits := []byte(strconv.Itoa(n))
	// from right to left , find first
	i := len(digits) - 2
	for i >= 0 && digits[i] >= digits[i+1] {
		i--
	}

	if i < 0 {
		return -1
	}
	// from right to left, find bigger than digits[i] but smallest
	j := len(digits) - 1
	for digits[j] <= digits[i] {
		j--
	}

	// swap
	digits[i], digits[j] = digits[j], digits[i]

	// reversed
	for k, l := i+1, len(digits)-1; k < l; k, l = k+1, l-1 {
		digits[k], digits[l] = digits[l], digits[k]
	}

	res, _ := strconv.Atoi(string(digits))
	return res

}
