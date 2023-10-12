package main

import "fmt"
//
func PairSum(arr []int, target int) []int {
	var result []int = []int{}

	var storage map[int]int = map[int]int{}

	for index, b := range arr {
		var a int = target - b

		_, exist := storage[a]
		if exist {
			result = append(result, storage[a], index)
			break
		}
		storage[b] = index
	}

	return result
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}
