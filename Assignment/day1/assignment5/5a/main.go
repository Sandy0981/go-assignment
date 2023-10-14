package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
