package kata

import "math"

func Doubles(maxk, maxn int) float64 {
	var result float64
	for i := 1.00; i <= float64(maxk); i++ {
		for j := 1.00; j <= float64(maxn); j++ {
			top := 1.00
			bottom := i * math.Pow(j+1, 2*i)
			result += top / bottom

		}
	}

	return result
}
