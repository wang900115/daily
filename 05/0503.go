package kata

import (
	"fmt"
	"sort"
)

func SumOfDivided(lst []int) string {
	primeSums := make(map[int]int)
	for _, num := range lst {
		n := abs(num)
		for i := 2; i <= n; i++ {
			if n%i == 0 && isPrime(i) {
				primeSums[i] += num
			}
		}
	}

	keys := make([]int, 0, len(primeSums))
	for k := range primeSums {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	result := ""
	for _, k := range keys {
		result += fmt.Sprintf("(%d %d)", k, primeSums[k])
	}
	return result
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
