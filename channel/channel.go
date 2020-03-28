package main

import (
	"fmt"
	"time"
)

func work(n int, c chan int) {
	for {
		fmt.Printf("Worker %d received %c\n", n, <-c)
	}
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go work(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}


	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
