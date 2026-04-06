package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hi guys, welcome to golang")

	var conferenceName string = "Go Conference"

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	var conferenceTickets int = 400

	var remainingTickets int = 300

	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

}
