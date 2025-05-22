package kata

import (
	"math/big"
	"sort"
)

func ListPosition(word string) *big.Int {
	counts := make(map[rune]int)

	// counts 為 freq of string
	for _, ch := range word {
		counts[ch]++
	}

	// letters 用來排序後(依照rune大小) 的 freq of string
	letters := make([]rune, 0, len(counts))
	for ch := range counts {
		letters = append(letters, ch)
	}

	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	rank := big.NewInt(1)
	// 針對字串每個字元，計算所有"比它小"的排列數，並累加
	for _, ch := range word {
		for _, c := range letters {
			if c >= ch {
				break
			}

			if counts[c] == 0 {
				continue
			}

			counts[c]--
			rank.Add(rank, permutations(counts))
			counts[c]++
		}

		counts[ch]--
	}
	return rank
}

// 計算該排列組合
func permutations(counts map[rune]int) *big.Int {
	total := 0
	for _, v := range counts {
		total += v
	}

	result := factorial(total)
	for _, v := range counts {
		if v > 1 {
			result.Div(result, factorial(v))
		}
	}
	return result
}

// 階層 n
func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
