package main

import "fmt"

func add(x int, y int) (int, int) {
	s1 := x * x
	s2 := y * y
	return s1, s2
}
func main() {
	var x, y int
	fmt.Printf("x = ")
	fmt.Scan(&x)
	fmt.Printf("y = ")
	fmt.Scan(&y)
	result1, result2 := add(x, y)
	fmt.Printf("Result1: %d; Result2: %d\n", result1, result2)
}
