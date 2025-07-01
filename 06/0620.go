package kata

func CountSheeps(numbers []bool) int {
	res := 0
	for _, number := range numbers {
		if number {
			res++
		}
	}

	return res
}
