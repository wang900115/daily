package kata

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	parts := strings.Fields(in)
	sorted := make([]int, len(parts))

	for i, char := range parts {
		ints, _ := strconv.Atoi(char)
		sorted[i] = ints
	}
	sort.Ints(sorted)

	return fmt.Sprintf("%d %d", sorted[len(sorted)-1], sorted[0])
}
