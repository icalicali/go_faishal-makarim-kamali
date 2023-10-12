package main

import "sort"

func main() {

}
//
func maxBuy(money int, products []int) int {
	sort.Ints(products)

	var i int = 0
	var maxAmount int = 0

	for money >= 0 && i < len(products) {
		money -= products[i]
		maxAmount++
		i++
	}
	return maxAmount - 1
}
