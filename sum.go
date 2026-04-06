package main

import "fmt"

func sum(x int) int {
	if x > 0 {
		return sum(x-1) + x
		// return sum(x-2) + x // jubut we tak sanlaryn jemi ucin
	}
	return 0
}
func main() {
	var x int
	fmt.Printf("Sany girizin: ")
	fmt.Scan(&x)
	sum := sum(x)
	fmt.Printf("Sum: %d\n", sum)
}
