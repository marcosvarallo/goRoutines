package main

import "fmt"

func main() {
	c := make(chan int)

	go send1(c)

	receive1(c)

	fmt.Println("about to exit")
}

// send channel
func send1(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// receive channel
func receive1(c <-chan int) {
	for v := range c {
		fmt.Println("the value received from the channel:", v)
	}
}
