package main

import "fmt"

func greet(user string) {
	fmt.Println("Hello" + user)
}
func main() {
	var user string
	fmt.Printf("What is your name? ")
	fmt.Scan(&user)
	greet(user)
}
