package main

import "fmt"

func fakt(n int) int {
	if n == 0 {
		return 1
	}
	return n * fakt(n-1) // 5 * 4!  
}

func main() {
	var n int
	fmt.Println("Sany girizin: ")
	fmt.Scan(&n)
	fmt.Printf("%d! = %d\n", n, fakt(n))
}
