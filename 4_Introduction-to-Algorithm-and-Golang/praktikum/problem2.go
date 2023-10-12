package main

import "fmt"
//
func main() {

	var i int

	fmt.Printf("Input : ")
	fmt.Scanf("%d", &i)

	if i%2 == 0 {
		fmt.Print("lampu menyala")
	} else {
		fmt.Print("lampu mati")
	}
}
