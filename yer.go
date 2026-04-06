package main

import "fmt"

//func calys(a int, b int) (int, int) {
//	return b, a
//}

func main() {
	var a, b int
	fmt.Printf("Sanlary girizin:\n")
	fmt.Printf(" a = ")
	fmt.Scan(&a)
	fmt.Printf(" b = ")
	fmt.Scan(&b)
	// a, b = b, a

	a = a + b
	b = a - b
	a = a - b

	// a, b = calys(a, b)

	fmt.Printf(" a = %v, b = %v\n", a, b)

}
