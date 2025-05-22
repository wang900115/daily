package kata

import "math/big"

type matrix2x2 [2][2]*big.Int

func multiply(a matrix2x2, b matrix2x2) matrix2x2 {
	var c matrix2x2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = big.NewInt(0)
			for k := 0; k < 2; k++ {
				temp := new(big.Int).Mul(a[i][k], b[k][j])
				c[i][j].Add(c[i][j], temp)
			}
		}
	}
	return c
}

func matrixPower(m matrix2x2, n int64) matrix2x2 {
	result := matrix2x2{
		{big.NewInt(1), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(1)},
	}

	for n > 0 {
		if n%2 == 1 {
			result = multiply(result, m)
		}
		m = multiply(m, m)
		n /= 2
	}

	return result
}

func Fib(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	sign := 1
	if n < 0 {
		if n%2 == 0 {
			sign = -1
		}
		n = -n
	}
	base := matrix2x2{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0)},
	}
	result := matrixPower(base, n-1)
	fibN := new(big.Int).Set(result[0][0])
	if sign == -1 {
		fibN.Neg(fibN)
	}
	return fibN
}
