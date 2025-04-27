package kata

import (
	"fmt"
)

func SeriesSum(n int) string {
	var res float64 = 0
	for i := 0; i < n; i++ {
		res += float64(1) / float64(1+3*i)
	}
	return fmt.Sprintf("%.2f", res)
}
