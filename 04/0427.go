package kata

func Persistence(n int) int {
	count := 0
	for n >= 10 {
		product := 1
		for n > 0 {
			product *= n % 10
			n /= 10
		}
		n = product
		count++
	}
	return count
}
