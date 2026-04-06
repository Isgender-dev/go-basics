package main

import (
	"fmt"
	"math"
)

var x, y float64

func sqrt(a float64) string {
	if a < 0 {
		return sqrt(-a) + "i"
	}
	return fmt.Sprint(math.Sqrt(a))
}

func main() {
	fmt.Printf("x = ")
	fmt.Scan(&x)
	fmt.Printf("y = ")
	fmt.Scan(&y)
	fmt.Printf("%v; %v\n", sqrt(x), sqrt(y))
}
