package main

import (
	"fmt"
)

func main() {
	var numbers = []int{3, 4, 5, 2, 6, 7, 8}
	fmt.Println("Reverse Slice ", reverse(numbers))
}
func reverse(a []int) []int {
	for i := 0; i < len(a)/2; i++ {
		var b int
		b = a[i]
		a[i] = a[len(a)-i-1]
		a[len(a)-i-1] = b
	}
	return a
}
