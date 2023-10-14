package main

import "fmt"

func main() {
	var num = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("The sum of even num in slice is", sum(num))
}

func sum(a []int) int {
	var sum int
	for i := range a {
		if a[i]%2 == 0 {
			sum += a[i]
		}
	}
	return sum
}
