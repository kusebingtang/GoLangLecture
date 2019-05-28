package main

import "fmt"

func main() {
	/*一个通道相当于一个先进先出（FIFO）的队列*/
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
}
