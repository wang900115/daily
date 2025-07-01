package kata

import (
	"strings"
)

func ToJadenCase(str string) string {
	res := []string{}
	for _, char := range strings.Split(str, " ") {
		n := len(char)
		res = append(res, strings.ToUpper(string(char[0]))+char[1:n])
	}
	restring := strings.Join(res, " ")
	return restring
}
