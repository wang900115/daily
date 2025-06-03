package kata

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func Gap(g, m, n int) []int {
	prev := 0
	for i := m; i <= n; i++ {
		if isPrime(i) {
			if prev != 0 && i-prev == g {
				return []int{prev, i}
			}
			prev = i
		}
	}
	return nil
}
