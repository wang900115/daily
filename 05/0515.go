package kata

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func countArray(arr []int) int {
	result := arr[0]

	for _, value := range arr[1:] {
		result = gcd(result, value)
	}

	return result
}

func Solution(ar []int) int {
	return len(ar) * countArray(ar)
}
