package kata

func FindOdd(seq []int) int {

	mapping := make(map[int]int)

	for _, num := range seq {
		mapping[num]++
	}

	for key, value := range mapping {
		if value%2 == 1 {
			return key
		}
	}
	return 0
}
