package kata

func SquareSum(numbers []int) int {
	res := 0
	for _, number := range numbers {
		res += number * number
	}
	return res
}
