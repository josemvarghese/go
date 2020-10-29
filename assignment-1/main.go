package main

import "fmt"

type numbers []int

func createNumbers(i int) numbers {
	count := 0
	n := numbers{}
	for count <= i {
		n = append(n, count)
		count++
	}
	return n
}
func main() {
	oddeven := createNumbers(20)
	for _, ov := range oddeven {
		if ov%2 == 0 {
			fmt.Printf("%d is even number \n", ov)
		} else {
			fmt.Printf("%d is odd number \n", ov)
		}

	}
}
