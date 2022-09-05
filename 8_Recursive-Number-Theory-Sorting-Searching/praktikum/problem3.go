package main

import (
	"fmt"
	"math"
)

func main() {
	primaSegiEmpat(2, 3, 13)
	/*
		17 19
		23 29
		31 37
	*/
}

func primaSegiEmpat(high, wide, start int) {
	var sum int = 0

	for i := 0; i < wide; i++ {
		for j := 0; j < high; j++ {
			start = nextPrime(start + 1)
			sum += start
			fmt.Print(start)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func nextPrime(start int) int {
	if checkPrime(start) {
		return start
	}
	return nextPrime(start + 1)
}

func checkPrime(n int) bool {
	if n == 1 {
		return false
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
