package main

import "fmt"

func main() {
	arr := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slc := []int{19, 18, 17, 16, 15, 14, 13, 12, 11, 10}

	fmt.Printf("array: %T, slice: %T\n", arr, slc)
	fmt.Printf("array[1]: %d, slice[1]: %d\n", arr[1], slc[1])
	arr[3] = 45
	slc[3] = 55

	fmt.Printf("array[3] %d, slice[3] %d\n", arr[3], slc[3])
	fmt.Printf("array[1:4]: %v, slice[1:4]: %v\n", arr[1:4], slc[1:4])

	fmt.Printf("len(array): %d, cap(array) %d\n", len(arr), cap(arr))

	slc = append(slc, 321)

	fmt.Printf("After append:\n")
	fmt.Printf("len(slice): %d, cap(slice): %d\n", len(slc), cap(slc))

	slc = append(slc[:8], slc[8+1:]...)
	fmt.Printf("After delete:\n")
	fmt.Printf("len(slice): %d, cap(slice): %d\n", len(slc), cap(slc))
	sub := slc[2:5:7]
	fmt.Printf("sub := slc[2:5:7] - %v, len: %d, cap: %d\n", sub, len(sub), cap(sub))

}
