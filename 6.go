package main

import "fmt"

func main() {
	for {
		var x int
		fmt.Print("Sayı giriniz: ")
		fmt.Scan(&x)

		if x == 0 || x == 1 {
			fmt.Printf("%d asal sayı değildir\n", x)
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

// primes=[2]
// for i=1 den N cenli
// for j=primes[:len(primes)-1]+1; true; j++
// If isPrime(j)
// primes append
// break
