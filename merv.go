package main

import "fmt"

func main() {
	var letter string
	for {
		fmt.Printf("Enter a letter: ")
		fmt.Scan(&letter)
		switch letter {
		case "a":
			fmt.Println("Bu gun a harpyny owrendik")
		case "b":
			fmt.Println("Bu gun b harpyny owrendik")
		case "c":
			fmt.Println("Bu gun c harpyny owrendik")
		default:
			fmt.Println("Biz bu gun sapak gecmedik")
		}
	}
}
