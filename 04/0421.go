package kata

func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0 {
		return []float64{}
	}

	res := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		if i < 3 {
			res = append(res, signature[i])
		} else {
			value := res[i-1] + res[i-2] + res[i-3]
			res = append(res, value)
		}
	}

	return res
}

// golang 陣列與切片

// 陣列 [n]T

// a := [3]int{1, 2, 3}
// a[0] = 10

// 切片 []T
// a := []int{1, 2, 3}
// a = append(a,4)
