package main

import "fmt"

func main() {
	var at rune
	fmt.Printf("Adynyzy girizin: ")
	fmt.Scan(&at)
	fmt.Printf("Adynyz tersine: ")

	for i := len(at) - 1; i >= 0; i-- {
		fmt.Printf("%c", at[i])
	}
}
