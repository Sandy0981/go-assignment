package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = []int{1, 2, 3, 4, 5}
	fmt.Println(compSlice(slice1, slice2))
}

func compSlice(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
