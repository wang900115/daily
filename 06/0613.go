package kata

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	vertical := 0
	horizontal := 0

	for _, char := range walk {
		switch char {
		case 'n':
			vertical++
		case 's':
			vertical--
		case 'w':
			horizontal--
		case 'e':
			horizontal++
		}
	}

	return vertical == 0 && horizontal == 0
}
