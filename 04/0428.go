package kata

func LongestConsec(strarr []string, k int) string {
	if len(strarr) > k || k < 0 {
		return ""
	}

	var longest string
	for i := 0; i <= len(strarr)-k; i++ {
		concatStr := ""
		for j := 0; j < k; j++ {
			concatStr += strarr[i+j]
		}

		if len(concatStr) > len(longest) {
			longest = concatStr
		}
	}

	return longest
}
