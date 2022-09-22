package main

import "fmt"

func main() {
	var sentence string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit; sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	var freqChannel chan map[string]int = make(chan map[string]int)

	go getFrequencies(sentence, freqChannel)

	for letter, freq := range <-freqChannel {
		fmt.Println(letter, ":", freq)
	}
}

func getFrequencies(sentence string, channel chan map[string]int) {
	var frequencies map[string]int = map[string]int{}

	for _, letter := range sentence {
		frequencies[string(letter)]++
	}

	channel <- frequencies
}
