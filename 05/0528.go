package kata

import (
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	res := "("
	for index, number := range numbers {

		if index == 3 {
			res += ") "
		}
		if index == 6 {
			res += "-"
		}

		res += strconv.Itoa(int(number))
	}

	return res
}
