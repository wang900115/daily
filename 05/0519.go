package kata

func Spiralize(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	// 每次畫完一圈 往內+=2
	for offset := 0; size-2*offset >= 1; offset += 2 {
		if offset-1 > 0 {
			matrix[offset][offset-1] = 1
		}

		// left -> right
		for i := offset; i < size-offset; i++ {
			matrix[offset][i] = 1
		}

		// top -> down
		for j := offset + 1; j < size-offset; j++ {
			matrix[j][size-offset-1] = 1
		}

		// right -> left
		for i := size - offset - 1; i > offset; i-- {
			matrix[size-offset-1][i] = 1
		}

		// down -> top
		for j := size - offset - 1; j > offset+1; j-- {
			matrix[j][offset] = 1
		}

	}
	return matrix
}
