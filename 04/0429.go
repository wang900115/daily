package kata

func DirReduc(arr []string) []string {
	mapping := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	stack := []string{}

	for _, str := range arr {
		n := len(stack)
		if n > 0 && mapping[str] == stack[n-1] {
			stack = stack[:n-1]
		} else {
			stack = append(stack, str)
		}
	}
	return stack
}
