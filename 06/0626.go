package kata

import "strconv"

func StringToNumber(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
