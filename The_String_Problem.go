package main

import (
	"fmt"
	"strings"
)

func main() {
	var testCases int
	var str string
	// var n int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		fmt.Scan(&str)
		if vowelsPresent(&str) {
			fmt.Println("lovely string")
		} else {
			fmt.Println("ugly string")
		}
	}
}

func vowelsPresent(input *string) bool {
	ml := "aeiou"
	mu := "AEIOU"

	for _, c := range *input {
		s := string(c)
		if mu = strings.Replace(mu, s, "", 1); len(mu) == 0 {
			return true
		}

		if ml = strings.Replace(ml, s, "", 1); len(ml) == 0 {
			return true
		}
	}

	return false
}
