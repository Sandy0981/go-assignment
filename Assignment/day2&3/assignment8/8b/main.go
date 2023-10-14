package main

import "fmt"

var ans int = 1

func main() {
	fact(5, ans)
}
func fact(n int, ans int) {
	if n == 1 {
		fmt.Println("Factorial is", ans)
		return
	}
	ans *= n
	fact(n-1, ans)
}
