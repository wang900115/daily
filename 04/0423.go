package kata

func SameArray(array1, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	if len(array1) != len(array2) {
		return false
	}

	squareCount := make(map[int]int)
	for _, v := range array1 {
		val := v * v
		squareCount[val]++
	}

	for _, v := range array2 {
		if squareCount[v] == 0 {
			return false
		}
		squareCount[v]--
	}

	return true
}
