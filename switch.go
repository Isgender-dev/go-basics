package main

import "fmt"

func main() {
	switch num := 4; num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("I don't know")
	}
}
