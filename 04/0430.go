package kata

import (
	"sort"
	"strings"
)

func Weight(strng string) int {
	sum := 0
	for _, char := range strng {
		sum += int(char - '0')
	}

	return sum
}

func OrderWeight(strng string) string {
	nums := strings.Fields(strng)

	sort.Slice(nums, func(i, j int) bool {
		weighti := Weight(nums[i])
		weightj := Weight(nums[j])
		if weighti == weightj {
			return nums[i] < nums[j]
		}
		return weighti < weightj
	})

	return strings.Join(nums, " ")
}
