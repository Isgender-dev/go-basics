package main

import "fmt"

func main() {
	for {
		var x int
		fmt.Print("Sany girizin: ")
		fmt.Scan(&x)

		if x == 0 {
			fmt.Printf("%d \n", x)
			continue
		}

		isPrime := true

		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d asal sayıdır\n", x)
		} else {
			fmt.Printf("%d asal sayı değildir\n", x)
		}
	}
}
