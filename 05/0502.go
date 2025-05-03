package kata

import "fmt"

func clamp(x int) int {
	if x < 0 {
		return 0
	}
	if x > 255 {
		return 255
	}
	return x
}

func RGB(r, g, b int) string {
	return fmt.Sprintf("%02X%02X%02X", clamp(r), clamp(g), clamp(b))
}
