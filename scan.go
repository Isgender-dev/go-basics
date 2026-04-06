package main

import "fmt"

func main() {
	var userName string
	var userEmail string
	var userAge uint

	for {
		fmt.Println("Enter your name:	")
		fmt.Scan(&userName)
		fmt.Println("Enter your email:  ")
		fmt.Scan(&userEmail)
		fmt.Printf("Enter your age: ")
		fmt.Scan(&userAge)
		fmt.Printf("My name is %v and I am %v years old.\n", userName, userAge)
		fmt.Printf("My email is %v\n", userEmail)
	}

}
