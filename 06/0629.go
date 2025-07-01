package kata

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	res := 0

	for _, char := range strconv.Itoa(n) {
		digit := int(char - '0')
		res += int(math.Pow(float64(digit), float64(p)))
		p++
	}

	if res%n == 0 {
		return res / n
	}

	return -1
}
