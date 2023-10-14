package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randNum := rand.Intn(10)
	fmt.Println("Enter a guess b/w 1 to 10")
	var guess int
	fmt.Scanln(&guess)
	for randNum != guess {
		if randNum < guess {
			fmt.Println("Too high")
		} else {
			fmt.Println("Too low")
		}
		fmt.Println("Enter another value")
		fmt.Scanln(&guess)
	}
	fmt.Println("Your guess is correct!!!!")
}
