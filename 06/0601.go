package kata

func PositiveSum(numbers []int) int {
	res := 0
	for _, num := range numbers {
		if num > 0 {
			res += num
		}
	}
	return res
}
