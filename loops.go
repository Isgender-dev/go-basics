package main

import "fmt"

func main() {

	var phoneNumber int

	for {
		fmt.Printf("Please give me your phone number")
		fmt.Printf("Yeah, write:  +993")
		fmt.Scan(&phoneNumber)
	}
}
