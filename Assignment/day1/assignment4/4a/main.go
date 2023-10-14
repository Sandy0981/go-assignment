package main

import "fmt"

var a int

func main() {
	fmt.Println("To check whether a number is even or odd")
	fmt.Println("Enter a nunber")
	fmt.Scanln(&a)
	fmt.Println(evenOdd(a))
}

func evenOdd(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
