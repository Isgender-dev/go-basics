package main

import "fmt"

type User struct {
	Id      int64
	Name    string
	Age     int
	friends []int64
}

func main() {
	fmt.Printf("value: %v type: %T\n", struct{}{}, struct{}{})
	var str struct{}
	fmt.Printf("value: %v type: %T\n", str, str)
	var user User
	fmt.Printf("value: %vn type: %T\n", user, user)
}
