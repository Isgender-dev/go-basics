package main

import "fmt"

func main() {
	for {
		var x int
		fmt.Printf("Sayi giriniz: ")
		fmt.Scan(&x)
		if x == 0 || x == 1 {
			fmt.Printf("%d asal sayi degildir\n", x)
			continue
		}
		if x == 2 {
			fmt.Printf("%d en kucuk asal sayidir\n", x)
		}
		for i := 2; i < x; i++ {
			if x%i == 0 {
				fmt.Printf("%d asal sayi degildir\n", x)
				break
			} else {
				fmt.Printf("%d asal sayidir\n", x)
			}
		}
		break
	}
}
