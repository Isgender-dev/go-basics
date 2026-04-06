package main

import (
	"fmt"
)

func main() {
	nums := []int{}
	var x int

	fmt.Println("Sanlary giriz (0 girizseň durýar):")

	for {
		fmt.Print("> ")
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		nums = append(nums, x)
	}

	if len(nums) == 0 {
		fmt.Println("San girizilmedi.")
		return
	}

	max := nums[0]

	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	fmt.Println("Iň uly san:", max)
}
