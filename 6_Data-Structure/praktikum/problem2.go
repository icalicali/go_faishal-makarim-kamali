package main

import (
	"fmt"
	"strconv"
)
//
func munculSekali(angka string) []int {
	var result []int
	var freq map[string]int = map[string]int{}

	for _, num := range angka {
		freq[string(num)]++
	}
	for key, freq := range freq {
		if freq == 1 {
			res, _ := strconv.Atoi(key)
			result = append(result, res)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
