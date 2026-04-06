package main

import "fmt"

func main() {
	for {
		var x int
		fmt.Printf("Sany girizin: ")
		fmt.Scan(&x)
		if x == 0 {
			fmt.Println("0 ne jubut ne tak sandyr")
			break
		}
		if x%2 == 0 {
			fmt.Printf("%d jubut san\n", x)
		} else {
			fmt.Printf("%d tak san\n", x)
		}
	}
}
