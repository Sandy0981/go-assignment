package main

import "fmt"

func main() {
	var num = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Index is :", findIndex(num, 7))
}

func findIndex(num []int, x int) int {
	for i, v := range num {
		if v == x {
			return i
		}
	}
	return -1
}
