package kata

import "sort"

func Permutations(s string) []string {
	res := []string{}
	chars := []rune(s)
	used := make([]bool, len(chars))
	path := []rune{}
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

	var Backtrack func(depth int)
	Backtrack = func(depth int) {
		if depth == len(chars) {
			res = append(res, string(path))
			return
		}

		for i := 0; i < len(chars); i++ {
			if used[i] {
				continue
			}

			if i > 0 && chars[i] == chars[i-1] && !used[i-1] {
				continue
			}

			used[i] = true
			path = append(path, chars[i])
			Backtrack(depth + 1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	Backtrack(0)
	return res
}
