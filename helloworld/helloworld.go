package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[1:])
	fmt.Println(s[1:2])
	fmt.Println(s[1:3])
	fmt.Println(s[:3])
	// fmt.Println(s[1:6])
	fmt.Println("Hello World")
}
