package main

import "fmt"

func main() {

}

func FindMinAndMax(arr []int) {
	var minimum int = arr[0]
	var maximum int = arr[0]

	var minIndex int = 0
	var maxIndex int = 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < minimum {
			minimum = arr[i]
			minIndex = i
		} else if arr[i] > maximum {
			maximum = arr[i]
			maxIndex = i
		}
	}

	fmt.Println("minimum: ", minimum, " index: ", minIndex)
	fmt.Println("maximum: ", maximum, " index: ", maxIndex)
}
