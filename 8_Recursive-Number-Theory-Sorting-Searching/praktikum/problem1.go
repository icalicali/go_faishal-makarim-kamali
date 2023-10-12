package main

import (
	"fmt"
	"math"
)
//
func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}

func primeX(number int) int {
	if number == 1 {
		return 2
	}

	if number < 1 {
		return 0
	}

	var countPrime int = 1
	var i int = 3

	for countPrime < number {
		if checkPrime(i) {
			countPrime++
		}
		i++
	}

	return i - 1
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
