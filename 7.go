package main

import "fmt"

func num(x int) {
	if x > 0 {
		fmt.Printf("%d ", x)
		num(x - 1)
	}
}
func main() {
	var x int
	fmt.Printf("Sany girizin: ")
	fmt.Scan(&x)
	num(x)

}
