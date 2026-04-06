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

	freq := make(map[rune]int)
	for _, v := range word {
		freq[v]++
	}
	var char rune
	var count int
	for i, v := range freq {
		if v > count {
			char = i
			count = v
		}
	}
	fmt.Printf("Most frequent: %c (%d times)", char, count)
}
