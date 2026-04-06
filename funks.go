package main

import "fmt"

func main() {
	func() {
		fmt.Printf("Hello ")

	}()

	sayWorld := func() {
		fmt.Print("say Hello world")
	}

	sayWorld()
}
