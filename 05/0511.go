package kata

import "fmt"

func Snail(snaipMap [][]int) []int {

	res := []int{}

	minX, maxX := 0, len(snaipMap[0])-1
	minY, maxY := 0, len(snaipMap)-1

	for minX <= maxX && minY <= maxY {

		// left to right , minY+1
		for i := minX; i <= maxX; i++ {
			res = append(res, snaipMap[minY][i])
		}
		minY++

		// top to bottom , maxX-1
		for j := minY; j <= maxY; j++ {
			res = append(res, snaipMap[j][maxX])
		}
		maxX--

		// right to left , maxY -1
		for i := maxX; i >= minX; i-- {
			res = append(res, snaipMap[maxY][i])
		}
		maxY--

		// bottom to top , minX +1
		for j := maxY; j >= minY; j-- {
			res = append(res, snaipMap[j][minX])
		}
		minX++
	}

	return res
}

func main() {
	var snaipMap = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	fmt.Println(Snail(snaipMap))
}
