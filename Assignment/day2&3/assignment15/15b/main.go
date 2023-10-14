package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
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

	go sender(ch)

	go receiver(ch)

	time.Sleep(2 * time.Second)
}
