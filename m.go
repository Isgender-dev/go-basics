package main

import "fmt"

func at(x string) string {
	if len(x) == 0 {
		return ""
	}
	return at(x[1:]) + x[:1]
}

func main() {
	var x string
	fmt.Printf("Adynyzy girizin: ")
	fmt.Scan(&x)
	fmt.Printf("Adynyz tersine: %s\n", at(x))
}
