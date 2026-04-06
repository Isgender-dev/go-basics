package main

import "fmt"

func swap(x, y, z string) (string, string, string) {
	return z, x, y
}

func main() {
	x, y, z := swap("20", "years old", "I'm")
	fmt.Println(x, y, z)
}
