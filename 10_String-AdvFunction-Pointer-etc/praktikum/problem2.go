package main

import "fmt"

func caesar(offset int, input string) string {
	result := make([]byte, 0, len(input))

	for _, char := range input {
		convert := (int(char)+offset-97)%26 + 97
		result = append(result, byte(convert))
	}

	return string(result)
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
}
