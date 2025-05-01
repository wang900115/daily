package kata

import "strconv"

func RotateString(str string) string {
	if len(str) <= 1 {
		return str
	}
	return str[1:] + string(str[0])
}

func ReverseString(str string) string {
	runes := []rune(str)
	for i := 0; i < len(str)/2; i++ {
		runes[i], runes[len(str)-1-i] = runes[len(str)-1-i], runes[i]
	}
	return string(runes)
}

func SumString(str string) int {
	sum := 0
	for _, char := range str {
		d, _ := strconv.Atoi(string(char))
		sum += d
	}
	return sum
}

func Revrot(s string, n int) string {
	if s == "" || n <= 0 {
		return ""
	}
	res := ""
	for i := 0; i+n <= len(s); i += n {
		chunk := s[i : i+n]

		if SumString(chunk)%2 == 0 {
			res += ReverseString(chunk)
		} else {
			res += RotateString(chunk)
		}
	}
	return res
}
