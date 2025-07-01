package kata

func GetMiddle(s string) string {
	n := len(s)
	if n%2 == 0 {
		return s[n/2-1 : n/2+1]
	}
	return string(s[n/2])
}
