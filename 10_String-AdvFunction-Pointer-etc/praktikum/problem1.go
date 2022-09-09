package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if len(a) < len(b) {
		b, a = a, b
	}

	isContain := strings.Contains(a, b)

	if isContain {
		return b
	}

	return ""
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
