package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {

	arrayC := append(arrayA, arrayB...)

	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range arrayC {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

}
