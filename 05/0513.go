package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(list []int) string {
	var res []string
	n := len(list)
	i := 0

	for i < n {
		start := i
		for i+1 < n && list[i]+1 == list[i+1] {
			i++
		}

		if i-start >= 2 {
			res = append(res, fmt.Sprintf("%d-%d", list[start], list[i]))
		} else {
			for j := start; j <= i; j++ {
				res = append(res, strconv.Itoa(list[j]))
			}
		}
		i++
	}

	return strings.Join(res, ",")

}
