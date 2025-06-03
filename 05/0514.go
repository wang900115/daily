package kata

func Determinant(maxtrix [][]int) int {
	n := len(maxtrix)
	if n == 1 {
		return maxtrix[0][0]
	}

	if n == 2 {
		return maxtrix[0][0]*maxtrix[1][1] - maxtrix[0][1]*maxtrix[1][0]
	}

	det := 0

	for i := 0; i < n; i++ {
		minor := getMinor(maxtrix, 0, i)
		cofactor := maxtrix[0][i]
		if i%2 == 1 {
			cofactor *= -1
		}

		det += cofactor * Determinant(minor)
	}

	return det

}

func getMinor(maxtrix [][]int, r, c int) [][]int {
	size := len(maxtrix)
	minor := [][]int{}
	for i := 0; i < size; i++ {
		if i == r {
			continue
		}
		row := []int{}
		for j := 0; j < size; j++ {
			if j == c {
				continue
			}

			row = append(row, maxtrix[i][j])
		}
		minor = append(minor, row)
	}
	return minor
}
