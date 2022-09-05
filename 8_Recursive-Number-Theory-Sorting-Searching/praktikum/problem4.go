package main

import "fmt"

func main() {
	fmt.Println(MaxSequence([]int{-2, -1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))    // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))    // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))    // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))     // 12
}

func MaxSequence(arr []int) int {
	var sum int = 0

	if len(arr)%2 == 0 {
		var middle int = len(arr)/2 - 1

		for i := middle - 1; i <= middle+3; i++ {
			sum += arr[i]
		}
	} else {
		var middle int = len(arr) / 2

		for i := middle - 1; i <= middle+2; i++ {
			sum += arr[i]
		}
	}
	return sum
}
