package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Printf("Write your word: ")
	fmt.Scan(&word)
	word = strings.ToLower(word)
	count := 0
	vowels := "aeiouy"
	for _, v := range word {

		if strings.ContainsRune(vowels, v) {
			count++
		}
	}
	fmt.Printf("Number of vowels: %v", count)
}
