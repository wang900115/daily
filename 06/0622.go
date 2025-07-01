package kata

func GetSum(a, b int) int {
	if a == b {
		return a
	}

	res := 0

	if a < b {
		for i := a; i <= b; i++ {
			res += i
		}
		return res
	}

	for i := b; i <= a; i++ {
		res += i
	}
	return res

}
