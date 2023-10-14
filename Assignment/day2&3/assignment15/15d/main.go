package main

import (
	"fmt"
	"sync"
)

func sender(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	wg.Done()
}

func receiver(ch chan int, wg *sync.WaitGroup) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sender(ch, &wg)
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go receiver(ch, &wg)
	}

	wg.Wait()

	close(ch)

	fmt.Println("End of program")
}
