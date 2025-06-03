package kata

func Sum(n int) int {
	res := 0
	for n > 0 {
		r := n % 10
		n /= 10
		res += r
	}
	return res
}

func DigitalRoot(n int) int {
	for n > 9 {
		n = Sum(n)
	}

	return n
}
