package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func receiver(ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int, 5)

	for i := 0; i < 3; i++ {
		go sender(ch)
	}

	for i := 0; i < 2; i++ {
		go receiver(ch)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("End of program")
}
