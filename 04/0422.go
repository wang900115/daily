package kata

func ArrayDiff(a, b []int) []int {
	set := make(map[int]bool)
	for _, v := range b {
		set[v] = true
	}
	var res []int
	for _, v := range a {
		if !set[v] {
			res = append(res, v)
		}
	}
	return res
}
