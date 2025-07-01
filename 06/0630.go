package kata

func sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func FindEvenIndex(arr []int) int {
	for i := 0; i < len(arr); i++ {
		if sum(arr[:i]) == sum(arr[i+1:]) {
			return i
		}
	}
	return -1
}

// package kata

// func FindEvenIndex(arr []int) int {
//   var sum, b int
//   for _,n := range arr { sum += n }
//   for i,n := range arr {
//     sum -= n
//     if sum == b { return i }
//     b += n
//   }
//   return -1
// }
