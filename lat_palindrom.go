package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string = "biji"
	fmt.Println("kata: ", word)
	fmt.Println("palindrom? ", palindrom(word))
}

func palindrom(word string) bool {
	n := len(word)
	runes := make([]rune, n)
	for _, rune := range word {
		n--
		runes[n] = rune
	}

	fmt.Println("String reversed: ", string(runes[n:]))
	if strings.Compare(word, string(runes[n:])) == -1 {
		return false
	} else {
		return true
	}
}
