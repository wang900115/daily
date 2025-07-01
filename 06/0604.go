package kata

func FindOutlier(integers []int) int {
	var even, odd []int

	for _, num := range integers[:3] {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}

	var isEvenMaj bool
	if len(even) > len(odd) {
		isEvenMaj = true
	} else {
		isEvenMaj = false
	}

	for _, n := range integers {
		if isEvenMaj && n%2 != 0 {
			return n
		}

		if !isEvenMaj && n%2 == 0 {
			return n
		}
	}
	return 0
}
