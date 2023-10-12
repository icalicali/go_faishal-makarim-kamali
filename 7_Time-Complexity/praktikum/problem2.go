package main

import (
	"fmt"
)
//
func pow(x, n int) int {
	var result int = 1

	for i := 1; i <= n; i++ {
		result *= x
	}
	return result
}

func powNew(x, n int) int {
	var result int = 1
	if n < 1 {
		return 1
	}
	for n >= 1 {
		if n%2 != 0 {
			result = result * x
			n--
		}

		x = x * x
		n = n / 2
	}
	return result
}

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 3))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}
