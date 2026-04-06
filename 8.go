package main

import (
	"fmt"
)

func sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, "-nji(y) gezek", "Jogap:", z)
	}
	return z
}

func main() {
	var x float64
	fmt.Println("x-y girizin: ")
	fmt.Scan(&x)
	res := sqrt(x)
	fmt.Println("Netije", res)
}
