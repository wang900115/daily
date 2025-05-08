package kata

import "math/big"

func LastDigit(n1, n2 string) int {
	if n2 == "0" {
		return 1
	}

	lastDigit := int(n1[len(n1)-1] - '0')

	cycle := []int{}
	seen := map[int]bool{}
	current := lastDigit

	for !seen[current] {
		seen[current] = true
		cycle = append(cycle, current)
		current = (current * lastDigit) % 10
	}

	bigB := new(big.Int)
	bigB.SetString(n2, 10)

	cycleLen := len(cycle)
	mod := new(big.Int).Mod(bigB, big.NewInt(int64(cycleLen))).Int64()

	if mod == 0 {
		return cycle[cycleLen-1]
	}

	return cycle[mod-1]
}

func a(n1, n2 string) int {
	a, b := big.NewInt(0), big.NewInt(0)
	a.SetString(n1, 10)
	b.SetString(n2, 10)

	a.Exp(a, b, big.NewInt(10))

	return int(a.Int64())
}
